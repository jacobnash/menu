package api

import (
	"menu/internal/database"
	"menu/model/drinks"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostDrink(c *gin.Context) {
	var drink drinks.Drinks

	if err := c.ShouldBindJSON(&drink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := database.GetDB().Create(&drink)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "created new Record", "data": drink})
	return
}
