package controllers

import (
	"net/http"
	"table-reservation/config"
	"table-reservation/models"

	"github.com/gin-gonic/gin"
)

// ดึงข้อมูลโต๊ะทั้งหมด
func GetTables(c *gin.Context) {
	var tables []models.Table
	config.DB.Find(&tables)
	c.JSON(http.StatusOK, tables)
}

// เพิ่มโต๊ะใหม่
func AddTable(c *gin.Context) {
	var table models.Table
	if err := c.ShouldBindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&table)
	c.JSON(http.StatusOK, gin.H{"message": "Table added successfully!", "data": table})
}
