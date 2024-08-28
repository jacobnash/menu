package api

import (
	"menu/internal/database"
	"menu/model/drinks"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteDrink(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Delete(&drinks.Drinks{}, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
