package models

type OrderProductTable struct {
	Id         int64   `json:"id"`
	OrderId    int64   `json:"order_id"  db:"order_id"`
	ProductId  int64   `json:"product_id"  db:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"  db:"total_price"`
}

type OrderProduct struct {
	OrderProductTable
	ProductName        string `json:"product_name"  db:"product_name"`
	ProductDescription string `json:"product_description"  db:"product_description"`
	ProductImage       string `json:"product_image"  db:"product_image"`
}
