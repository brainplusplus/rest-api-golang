package requests

type ProductBaseRequest struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

type ProductSaveRequest struct {
	ProductBaseRequest
}

type ProductUpdateRequest struct {
	Id int64 `json:"id"`
	ProductBaseRequest
}
