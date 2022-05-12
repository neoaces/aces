package main

import (
	"github.com/gin-gonic/gin"
	"github.com/neoaces/aces/controller"
)

func main() {
	//
	// INITIALIZING GIN
	// Set up a default server for Gin.
	r := gin.Default()

	// Returns an initialized sqlite database
	controller.NewDb()

	//
	// ROUTING - through Gin
	//
	r.GET("/", controller.GetRandCard)
	r.GET("/:ID", controller.GetCard)
	// r.POST("/api", db.add)

	r.Run() //  Run the gin client
}
