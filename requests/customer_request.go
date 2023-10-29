package requests

type CustomerBaseRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

type CustomerSaveRequest struct {
	CustomerBaseRequest
}

type CustomerUpdateRequest struct {
	Id int64 `json:"id"`
	CustomerBaseRequest
}

type CustomerRegisterRequest struct {
	CustomerBaseRequest
}
