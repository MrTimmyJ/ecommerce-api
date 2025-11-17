//models/setup.go

package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// Create a shorthand function to check for errors
func Check(e error, m string) {
	if e != nil {
		log.Panic(m+": ", e.Error())
	}
}

func ConnectDatabase() {
	// Connect to the database
	database, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	Check(err, "Database connection error")

	// Migrate product definition to database
	err = database.AutoMigrate(&Product{})
	Check(err, "Product migration error")

	err = database.AutoMigrate(&OrderItem{})
	Check(err, "Order migration error")

	// Migrate order definition to database
	err = database.AutoMigrate(&Order{})
	Check(err, "Order migration error")

	err = database.AutoMigrate(&User{})
	Check(err, "User migration error")

	// Set global database variable
	DB = database
}
