package services

import (
	"github.com/go-playground/validator/v10"
	"simple-ecommerce/models"
	"simple-ecommerce/repositories"
	"simple-ecommerce/requests"
	"simple-ecommerce/responses"
)

type ProductService interface {
	GetAll() *responses.ResponseList
	GetById(id int64) *responses.Response
	DeleteById(id int64) *responses.Response
	Save(req *requests.ProductSaveRequest) *responses.Response
	Update(req *requests.ProductUpdateRequest) *responses.Response
}

type productService struct {
}

func GetProductService() ProductService {
	return &productService{}
}

func (r *productService) GetAll() *responses.ResponseList {
	resp := new(responses.ResponseList)
	data, err := repositories.GetProductRepository().GetAll()
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Rows = data
		resp.Success = true
	}
	return resp
}

func (r *productService) GetById(id int64) *responses.Response {
	resp := new(responses.Response)
	data, err := repositories.GetProductRepository().GetById(id)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *productService) DeleteById(id int64) *responses.Response {
	resp := new(responses.Response)
	err := repositories.GetProductRepository().DeleteById(id)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Success = true
		resp.Message = "Success Delete"
	}
	return resp
}

func (r *productService) Save(req *requests.ProductSaveRequest) *responses.Response {
	resp := new(responses.Response)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		log.Info(err.Error())
		resp.Message = err.Error()
		return resp
	}
	reqData := models.Product{}
	reqData.Price = req.Price
	reqData.Name = req.Name
	reqData.Description = req.Description
	reqData.Image = req.Image
	data, err := repositories.GetProductRepository().Save(&reqData)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
		resp.Message = "Success Save"
	}
	return resp
}

func (r *productService) Update(req *requests.ProductUpdateRequest) *responses.Response {
	resp := new(responses.Response)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		log.Info(err.Error())
		resp.Message = err.Error()
		return resp
	}
	reqData, _ := repositories.GetProductRepository().GetById(req.Id)
	reqData.Price = req.Price
	reqData.Name = req.Name
	reqData.Description = req.Description
	if req.Image != "" {
		reqData.Image = req.Image
	}
	data, err := repositories.GetProductRepository().Update(reqData)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
		resp.Message = "Success Update"
	}
	return resp
}
