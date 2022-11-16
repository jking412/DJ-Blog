package service

import (
	"DJ-Blog/internal/model"
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
)

type ArticleService struct {
}

type Article struct {
	model.ArticleModel
	tags       []model.TagModel
	categories []model.CategoryModel
}

func NewArticleService() *ArticleService {
	return new(ArticleService)
}

func (a *ArticleService) GetArticleList() ([]Article, bool) {
	articles := make([]Article, 0)
	if err := database.DB.Model(&model.ArticleModel{}).Find(articles).Error; err != nil {
		logrus.Warn("Get article list failed", err)
		return nil, false
	}
	return articles, true
}

func (a *ArticleService) GetArticleById(id uint32) (*Article, bool) {
	article := &Article{}
	if err := database.DB.Model(&model.ArticleModel{}).Where("id = ?", id).First(article).Error; err != nil {
		logrus.Warn("Get article failed", err)
		return nil, false
	}
	return article, true
}

func (a *ArticleService) GetArticleByTitle(title string) (*Article, bool) {
	article := &Article{}
	if err := database.DB.Model(&model.ArticleModel{}).Where("title = ?", title).First(&article).Error; err != nil {
		logrus.Warn("Get article failed", err)
		return nil, false
	}
	return article, true
}
