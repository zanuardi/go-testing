package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zanuardinovanda/go-testing/controllers"
	"github.com/zanuardinovanda/go-testing/models"
)

func main() {

	route := gin.Default()

	//Model
	db := models.SetupModels()
	route.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to my API"})
	})

	//User
	route.GET("/user", controllers.GetUser)
	route.PUT("/user/:id", controllers.EditUser)
	route.DELETE("/user/:id", controllers.DeleteUser)

	//Register
	route.POST("/register", controllers.RegisterUser)

	route.Run()
}
