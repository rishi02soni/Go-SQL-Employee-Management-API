package main

import (
	"employee-api/database"
	"employee-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	router.POST("/employees", handlers.CreateEmployee)
	router.GET("/employees", handlers.GetEmployees)
	router.GET("/employees/:id", handlers.GetEmployee)
	router.PUT("/employees/:id", handlers.UpdateEmployee)
	router.DELETE("/employees/:id", handlers.DeleteEmployee)

	router.Run(":8080")
}
