package main

import (
	"github.com/gin-gonic/gin"
	"github.com/neoaces/aces/controller"
)

func main() {
	r := gin.Default()

	// Returns the initialized sqlite database
	db := controller.NewDb()

	r.GET("/", db.GetCard)
	// r.GET("/:cardID", db.)
	// r.POST("/api", db.add)

	r.Run()
}
