package service

import (
	"DJ-Blog/internal/model"
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
)

// Article 是一个服务层的结构体
// 用于处理业务逻辑，也用于处理和返回数据
type Article struct {
	// 当model.ArticleModel改为*model.ArticleModel时，gorm会有奇怪的问题
	// example: 当调用GetArticleList()时，假如此时有10条数据，那么返回的10条数据会一模一样
	model.ArticleModel `json:"article,omitempty"`
	Tags               []model.TagModel      `json:"tags" gorm:"many2many:article_tag_relation"`
	Categories         []model.CategoryModel `json:"categories" gorm:"many2many:article_category_relation"`
}

func CreateArticle(article *Article) *Article {

	// 当tag和category的Id不存在时，gorm会创建name为空的tag和category
	if err := database.DB.Create(article).Error; err != nil {
		logrus.Warn("Create article failed", err)
		return nil
	}

	return article
}

func DeleteArticleById(id uint32) bool {

	if err := database.DB.Model(&Article{}).Where("id = ?", id).Delete(&model.ArticleModel{}).Error; err != nil {
		logrus.Warn("Delete article failed", err)
		return false
	}

	return true
}

func GetArticleList() ([]Article, bool) {

	articles := make([]Article, 0)
	if err := database.DB.Model(&Article{}).
		Preload("Tags").
		Preload("Categories").
		// 这里的articles如果前面不加 &会报错
		Find(&articles).Error; err != nil {
		logrus.Warn("Get article list failed", err)
		return nil, false
	}

	return articles, true
}

func GetArticleOrderByTime() ([]Article, bool) {

	var articles []Article
	if err := database.DB.Model(&Article{}).
		Preload("Tags").
		Preload("Categories").
		Order("created_at desc").
		Find(&articles).Error; err != nil {
		logrus.Warn("Get article list failed", err)
		return nil, false
	}

	return articles, true
}

func GetArticleById(id uint32) (*Article, bool) {

	var article *Article
	if err := database.DB.Model(&Article{}).Where("id = ?", id).
		Preload("Tags").
		Preload("Categories").
		// 这里的articles如果前面不加 &会报错
		First(&article).Error; err != nil {
		logrus.Warn("Get article failed ", err)
		return nil, false
	}
	return article, true
}

func GetArticleByTitle(title string) (*Article, bool) {

	var article *Article
	if err := database.DB.Model(&model.ArticleModel{}).Where("title = ?", title).
		Preload("Tags").
		Preload("Categories").
		First(article).Error; err != nil {
		logrus.Warn("Get article failed", err)
		return nil, false
	}
	return article, true
}

// TODO: 下面的service层代码都待修改

func GetArticleByTagId(tagId uint32) ([]Article, bool) {

	var articles []Article
	if err := database.DB.Model(&Article{}).
		Preload("Tags").
		Preload("Categories").
		Find(&articles).Error; err != nil {
		logrus.Warn("Get article list failed", err)
		return nil, false
	}

	var newArticles []Article
	for _, article := range articles {
		for _, tag := range article.Tags {
			if tag.Id == tagId {
				newArticles = append(newArticles, article)
			}
		}
	}

	return newArticles, true
}

func GetArticleByCategoryId(categoryId uint32) ([]Article, bool) {

	var articles []Article
	if err := database.DB.Model(&Article{}).
		Joins("join article_category_relation on category_id = ?", categoryId).
		Where("article.id = article_id").
		Find(articles).Error; err != nil {
		logrus.Warn("Get article list failed", err)
		return nil, false
	}

	return articles, true
}

func UpdateArticle(article *Article) bool {

	if err := database.DB.Model(&Article{}).Where("id = ?", article.Id).Updates(article).Error; err != nil {
		logrus.Warn("Update article failed", err)
		return false
	}

	return true
}

func CreateTag(tag *model.TagModel) bool {

	if err := database.DB.Create(tag).Error; err != nil {
		logrus.Warn("Create tag failed", err)
		return false
	}
	return true
}

func GetTags() ([]model.TagModel, bool) {

	var tags []model.TagModel
	if err := database.DB.Model(&model.TagModel{}).Find(&tags).Error; err != nil {
		logrus.Warn("Get tags failed", err)
		return tags, false
	}
	return tags, true
}

func CreateCategory(category *model.CategoryModel) bool {

	if err := database.DB.Create(category).Error; err != nil {
		logrus.Warn("Create category failed", err)
		return false
	}
	return true
}

func GetCategories() ([]model.CategoryModel, bool) {

	var categories []model.CategoryModel
	if err := database.DB.Model(&model.CategoryModel{}).Find(&categories).Error; err != nil {
		logrus.Warn("Get categories failed", err)
		return categories, false
	}
	return categories, true
}

func IsExistArticleById(id uint32) bool {

	var count int64
	if err := database.DB.Model(&Article{}).Where("id = ?", id).Count(&count).Error; err != nil {
		logrus.Warn("Get article failed", err)
		return false
	}
	return count > 0
}

func (a *Article) String() string {

	articleStr := a.ArticleModel.String()
	tagsStr := "\n"
	for _, tag := range a.Tags {
		tagsStr += tag.String()
	}
	categoriesStr := "\n"
	for _, category := range a.Categories {
		categoriesStr += category.String()
	}

	return articleStr + tagsStr + categoriesStr + "\n"
}
