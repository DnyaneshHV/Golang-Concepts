package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Map acts as our in-memory database
var customers = make(map[int]Customer)  // declaration of the map , as key is int and type of data we are storing in the value is Customer form the struct


var nextID = 1

func main() {
	router := gin.Default()

	router.POST("/customers", createCustomer)
	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getCustomer)
	router.PUT("/customers/:id", updateCustomer)
	router.DELETE("/customers/:id", deleteCustomer)

	router.Run(":8080")
}

// CREATE
func createCustomer(c *gin.Context) {
	var customer Customer  // variable of type Customer that we are passing in the map

	if err := c.ShouldBindJSON(&customer); err != nil {   // we are binding the data which stroed in c in the customer variable and then passing in the map
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer.ID = nextID  // the first customer created will get the id 1
	nextID++

	customers[customer.ID] = customer

	c.JSON(http.StatusCreated, customer)
}

// READ ALL
func getCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, customers)
}

// READ ONE
func getCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	customer, exists := customers[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// UPDATE
func updateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	customer, exists := customers[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
		return
	}

	var updated Customer

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer.Name = updated.Name
	customer.Email = updated.Email

	customers[id] = customer

	c.JSON(http.StatusOK, customer)
}

// DELETE
func deleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	_, exists := customers[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
		return
	}

	delete(customers, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted successfully",
	})
}