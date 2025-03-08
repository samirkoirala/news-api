package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samirkoirala/news-api/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/news", controllers.GetAllNews)
	r.POST("/newsadd", controllers.CreateNews)
}
