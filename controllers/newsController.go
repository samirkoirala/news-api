package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samirkoirala/news-api/config"
	"github.com/samirkoirala/news-api/services"
)

func GetAllNews(c *gin.Context) {
	news, err := services.FetchAllNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch news"})
		return
	}
	c.JSON(http.StatusOK, news)
}

func CreateNews(c *gin.Context) {
	var news services.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	createdNews, err := services.CreateNews(config.News(news))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create news"})
		return
	}

	c.JSON(http.StatusCreated, createdNews)
}
