// models/order.go

package models

import (
	"gorm.io/gorm"
	"time"
)

type OrderItem struct {
	gorm.Model
	Quantity  int     `json:"quantity" binding:"required"`
	Price     int     `json:"price"` //in pennies
	Product   Product `json:"product"`
	ProductID uint    `json:"product_id" binding:"required"`
	OrderID   uint    `json:"order_id"`
}

type OrderRequestData struct {
	Orders    []Order
	Retrieved time.Time
}

type Order struct {
	gorm.Model
	Items    []OrderItem `json:"items" binding:"required"`
	Name     string      `json:"name" binding:"required"`
	Email    string      `json:"email" binding:"required"`
	Address1 string      `json:"address_one" binding:"required"`
	Address2 string      `json:"address_two"`
	City     string      `json:"city" binding:"required"`
	State    string      `json:"state" binding:"required"`
	Zip      string      `json:"zip" binding:"required"`
}
