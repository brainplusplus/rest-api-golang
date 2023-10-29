package services

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"simple-ecommerce/middlewares"
	"simple-ecommerce/models"
	"simple-ecommerce/repositories"
	"simple-ecommerce/requests"
	"simple-ecommerce/responses"
	"simple-ecommerce/utils"
	"time"
)

type CustomerService interface {
	Login(req *requests.AuthLoginCustomerRequest) *responses.Response
	GetAll() *responses.ResponseList
	GetById(id int64) *responses.Response
	GetByEmail(email string) *responses.Response
	DeleteById(id int64) *responses.Response
	Save(req *requests.CustomerSaveRequest) *responses.Response
	Update(req *requests.CustomerUpdateRequest) *responses.Response
	//ChangeProfile(req *requests.ProfileRequest) *responses.Response
	//ChangePassword(req *requests.ProfileChangePasswordRequest) *responses.Response
	Register(req *requests.CustomerRegisterRequest) *responses.Response
}

type customerService struct {
}

func GetCustomerService() CustomerService {
	return &customerService{}
}

// FindProfileById search customer profile by customer id and returns customerProfile
func (r *customerService) Login(req *requests.AuthLoginCustomerRequest) *responses.Response {
	resp := new(responses.Response)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		log.Info(err.Error())
		resp.Message = err.Error()
		return resp
	}
	log.Info("LOGIN")
	log.Info(req.Email)
	data, err := repositories.GetCustomerRepository().GetByEmail(req.Email)
	if err != nil {
		resp.Message = err.Error()
	} else {
		log.Info(req.Password + " = " + data.Password)
		match := utils.CheckPasswordHash(req.Password, data.Password)
		log.Info(match)
		if match {
			now := time.Now()
			et := middlewares.EasyToken{
				UserID:   data.Id,
				Username: req.Email,
				Role:     "User",
				Expires:  &jwt.NumericDate{now.Add(time.Hour * 24)},
			}
			resp.Data, _ = et.GetJWTToken()
		} else {
			resp.Message = "Invalid Email or password"
		}
		resp.Success = match
	}
	log.Info("END LOGIN")
	return resp
}

func (r *customerService) GetAll() *responses.ResponseList {
	resp := new(responses.ResponseList)
	data, err := repositories.GetCustomerRepository().GetAll()
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Rows = data
		resp.Success = true
	}
	return resp
}

func (r *customerService) GetByEmail(email string) *responses.Response {
	resp := new(responses.Response)
	data, err := repositories.GetCustomerRepository().GetByEmail(email)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *customerService) GetById(id int64) *responses.Response {
	resp := new(responses.Response)
	data, err := repositories.GetCustomerRepository().GetById(id)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *customerService) DeleteById(id int64) *responses.Response {
	resp := new(responses.Response)
	err := repositories.GetCustomerRepository().DeleteById(id)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Success = true
		resp.Message = "Success Delete"
	}
	return resp
}

func (r *customerService) Save(req *requests.CustomerSaveRequest) *responses.Response {
	resp := new(responses.Response)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		log.Info(err.Error())
		resp.Message = err.Error()
		return resp
	}
	reqData := models.Customer{}
	reqData.Password = req.Password
	reqData.Name = req.Name
	reqData.Email = req.Email
	checkcustomer, err := repositories.GetCustomerRepository().GetByEmail(req.Email)
	if checkcustomer == nil && err == sql.ErrNoRows {
		data, err2 := repositories.GetCustomerRepository().Save(&reqData)
		if err2 != nil {
			resp.Message = err2.Error()
		} else {
			resp.Data = data
			resp.Success = true
			resp.Message = "Success Save"
		}
	} else if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Message = "Already registered"
	}
	return resp
}

func (r *customerService) Update(req *requests.CustomerUpdateRequest) *responses.Response {
	resp := new(responses.Response)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		log.Info(err.Error())
		resp.Message = err.Error()
		return resp
	}
	reqData, _ := repositories.GetCustomerRepository().GetById(req.Id)
	reqData.Email = req.Email
	if req.Password != reqData.Password && req.Password != "" {
		reqData.Password, _ = utils.HashPassword(req.Password)
	}
	reqData.Name = req.Name
	reqData.Email = req.Email
	data, err := repositories.GetCustomerRepository().Update(reqData)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
		resp.Message = "Success Update"
	}
	return resp
}

func (r *customerService) Register(req *requests.CustomerRegisterRequest) *responses.Response {
	resp := new(responses.Response)
	reqData := models.Customer{}
	reqData.Password, _ = utils.HashPassword(req.Password)
	reqData.Name = req.Name
	reqData.Email = req.Email
	checkcustomer, err := repositories.GetCustomerRepository().GetByEmail(req.Email)
	if checkcustomer == nil && err == sql.ErrNoRows {
		data, err2 := repositories.GetCustomerRepository().Save(&reqData)
		if err2 != nil {
			resp.Message = err2.Error()
		} else {
			resp.Data = data
			resp.Success = true
			resp.Message = "Success Register"
		}
	} else if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Message = "Already registered"
	}
	return resp
}
