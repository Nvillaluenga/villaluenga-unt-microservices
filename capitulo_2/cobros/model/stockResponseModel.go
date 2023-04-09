package model

type StockResponse struct {
	ProductId int     `json:"product_id" binding:"required"`
	Stock     int     `json:"stock"  binding:"required"`
	Price     float64 `json:"price"  binding:"required"`
}
