package services

import (
	"github.com/samirkoirala/news-api/config"
)

func FetchAllNews() ([]config.News, error) {
	var news []config.News
	if err := config.DB.Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func CreateNews(news config.News) (config.News, error) {
	if err := config.DB.Create(&news).Error; err != nil {
		return config.News{}, err
	}
	return news, nil
}
