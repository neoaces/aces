package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var cards []string

func returnCards(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{"cards": cards})
}

func main() {
	r := gin.Default()
	cards = append(cards, "New Card")

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(200, "Welcome back.")
	// })

	r.GET("/", returnCards)

	r.Run()
}