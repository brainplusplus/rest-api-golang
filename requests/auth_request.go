package requests

type AuthLoginAdminRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

type AuthLoginCustomerRequest struct {
	AuthLoginAdminRequest
}
