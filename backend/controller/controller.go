package controller

import (
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
	db.AutoMigrate(&models.Card{})
	return &UserDatabase{DB: db}
}

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

func (db *UserDatabase) ReturnCards(ctx *gin.Context) {
	// Finds all the cards in the database
	uncarded := db.DB.Find(&models.Card{}).Error

	ctx.JSON(http.StatusAccepted, gin.H{"cards": uncarded})
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
