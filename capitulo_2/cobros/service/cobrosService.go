package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"compralibre.com/cobros/model"
	"compralibre.com/cobros/repository"
)

const (
	DIMINISH_STOCK_URL = "http://localhost:8081/stock/diminish/%d/%d"
	STOCK_URL          = "http://localhost:8081/stock/%d"
)

func CalculateOrderAmount(order model.Order) (float64, error) {
	var totalAmount float64 = 0
	for _, product := range order.Products {
		productId := product.ProductId
		quantity := product.Quantity
		price, err1 := getPrice(productId)
		if err1 != nil {
			log.Print("CalculateOrderAmount - getPrice -", err1)
			return 0, err1
		}
		amount := float64(quantity) * price
		totalAmount = totalAmount + amount
	}
	return totalAmount, nil
}

func SaveOrder(order *model.Order) error {
	_, err := repository.InsertNewOrder(order)
	return err
}

func GetOrder(orderId int) (*model.Order, error) {
	order, err := repository.GetOrder(orderId)
	return order, err
}

func ProcessPayment(payment model.Payment) (float64, error) {
	order, err := repository.GetOrder(payment.OrderId)
	if err != nil {
		return 0, err
	}

	if order.Status != "PENDING" {
		return 0, errors.New("the order is no longer available")
	}

	for _, product := range order.Products {
		productId := product.ProductId
		quantity := product.Quantity
		callSuccesful, err1 := diminishStock(productId, quantity)
		if err1 != nil {
			return 0, err1
		} else if !callSuccesful {
			return 0, errors.New("out of stock")
		}
	}

	order, err = repository.UpdateOrderStatus(payment.OrderId, "DONE")
	if err != nil {
		return 0, err
	}
	return order.TotalAmount, nil
}

func getPrice(productId int) (float64, error) {
	uri := fmt.Sprintf(STOCK_URL, productId)
	resp, err := http.Get(uri)
	if err != nil {
		log.Print("getPrice - POST call -", err)
		return -1, err
	}

	defer resp.Body.Close()
	var data model.StockResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Print("getPrice - decode -", err)
		return -1, err
	}

	return float64(data.Price), nil
}

func diminishStock(productId int, quantity int) (bool, error) {
	uri := fmt.Sprintf(DIMINISH_STOCK_URL, productId, quantity)
	resp, err := http.Post(uri, "application/json", http.NoBody)
	if err != nil {
		return false, err
	}
	if resp.StatusCode == http.StatusConflict {
		return false, nil
	}

	return true, nil
}
