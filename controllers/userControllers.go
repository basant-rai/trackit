package controllers

import (
	initializers "go-app/init"
	"go-app/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body struct {
		UserName    string `json:"username"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Password    string `json:"password"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
		Address     string `json:"address"`
		Country     string `json:"country"`
	}
	c.Bind((&body))

	println(body.FirstName)

	post := models.User{
		FirstName:   body.FirstName,
		UserName:    body.UserName,
		LastName:    body.LastName,
		Email:       body.Email,
		Address:     body.Address,
		Country:     body.Country,
		PhoneNumber: body.PhoneNumber,
		Password:    body.Password, // Make sure to hash the password before storing it
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})

}
