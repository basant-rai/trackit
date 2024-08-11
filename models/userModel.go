package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `json:"username"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Country     string `json:"country"`
}
