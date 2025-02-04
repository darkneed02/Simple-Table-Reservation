package routes

import (
	"table-reservation/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/tables", controllers.GetTables)
	r.POST("/tables", controllers.AddTable) // เพิ่ม API เพิ่มโต๊ะ
	r.POST("/reserve", controllers.ReserveTable)
}
