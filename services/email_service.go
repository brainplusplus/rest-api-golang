package services

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tushar2708/altcsv"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"simple-ecommerce/configs"
	"simple-ecommerce/objects"
	"simple-ecommerce/repositories"
	"simple-ecommerce/requests"
	"simple-ecommerce/responses"
	"simple-ecommerce/utils"
)

type EmailService interface {
	Send(req requests.EmailRequest) *responses.Response
	SendCsvOrdersToAdmins() *responses.Response
	SendEachPendingOrderToCustomer() *responses.Response
}

type emailService struct {
	credential *objects.SmtpCredential
	from       *string
}

func GetEmailService() EmailService {
	return &emailService{credential: smtpCredential, from: &emailFrom}
}

func (r *emailService) Send(req requests.EmailRequest) *responses.Response {
	var err error
	resp := new(responses.Response)
	m := gomail.NewMessage()
	m.SetHeader("From", req.From)
	m.SetHeader("To", req.To...)
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/html", req.Content)
	if len(req.Attachments) > 0 {
		for _, fileAttachment := range req.Attachments {
			m.Attach(fileAttachment)
		}
	}

	d := gomail.NewDialer(r.credential.Host, r.credential.Port, r.credential.Username, r.credential.Password)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	//can send multiple email in one request d.DialAndSend(m1,m2...)
	if err = d.DialAndSend(m); err != nil {
		resp.Message = err.Error()
		log.Info(resp.Message)
	} else {
		resp.Data = req
		resp.Success = true
	}
	return resp
}

func (r *emailService) SendCsvOrdersToAdmins() *responses.Response {
	resp := new(responses.Response)
	fileCsv, err := os.CreateTemp("", "orders-*.csv") // in Go version older than 1.17 you can use ioutil.TempFile
	if err != nil {
		resp.Message = err.Error()
		log.Info(err)
		return resp
	}
	//defer fileCsv.Close()
	defer os.Remove(fileCsv.Name())

	orders := [][]string{}
	headers := []string{"Order ID", "Customer Name", "Order Date", "Total Price", "Status"}
	orders = append(orders, headers)

	rows, err := repositories.GetOrderRepository().GetAll()
	if err != nil {
		resp.Message = err.Error()
		log.Info(err)
		return resp
	}
	for _, ord := range rows {
		ordArr := []string{
			fmt.Sprintf("%v", ord.Id),
			ord.CustomerName,
			ord.OrderDate.ToTimeString(),
			fmt.Sprintf("%v", ord.TotalPrice),
			ord.Status,
		}
		orders = append(orders, ordArr)
	}

	wr := altcsv.NewWriter(fileCsv)
	wr.AllQuotes = false
	//in My Ms Excel, CSV using this delimiter ; not ,
	wr.Comma = ';'

	// Write all items and deal with errors
	if err := wr.WriteAll(orders); err != nil {
		resp.Message = err.Error()
		log.Info(err)
		return resp
	}
	admins, err := repositories.GetAdminRepository().GetAll()
	if err != nil {
		resp.Message = err.Error()
		log.Info(err)
		return resp
	}
	var adminEmails []string
	for _, admin := range admins {
		adminEmails = append(adminEmails, admin.Email)
	}
	emailReq := requests.EmailRequest{}
	emailReq.From = *r.from
	emailReq.To = adminEmails
	emailReq.Subject = "All Orders"
	emailReq.Content = "<p>Below we attach a CSV file for all orders</p>"
	emailReq.Attachments = []string{fileCsv.Name()}
	resp = r.Send(emailReq)
	if !resp.Success {
		log.Info(resp.Message)
	}
	return resp
}

func (r *emailService) SendEachPendingOrderToCustomer() *responses.Response {
	resp := new(responses.Response)
	status := "Pending"
	pendingOrderList, err := repositories.GetOrderRepository().GetAllByStatusWithProducts(status)
	if err != nil {
		resp.Message = err.Error()
		log.Info(err)
		return resp
	}
	for _, pendingOrder := range pendingOrderList {
		emailReq := requests.EmailRequest{}
		emailReq.From = *r.from
		emailReq.To = []string{pendingOrder.CustomerEmail}
		emailReq.Subject = fmt.Sprintf("You have a pending order in %v", pendingOrder.OrderDate.ToTimeString())
		templateFileName := "email_order_reminder.html"
		pathFile := "./views/" + templateFileName
		t, err := template.New(templateFileName).Funcs(
			utils.GetAllFuncMap()).ParseFiles(pathFile)
		if err != nil {
			log.Info(err.Error())
		} else {
			buf := new(bytes.Buffer)
			if err = t.Execute(buf, fiber.Map{
				"Order":   pendingOrder,
				"BaseURL": configs.GetConfigString("server.base_url"),
			}); err != nil {
				log.Info(err.Error())
			} else {
				htmlContent := buf.String()
				emailReq.Content = htmlContent
				respSend := r.Send(emailReq)
				if !respSend.Success {
					log.Info(resp.Message)
				}
			}
		}
	}

	return resp
}
