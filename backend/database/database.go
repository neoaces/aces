package database

// This package implements all the functions required to connect
// and initialize the database, NOT ACCESS THE DATABASE DIRECTLY.

import (
	"github.com/neoaces/aces/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// dsn := "postgres://admin:admin@localhost:5432/gorm"
	dsn := "postgres://admin:admin@psql:5432/gorm"
	// Initialize the database through the gorm sqlite driver.
	// Open a connection through gorm
	// db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})  sqlite connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Something went wrong with the database connection")
	}

	// Make migrations from the sqlite database
	db.AutoMigrate(&models.Card{}, &models.CardSet{}, &models.User{})

	// Return the pointer to the database connection, provided by GORM.
	return db
}

func AddCard(db *gorm.DB, card *models.Card) (err error) {
	db.Save(card)
	return nil
}
