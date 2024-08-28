package endpoints

import (
	"log"
	"menu/internal/database"
	"menu/model/drinks"
	"net/http"

	"github.com/gin-gonic/gin"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Createdrink(c *gin.Context) {
	var data drinks.Drinks

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	drink := drinks.Drinks{Name: data.Name}
	result := database.GetDB().Create(&drink)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	}

	c.JSON(http.StatusOK,
		gin.H{
			"message": "create drink",
			"data":    drink,
		})
}
