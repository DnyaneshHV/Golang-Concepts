package main

import (
	"net/http"   // used for http status code
	"strconv"    // this package is used to convert the string in to numbers, as url parameters are always in string.

	"github.com/gin-gonic/gin"  // So gin is the web framework , we got that package.
)

var customers []Customer
var nextID = 1

func main() {
	router := gin.Default()  // this creates the web server.

	router.POST("/customers", createCustomer)  // calling the function here.
	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getCustomer)
	router.PUT("/customers/:id", updateCustomer)
	router.DELETE("/customers/:id", deleteCustomer)

	router.Run(":8080")   // and running on the port 8080
}



// The Functions that performs the CRUD operation

// gin is package and context is the type of that package.

// CREATE
func createCustomer(c *gin.Context) {  // in gin.context , the data come from the fronted and passesd in c , it may be data, head , info
	var customer Customer   // this creates the empty struct.

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})  // If we don't pass even one of the value in the json it will give and error and return.
		return
	}

	customer.ID = nextID // customer id will be incremented here.
	nextID++

	customers = append(customers, customer) append adds the custmer data into the slice.

	c.JSON(http.StatusCreated, customer) // Context type has the method called JSON , that we are using here.
}

// READ ALL
func getCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, customers)  // Reads all the customers in the slice
}

// READ ONE
func getCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, customer := range customers {
		if customer.ID == id {
			c.JSON(http.StatusOK, customer)
			return
		}  // Using for loop we iterate over each customer and when the id is matched we return that.
	}

	c.JSON(http.StatusNotFound, gin.H{  // And if the id is not found then we, return customer not found error.
		"message": "customer not found",
	})
}

// UPDATE
func updateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updated Customer

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, customer := range customers {
		if customer.ID == id {
			customers[i].Name = updated.Name
			customers[i].Email = updated.Email

			c.JSON(http.StatusOK, customers[i])
			return
		}  // Using the for loop matching the id and updating the values.
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "customer not found",
	})
}

// DELETE
func deleteCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, customer := range customers {
		if customer.ID == id {
			customers = append(customers[:i], customers[i+1:]...) //Remocing the cutomer with the matching id in deleting operation.

			c.JSON(http.StatusOK, gin.H{
				"message": "customer deleted",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{     // if incorrect id passed , then customer not found will display.
		"message": "customer not found",
	})
}
