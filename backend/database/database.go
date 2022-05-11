package database

// This package implements all the functions required to connect
// and initialize the database, NOT ACCESS THE DATABASE DIRECTLY.

import (
	"github.com/neoaces/aces/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Initialize the database through the gorm sqlite driver.
	// Open a connection through gorm
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("Something went wrong")
	}

	// Make migrations from the sqlite database
	db.AutoMigrate(&models.Card{})

	// Return the pointer to the database connection, provided by GORM.
	return db
}

func AddCard(db *gorm.DB, card *models.Card) (err error) {
	db.Save(card)
	return nil
}
