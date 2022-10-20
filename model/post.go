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
	UserId  uint64 `gorm:"column:user_id"`
}

func (p *Post) TableName() string {
	return "posts"
}

func (p *Post) GetAllPosts() ([]Post, error) {
	var posts []Post
	if err := database.DB.Find(&posts).Error; err != nil {
		logrus.Warn("Get all posts failed")
		return nil, err
	}
	return posts, nil
}

func (p *Post) Create() error {
	if err := database.DB.Create(p).Error; err != nil {
		logrus.Warn("Create post failed")
		return err
	}
	return nil
}

func (p *Post) GetPostById() (*Post, error) {
	var post *Post
	if err := database.DB.Where("id = ?", p.Id).First(&post).Error; err != nil {
		logrus.Warn("Get post by id failed")
		return nil, err
	}
	return post, nil
}

func (p *Post) Update() error {
	if err := database.DB.Where("id = ?", p.Id).Updates(p).Error; err != nil {
		logrus.Warn("Update post by id failed")
		return err
	}
	return nil
}

func (p *Post) DeletePostById() error {
	if err := database.DB.Where("id = ?", p.Id).Delete(p).Error; err != nil {
		logrus.Warn("Delete post by id failed")
		return err
	}
	return nil
}
