package api

import (
	"menu/internal/database"
	"menu/model/drinks"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDrinkList(c *gin.Context) {
	var drinks []drinks.Drinks
	results := database.GetDB().Find(&drinks)

	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": drinks})
}
