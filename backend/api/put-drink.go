package api

import (
	"menu/internal/database"
	"menu/model/drinks"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PutDrink(c *gin.Context) {
	var payload drinks.Drinks
	var drink drinks.Drinks

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	result := database.GetDB().First(&drink, id)
	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record Not Found"})
	}

	payload.ID = drink.ID

	database.GetDB().Save(&payload)

	c.JSON(http.StatusAccepted, gin.H{"data": payload})
}
