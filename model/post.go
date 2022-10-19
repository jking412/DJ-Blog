package model

import (
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
)

type Post struct {
	BaseModel
	Title   string `gorm:"column:title;type:varchar(255);not null"`
	Tag     string `gorm:"column:tag;type:varchar(255);default:''"`
	Content string `gorm:"column:content;type:longtext;not null"`
	Likes   uint64 `gorm:"column:likes;default:0"`
	Stared  bool   `gorm:"column:stared;default:false"`
}

func (p *Post) TableName() string {
	return "posts"
}

func (p *Post) GetAllPosts() []Post {
	var posts []Post
	if err := database.DB.Find(&posts).Error; err != nil {
		logrus.Warn("Get all posts failed")
	}
	return posts
}

func (p *Post) GetPostById() *Post {
	var post *Post
	if err := database.DB.Where("id = ?", p.Id).First(&post).Error; err != nil {
		logrus.Warn("Get post by id failed")
	}
	return post
}

func (p *Post) DeletePostById() {
	if err := database.DB.Where("id = ?", p.Id).Delete(p).Error; err != nil {
		logrus.Warn("Delete post by id failed")
	}
}
