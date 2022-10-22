package model

import (
	"DJ-Blog/pkg/database"
	"DJ-Blog/pkg/search"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Post struct {
	BaseModel
	Title   string `gorm:"column:title;type:varchar(255);not null"`
	Tag     string `gorm:"column:tag;type:varchar(255);default:''"`
	Content string `gorm:"column:content;type:longtext;not null"`
	Author  string `gorm:"column:author;type:varchar(255);default:''"`
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

func (p *Post) AfterSave(tx *gorm.DB) (err error) {
	// save post to es
	dataMap := map[string]interface{}{
		"id":         p.Id,
		"title":      p.Title,
		"tag":        p.Tag,
		"content":    p.Content,
		"author":     p.Author,
		"likes":      p.Likes,
		"stared":     p.Stared,
		"created_at": p.CreatedAt,
		"updated_at": p.UpdatedAt,
		"user_id":    p.UserId,
	}
	if ok := search.ZincCli.PutDoc(search.ZincCli.IndexName, p.Id, dataMap); !ok {
		logrus.Warn("Put post to es failed")
		return fmt.Errorf("put post to es failed")
	}
	return nil
}

func (p *Post) AfterUpdate(tx *gorm.DB) (err error) {
	// update post to es
	dataMap := map[string]interface{}{
		"id":         p.Id,
		"title":      p.Title,
		"tag":        p.Tag,
		"content":    p.Content,
		"author":     p.Author,
		"likes":      p.Likes,
		"stared":     p.Stared,
		"created_at": p.CreatedAt,
		"updated_at": p.UpdatedAt,
		"user_id":    p.UserId,
	}
	if ok := search.ZincCli.PutDoc(search.ZincCli.IndexName, p.Id, dataMap); !ok {
		logrus.Warn("Update post to es failed")
		return fmt.Errorf("update post to es failed")
	}
	return nil
}

func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	// delete post from es
	if ok := search.ZincCli.DeleteDoc(search.ZincCli.IndexName, p.Id); !ok {
		logrus.Warn("Delete post from es failed")
		return fmt.Errorf("delete post from es failed")
	}
	return nil
}
