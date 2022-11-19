package service

import (
	"DJ-Blog/internal/model"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestCreateArticle(t *testing.T) {

	article := &Article{
		ArticleModel: &model.ArticleModel{
			Title:         faker.Sentence(),
			OriginContent: faker.Paragraph(),
			ParseContent:  faker.Paragraph(),
			UserId:        1,
		},
		Tags: []model.TagModel{
			{
				Id: 1,
			},
			{
				Id: 2,
			},
		},
		Categories: []model.CategoryModel{
			{
				Id: 1,
			},
			{
				Id: 2,
			},
		},
	}

	serviceArticle := CreateArticle(article)

	if serviceArticle.Id == 0 {
		t.Error("创建文章失败")
	}

	fmt.Println(serviceArticle)
}

func TestDeleteArticleById(t *testing.T) {

	if !DeleteArticleById(7) {
		t.Error("删除文章失败")
	}
}

func TestGetArticleList(t *testing.T) {

	articles, ok := GetArticleList()
	if !ok {
		t.Error("获取文章列表失败")
	}
	fmt.Println(articles)
}

func TestCreateTag(t *testing.T) {

	tag := &model.TagModel{
		Name: faker.Word(),
	}

	if !CreateTag(tag) {
		t.Error("创建标签失败")
	}
	fmt.Println(tag)
}

func TestGetTags(t *testing.T) {

	tags, ok := GetTags()
	if !ok {
		t.Error("获取标签失败")
	}
	fmt.Println(tags)
}

func TestCreateCategory(t *testing.T) {

	category := &model.CategoryModel{
		Name: faker.Word(),
	}

	if !CreateCategory(category) {
		t.Error("创建分类失败")
	}
	fmt.Println(category)
}

func TestGetCategories(t *testing.T) {

	categories, ok := GetCategories()
	if !ok {
		t.Error("获取分类失败")
	}
	fmt.Println(categories)
}
