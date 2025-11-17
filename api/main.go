package main

import (
	"github.com/Acstrayer/TESCSE-Ecom/api/controllers"
	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// For Jess; from Austin: excuse me?
func poop(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("💩"))
}

func main() {
	// Initialize the router and static routes
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/api/", poop)
	r.GET("/api/poop", poop)

	//Product endpoints
	r.GET("/api/products", controllers.GetProducts)
	r.GET("/api/products/search", controllers.SearchProducts)
	r.POST("/api/products/new", controllers.CreateProduct)
	r.GET("/api/product/:id", controllers.FindProduct)
	r.PATCH("/api/product/:id/update", controllers.UpdateProduct)
	r.DELETE("/api/product/:id/delete", controllers.DeleteProduct)

	//Order endpoints
	r.GET("/api/orders", controllers.GetOrders)
	r.GET("/api/orders/search", controllers.SearchOrders)
	r.POST("/api/orders/new", controllers.CreateOrder)

	//User Endpoints
	r.POST("/api/users/new", controllers.CreateUser)
	r.POST("/api/user/login", controllers.UserLogin)

	err := r.Run("127.0.0.1:8079")
	models.Check(err, "Error starting API")
}
