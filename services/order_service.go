package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"simple-ecommerce/models"
	"simple-ecommerce/repositories"
	"simple-ecommerce/requests"
	"simple-ecommerce/responses"
)

type OrderService interface {
	GetAll() *responses.ResponseList
	GetAllByCustomerId(customerId int64) *responses.ResponseList
	GetById(id int64) *responses.Response
	GetAllWithProducts() *responses.ResponseList
	GetAllByCustomerEmailWithProducts(customerEmail string) *responses.ResponseList
	GetAllByCustomerIdWithProducts(customerId int64) *responses.ResponseList
	GetAllByStatusWithProducts(status string) *responses.ResponseList
	GetByIdWithProducts(id int64) *responses.Response
	GetByIdAndCustomerIdWithProducts(id int64, customerId int64) *responses.Response
	GetByTokenWithProducts(token string) *responses.Response
	DeleteById(id int64) *responses.Response
	Save(req *requests.OrderSaveRequest) *responses.Response
	Update(req *requests.OrderUpdateRequest) *responses.Response
	StatusUpdate(id int64, status string) error
}

type orderService struct {
}

func GetOrderService() OrderService {
	return &orderService{}
}

func (r *orderService) GetAll() *responses.ResponseList {
	resp := new(responses.ResponseList)
	data, err := repositories.GetOrderRepository().GetAll()
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Rows = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) GetAllByCustomerId(customerId int64) *responses.ResponseList {
	resp := new(responses.ResponseList)
	data, err := repositories.GetOrderRepository().GetAllByCustomerId(customerId)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Rows = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) GetById(id int64) *responses.Response {
	resp := new(responses.Response)
	data, err := repositories.GetOrderRepository().GetById(id)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) GetAllWithProducts() *responses.ResponseList {
	resp := new(responses.ResponseList)
	data, err := repositories.GetOrderRepository().GetAllWithProducts()
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Rows = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) GetAllByCustomerEmailWithProducts(customerEmail string) *responses.ResponseList {
	resp := new(responses.ResponseList)
	data, err := repositories.GetOrderRepository().GetAllByCustomerEmailWithProducts(customerEmail)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Rows = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) GetAllByCustomerIdWithProducts(customerId int64) *responses.ResponseList {
	resp := new(responses.ResponseList)
	data, err := repositories.GetOrderRepository().GetAllByCustomerIdWithProducts(customerId)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Rows = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) GetAllByStatusWithProducts(status string) *responses.ResponseList {
	resp := new(responses.ResponseList)
	data, err := repositories.GetOrderRepository().GetAllByStatusWithProducts(status)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Rows = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) GetByIdWithProducts(id int64) *responses.Response {
	resp := new(responses.Response)
	data, err := repositories.GetOrderRepository().GetByIdWithProducts(id)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) GetByIdAndCustomerIdWithProducts(id int64, customerId int64) *responses.Response {
	resp := new(responses.Response)
	data, err := repositories.GetOrderRepository().GetByIdAndCustomerIdWithProducts(id, customerId)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) GetByTokenWithProducts(token string) *responses.Response {
	resp := new(responses.Response)
	data, err := repositories.GetOrderRepository().GetByTokenWithProducts(token)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *orderService) DeleteById(id int64) *responses.Response {
	resp := new(responses.Response)
	err := repositories.GetOrderRepository().DeleteById(id)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Success = true
		resp.Message = "Success Delete"
	}
	return resp
}

func (r *orderService) Save(req *requests.OrderSaveRequest) *responses.Response {
	resp := new(responses.Response)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		log.Info(err.Error())
		resp.Message = err.Error()
		return resp
	}
	reqData := models.Order{}
	reqData.CustomerId = req.CustomerId
	reqData.OrderDate = req.OrderDate
	reqData.Status = req.Status
	reqData.OrderProducts = []models.OrderProduct{}
	for _, orderProductReq := range req.OrderProducts {
		orderProduct := models.OrderProduct{}
		orderProduct.ProductId = orderProductReq.ProductId
		orderProduct.Quantity = orderProductReq.Quantity
		if orderProduct.ProductId != 0 {
			product, err := repositories.GetProductRepository().GetById(orderProduct.ProductId)
			if err != nil {
				resp.Message = fmt.Sprintf("Error Occured : %s", err)
				return resp
			}
			orderProduct.Price = product.Price
			orderProduct.TotalPrice = orderProduct.Price * float64(orderProduct.Quantity)
			reqData.TotalPrice += orderProduct.TotalPrice

		} else {
			log.Info("product not found", orderProduct.ProductId)
			resp.Message = fmt.Sprintf("One of Product Not Found")
			return resp
		}
		reqData.OrderProducts = append(reqData.OrderProducts, orderProduct)
	}
	data, err := repositories.GetOrderRepository().Save(&reqData)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
		resp.Message = "Success Save"
	}
	return resp
}

func (r *orderService) Update(req *requests.OrderUpdateRequest) *responses.Response {
	resp := new(responses.Response)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		log.Info(err.Error())
		resp.Message = err.Error()
		return resp
	}
	reqData, _ := repositories.GetOrderRepository().GetById(req.Id)
	reqData.CustomerId = req.CustomerId
	reqData.OrderDate = req.OrderDate
	reqData.Status = req.Status
	reqData.OrderProducts = []models.OrderProduct{}
	for _, orderProductReq := range req.OrderProducts {
		orderProduct := models.OrderProduct{}
		orderProduct.ProductId = orderProductReq.ProductId
		orderProduct.Quantity = orderProductReq.Quantity
		if orderProduct.ProductId != 0 {
			product, err := repositories.GetProductRepository().GetById(orderProduct.ProductId)
			if err != nil {
				resp.Message = fmt.Sprintf("Error Occured : %s", err)
				return resp
			}
			orderProduct.Price = product.Price
			orderProduct.TotalPrice = orderProduct.Price * float64(orderProduct.Quantity)
			reqData.TotalPrice += orderProduct.TotalPrice

		} else {
			log.Info("product not found", orderProduct.ProductId)
			resp.Message = fmt.Sprintf("One of Product Not Found")
			return resp
		}
		reqData.OrderProducts = append(reqData.OrderProducts, orderProduct)
	}
	data, err := repositories.GetOrderRepository().Update(reqData)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
		resp.Message = "Success Update"
	}
	return resp
}

func (r *orderService) StatusUpdate(id int64, status string) error {
	err := repositories.GetOrderRepository().StatusUpdate(id, status)
	if err != nil {
		log.Info(err)
	}
	return err
}
