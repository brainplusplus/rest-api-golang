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

type AdminService interface {
	Login(req *requests.AuthLoginAdminRequest) *responses.Response
	GetAll() *responses.ResponseList
	GetById(id int64) *responses.Response
	GetByEmail(email string) *responses.Response
	DeleteById(id int64) *responses.Response
	Save(req *requests.AdminSaveRequest) *responses.Response
	Update(req *requests.AdminUpdateRequest) *responses.Response
	//ChangeProfile(req *requests.ProfileRequest) *responses.Response
	//ChangePassword(req *requests.ProfileChangePasswordRequest) *responses.Response
}

type adminService struct {
}

func GetAdminService() AdminService {
	return &adminService{}
}

// FindProfileById search Admin profile by Admin id and returns AdminProfile
func (r *adminService) Login(req *requests.AuthLoginAdminRequest) *responses.Response {
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
	data, err := repositories.GetAdminRepository().GetByEmail(req.Email)
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
				Role:     "Admin",
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

func (r *adminService) GetAll() *responses.ResponseList {
	resp := new(responses.ResponseList)
	data, err := repositories.GetAdminRepository().GetAll()
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Rows = data
		resp.Success = true
	}
	return resp
}

func (r *adminService) GetByEmail(email string) *responses.Response {
	resp := new(responses.Response)
	data, err := repositories.GetAdminRepository().GetByEmail(email)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *adminService) GetById(id int64) *responses.Response {
	resp := new(responses.Response)
	data, err := repositories.GetAdminRepository().GetById(id)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
	}
	return resp
}

func (r *adminService) DeleteById(id int64) *responses.Response {
	resp := new(responses.Response)
	err := repositories.GetAdminRepository().DeleteById(id)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Success = true
		resp.Message = "Success Delete"
	}
	return resp
}

func (r *adminService) Save(req *requests.AdminSaveRequest) *responses.Response {
	resp := new(responses.Response)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		log.Info(err.Error())
		resp.Message = err.Error()
		return resp
	}
	reqData := models.Admin{}
	reqData.Password = req.Password
	reqData.Name = req.Name
	reqData.Email = req.Email
	checkAdmin, err := repositories.GetAdminRepository().GetByEmail(req.Email)
	if checkAdmin == nil && err == sql.ErrNoRows {
		data, err2 := repositories.GetAdminRepository().Save(&reqData)
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

func (r *adminService) Update(req *requests.AdminUpdateRequest) *responses.Response {
	resp := new(responses.Response)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		log.Info(err.Error())
		resp.Message = err.Error()
		return resp
	}
	reqData, _ := repositories.GetAdminRepository().GetById(req.Id)
	reqData.Email = req.Email
	if req.Password != reqData.Password && req.Password != "" {
		reqData.Password, _ = utils.HashPassword(req.Password)
	}
	reqData.Name = req.Name
	reqData.Email = req.Email
	data, err := repositories.GetAdminRepository().Update(reqData)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Data = data
		resp.Success = true
		resp.Message = "Success Update"
	}
	return resp
}
