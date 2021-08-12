package models

import "gorm.io/gorm"

//Product -> model for products table
type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

//TableName --> Table for Product Model
func (Product) TableName() string {
	return "products"
}
