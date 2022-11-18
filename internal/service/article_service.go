package service

import (
	"DJ-Blog/internal/model"
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
)

// Article 是一个服务层的结构体
// 用于处理业务逻辑，也用于处理和返回数据
type Article struct {
	*model.ArticleModel `json:"article,omitempty"`
	Tags                []model.TagModel      `json:"tags,omitempty"`
	Categories          []model.CategoryModel `json:"categories,omitempty"`
}

func CreateArticle(article *model.ArticleModel) *Article {

	if err := database.DB.Create(article).Error; err != nil {
		logrus.Warn("Create article failed", err)
		return nil
	}

	serviceArticle := &Article{
		ArticleModel: article,
	}

	database.DB.Raw("SELECT LAST_INSERT_ID()").Scan(&serviceArticle.Id)

	return serviceArticle
}

func DeleteArticleById(id uint32) bool {
	if err := database.DB.Model(&model.ArticleModel{}).Where("id = ?", id).Delete(&model.ArticleModel{}).Error; err != nil {
		logrus.Warn("Delete article failed", err)
		return false
	}
	return true
}

func GetArticleList() ([]Article, bool) {
	articles := make([]Article, 0)
	if err := database.DB.Model(&model.ArticleModel{}).Preload("Tags").
		Preload("Categories").
		Find(articles).Error; err != nil {
		logrus.Warn("Get article list failed", err)
		return nil, false
	}
	return articles, true
}

func GetArticleOrderByTime() ([]Article, bool) {

	var articles []Article
	if err := database.DB.Model(&model.ArticleModel{}).
		Preload("Tags").
		Preload("Categories").
		Order("created_at desc").
		Find(articles).Error; err != nil {
		logrus.Warn("Get article list failed", err)
		return nil, false
	}
	return articles, true
}

func GetArticleById(id uint32) (*Article, bool) {
	article := &Article{}
	if err := database.DB.Model(&model.ArticleModel{}).Where("id = ?", id).
		Preload("Tags").
		Preload("Categories").
		First(article).Error; err != nil {
		logrus.Warn("Get article failed", err)
		return nil, false
	}
	return article, true
}

func GetArticleByTitle(title string) (*Article, bool) {

	article := &Article{}
	if err := database.DB.Model(&model.ArticleModel{}).Where("title = ?", title).
		Preload("Tags").
		Preload("Categories").
		First(&article).Error; err != nil {
		logrus.Warn("Get article failed", err)
		return nil, false
	}
	return article, true
}

func GetArticleByTagId(tagId uint32) ([]Article, bool) {

	var articles []Article
	if err := database.DB.Model(&model.ArticleModel{}).
		Joins("join article_tag_relation on tag_id = ?", tagId).
		Where("article.id = article_id").
		Find(articles).Error; err != nil {
		logrus.Warn("Get article list failed", err)
		return nil, false
	}

	return articles, true
}

func GetArticleByCategoryId(categoryId uint32) ([]Article, bool) {

	var articles []Article
	if err := database.DB.Model(&model.ArticleModel{}).
		Joins("join article_category_relation on category_id = ?", categoryId).
		Where("article.id = article_id").
		Find(articles).Error; err != nil {
		logrus.Warn("Get article list failed", err)
		return nil, false
	}
	return articles, true
}

func GetTags() ([]model.TagModel, bool) {

	var tags []model.TagModel
	if err := database.DB.Model(&model.TagModel{}).Find(&tags).Error; err != nil {
		logrus.Warn("Get tags failed", err)
		return tags, false
	}
	return tags, true
}

func GetCategories() ([]model.CategoryModel, bool) {

	var categories []model.CategoryModel
	if err := database.DB.Model(&model.CategoryModel{}).Find(&categories).Error; err != nil {
		logrus.Warn("Get categories failed", err)
		return categories, false
	}
	return categories, true
}

func UpdateArticle(article *model.ArticleModel) bool {
	if err := database.DB.Model(&model.ArticleModel{}).Where("id = ?", article.Id).Updates(article).Error; err != nil {
		logrus.Warn("Update article failed", err)
		return false
	}
	return true
}
