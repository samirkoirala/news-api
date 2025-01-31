package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/news-api/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/news", controllers.GetAllNews)
	r.POST("/news", controllers.CreateNews)
}
