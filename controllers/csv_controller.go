package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"simple-ecommerce/models"

	//"encoding/csv"
	"github.com/tushar2708/altcsv"
	"simple-ecommerce/responses"
	"simple-ecommerce/services"
)

type CsvApiControllerInterface interface {
	ServeCSV(c *fiber.Ctx) error
}

type csvApiControllers struct{}

var (
	CsvApiController CsvApiControllerInterface = &csvApiControllers{}
)

func (ctl *csvApiControllers) ServeCSV(c *fiber.Ctx) error {
	resp := new(responses.Response)
	orders := [][]string{}
	headers := []string{"Order ID", "Customer Name", "Order Date", "Total Price", "Status"}
	orders = append(orders, headers)

	responseList := services.GetOrderService().GetAll()
	rows, ok := responseList.Rows.([]models.Order)
	if !ok {
		resp.Message = "Conversion Failed"
		log.Info(resp.Message)
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
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
	// Set our headers so browser will download the file
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment;filename=orders.csv")
	// Create a CSV writer using our HTTP response writer as our io.Writer
	wr := altcsv.NewWriter(c)
	wr.AllQuotes = false
	//in My Ms Excel, CSV using this delimiter ; not ,
	wr.Comma = ';'

	// Write all items and deal with errors
	if err := wr.WriteAll(orders); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp)
	}
	return nil
}
