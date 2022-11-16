package service

import "DJ-Blog/internal/model"

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
