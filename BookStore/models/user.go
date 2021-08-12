package models

import "gorm.io/gorm"

//User -> model for users table
type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//TableName --> Table for User Model
func (User) TableName() string {
	return "users"
}
