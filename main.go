package main

import (
	"go-app/controllers"
	initializers "go-app/init"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/register", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default to port 8000 if not set
	}
	r.Run(":" + port)
}
