package main

import (
	"log"
	"net/http"
	"strconv"

	"compralibre.com/cobros/model"
	"compralibre.com/cobros/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the "payment" endpoint
	router.POST("/order", calculateOrder)
	router.GET("/order/:id", getOrder)
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
			"error": err.Error(),
		})
		return
	}
	order.TotalAmount = totalAmount

	// Save order
	err = service.SaveOrder(&order)
	if err != nil {
		log.Print("calculateOrder - SaveOrder -", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
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

func getOrder(c *gin.Context) {
	orderId, _ := strconv.Atoi(c.Param("id"))
	order, _ := service.GetOrder(orderId)
	c.JSON(http.StatusBadRequest, order)
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
			"error": err.Error(),
		})
		return
	}

	// Return a JSON response with a message and the payment amount
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment processed successfully",
		"amount":  amount,
	})
}
