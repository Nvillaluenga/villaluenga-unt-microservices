package model

type Order struct {
	Id          int            `json:"id"`
	Products    []ProductOrder `json:"products" binding:"required"`
	User        string         `json:"user" binding:"required"`
	Status      string         `json:"status"`
	TotalAmount float64        `json:"totalAmount"`
}

type ProductOrder struct {
	ProductId int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type Product struct {
	Id          int     `json:"id" binding:"required"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type Payment struct {
	OrderId       int    `json:"order_id" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	User          string `json:"user" binding:"required"`
}
