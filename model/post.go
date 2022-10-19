package model

import (
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
)

type Post struct {
	BaseModel
	Title   string `gorm:"column:title;type:varchar(255);not null"`
	Tag     string `gorm:"column:tag;type:varchar(255);not null"`
	Content string `gorm:"column:content;type:longtext;not null"`
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
