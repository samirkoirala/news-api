package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/news-api/services"
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

	createdNews, err := services.CreateNews(news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create news"})
		return
	}

	c.JSON(http.StatusCreated, createdNews)
}
