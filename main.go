package main

import (
	"net/http"
	"simpleapi_go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/product/:code", getProduct)
	router.POST("/products", addProduct)
	router.Run("localhost:8083")
}

func getProducts(c *gin.Context) {
	products := models.GetProducts()

	if products == nil || len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}

//c.Param() is used to check/get the whole data structure using one of the parameters we defined in JSON format in our data structure
func getProduct(c *gin.Context) {
	code := c.Param("code")

	product := models.GetProduct(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func addProduct(c *gin.Context) {
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}
}
