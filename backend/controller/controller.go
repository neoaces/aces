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

var in_db *gorm.DB

// Initialize the database, and register it into the controller file.
func NewDb() *gorm.DB {
	db := database.InitDB()
	// fmt.Printf("%T, %v\n", db, db)

	// Whenever the table or model changes, the database
	// will automatically update (without changing the data)

	// Note how the function DOES return a pointer to the database
	in_db = db
	return db
}

// GET "/" - fetches a random user
func GetRandCard(ctx *gin.Context) {
	found_cards := []models.Card{}
	// Finds all the cards in the database
	in_db.Table("cards").Take(&found_cards)

	for _, card := range found_cards {
		fmt.Printf("The card is %v, and the answer is %v\n", card.Name, card.Answer)
	}

	ctx.JSON(http.StatusFound, gin.H{"cards": found_cards})
}

// GET "/:ID" - fetches a specific card
func GetCard(ctx *gin.Context) {
	cardID := ctx.Param("ID")

	if index, err := strconv.Atoi(cardID); err == nil {
		var card models.Card
		in_db.First(&card, index)
		ctx.JSON(http.StatusOK, &card)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No card under this index."})
	}
}

// func AddCards(ctx *gin.Context) { // TODO: make query create a card in the database
// 	var json models.Card
// 	err := ctx.BindJSON(&json)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}

// 	database.AddCard(in_db, &json)

// 	ctx.String(http.StatusCreated, ctx.FullPath())
// }
