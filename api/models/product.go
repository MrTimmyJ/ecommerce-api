// models/product.go

package models

import (
	"gorm.io/gorm"
	"time"
)

type ProductRequestData struct {
	Products  []Product
	Retrieved time.Time
}

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Quantity    uint   `json:"quantity"`
	Price       int    `json:"price"` //In pennies
	PLU         int    `json:"plu"`
}
