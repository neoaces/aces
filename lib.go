package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

var cards []string

func getCard(ctx *gin.Context) {
	cardID := ctx.Param("cardID")

	if index, err := strconv.Atoi(cardID); err == nil {
		ctx.JSON(http.StatusOK, &cards[index])
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong card."})

	}
}

func returnCards(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{"cards": cards})
}

func addCards(ctx *gin.Context) {
	item := ctx.PostForm("item")
	cards = append(cards, item)

	ctx.String(http.StatusCreated, ctx.FullPath())
}

func main() {
	r := gin.Default()
	cards = append(cards, "New Card")

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(200, "Welcome back.")
	// })

	r.GET("/", returnCards)
	r.GET("/:cardID", getCard)
	r.POST("/api", addCards)

	r.Run()
}