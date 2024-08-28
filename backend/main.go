package main

import (
	"menu/api"
	"menu/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.GetDB()
	// Start monitoring the database connection in a separate goroutine
	// go database.MonitorConnection()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/Drinks", api.PostDrink)
	r.GET("/Drinks", api.GetDrinkList)
	r.PUT("/Drinks/:id", api.PutDrink)
	r.DELETE("/Drinks/:id", api.DeleteDrink)

	r.Run()
}
