package controllers

import (
	"net/http"
	"table-reservation/config"
	"table-reservation/models"

	"github.com/gin-gonic/gin"
)

func ReserveTable(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&reservation)
	c.JSON(http.StatusOK, gin.H{"message": "Table reserved"})
}
