package models

import (
	"simple-ecommerce/types"
)

type OrderTable struct {
	Id         int64          `json:"id"`
	Token      string         `json:"-" db:"token"`
	CustomerId int64          `json:"customer_id" db:"customer_id"`
	OrderDate  types.JsonDate `json:"order_date" db:"order_date"`
	TotalPrice float64        `json:"total_price"  db:"total_price"`
	Status     string         `json:"status"`
}

type Order struct {
	OrderTable
	CustomerName  string         `json:"customer_name" db:"customer_name"`
	CustomerEmail string         `json:"customer_email" db:"customer_email"`
	OrderProducts []OrderProduct `json:"order_products"`
}
