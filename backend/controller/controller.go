package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/neoaces/aces/database"
	"github.com/neoaces/aces/models"
	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

func NewDb() *UserDatabase {
	db := database.InitDB()
	// fmt.Printf("%T, %v\n", db, db)

	// Whenever the table or model changes, the database
	// will automatically update (without changing the data)
	db.AutoMigrate(&models.Card{})

	// Note how the function DOES NOT return a pointer to the database. Without this
	// typing, and using db.DB to access the new database, an error is thrown.
	return &UserDatabase{DB: db}
}

// GET "" - fetches a specific card
func (db *UserDatabase) GetCard(ctx *gin.Context) {
	cardID := ctx.Param("cardID")

	if index, err := strconv.Atoi(cardID); err == nil {
		var card models.Card
		db.DB.First(&card, index)
		ctx.JSON(http.StatusOK, &card.Answer)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No card under this index."})
	}
}

// GET "/" - fetches a random user
func (db *UserDatabase) ReturnCards(ctx *gin.Context) {
	found_cards := []models.Card{}
	// Finds all the cards in the database
	db.DB.Find(&found_cards)

	for _, card := range found_cards {
		fmt.Printf("The card is %v, and the answer is %v\n", card.Name, card.Answer)
	}

	ctx.JSON(http.StatusFound, gin.H{"cards": found_cards})
}

func (db *UserDatabase) AddCards(ctx *gin.Context) {
	var json models.Card
	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	database.AddCard(db.DB, &json)

	ctx.String(http.StatusCreated, ctx.FullPath())
}
