package main

import (
	"log"
	"net/http"

	"compralibre.com/cobros/model"
	"compralibre.com/cobros/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the "payment" endpoint
	router.POST("/order", calculateOrder)
	router.POST("/payment", paymentHandler)

	// Start the server on port 8080
	router.Run(":8080")
}

func calculateOrder(c *gin.Context) {
	// Read the JSON payload from the request body
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Calculate total ammount
	totalAmount, err := service.CalculateOrderAmount(order)
	if err != nil {
		log.Print("calculateOrder - CalculateOrderAmount -", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There was an error while claculating your order",
		})
		return
	}
	order.TotalAmount = totalAmount

	// Save order
	err = service.SaveOrder(&order)
	if err != nil {
		log.Print("calculateOrder - SaveOrder -", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There was an error while processing your order",
		})
		return
	}

	// Return a JSON response with a message and the payment amount
	c.JSON(http.StatusOK, gin.H{
		"message": "Order processed succesfully",
		"amount":  order.TotalAmount,
		"orderId": order.Id,
	})
}

func paymentHandler(c *gin.Context) {
	// Read the JSON payload from the request body
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Process the payment
	amount, err := service.ProcessPayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There was an error while processing your payment",
		})
		return
	}

	// Return a JSON response with a message and the payment amount
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment processed successfully",
		"amount":  amount,
	})
}
