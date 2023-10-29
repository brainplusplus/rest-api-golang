package requests

import "simple-ecommerce/types"

type OrderBaseRequest struct {
	CustomerId int64          `json:"customer_id"`
	OrderDate  types.JsonDate `json:"order_date"`
	Status     string         `json:"status"`
}

type OrderSaveRequest struct {
	OrderBaseRequest
	OrderProducts []OrderProduct `json:"order_products" validate:"required,min=0"`
}

type OrderUpdateRequest struct {
	Id int64 `json:"id"`
	OrderBaseRequest
	OrderProducts []OrderProduct `json:"order_products" validate:"required,min=0"`
}

type OrderProduct struct {
	//Id         int64   `json:"id"`
	//OrderId    int64   `json:"order_id"`
	ProductId int64 `json:"product_id" validate:"required,gt=0"`
	Quantity  int   `json:"quantity" validate:"required,gt=0"`
	//Price      float64 `json:"price"`
	//TotalPrice float64 `json:"total_price"`
}
