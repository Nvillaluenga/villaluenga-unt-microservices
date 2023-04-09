package repository

import (
	"errors"

	"compralibre.com/cobros/model"
)

var orders []model.Order

func InitDb() {
	orders = []model.Order{
		{
			Id:          1,
			User:        "user1",
			Status:      "DONE",
			TotalAmount: 100.00,
			Products: []model.ProductOrder{
				{
					ProductId: 1,
					Quantity:  1,
				},
			},
		},
	}
}

func InsertNewOrder(order *model.Order) (*model.Order, error) {
	orderId := GetMaxId() + 1
	order.Id = orderId
	order.Status = "PENDING"
	orders = append(orders, *order)
	return order, nil
}

func UpdateOrderStatus(orderId int, status string) (*model.Order, error) {
	order, err := GetOrder(orderId)
	if err != nil {
		return nil, err
	}
	order.Status = status
	return order, nil
}

func GetOrder(orderId int) (*model.Order, error) {
	for _, order := range orders {
		if order.Id == orderId {
			return &order, nil
		}
	}
	return nil, errors.New("order not found")
}

// Helper functions
func GetMaxId() int {
	maxId := -1
	for _, order := range orders {
		maxId = max(maxId, order.Id)
	}
	return maxId
}

func max(a int, b int) int {
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}
