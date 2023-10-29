package requests

type AdminBaseRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

type AdminSaveRequest struct {
	AdminBaseRequest
}

type AdminUpdateRequest struct {
	Id int64 `json:"id"`
	AdminBaseRequest
}

type AdminRegisterRequest struct {
	AdminBaseRequest
}
