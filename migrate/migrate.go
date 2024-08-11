package main

import (
	initializers "go-app/init"
	"go-app/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate((&models.User{}))
}
