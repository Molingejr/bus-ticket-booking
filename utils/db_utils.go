package utils

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Or your specific database dialect
)

// Global DB variable
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	// Set up the database connection (SQLite in this case)
	DB, err = gorm.Open("sqlite3", "bus_ticket_booking.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}
