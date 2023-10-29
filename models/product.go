package models

type ProductTable struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

type Product struct {
	ProductTable
}
