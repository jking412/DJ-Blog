package model

import (
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
	"time"
)

type ArticleModel struct {
	Id            uint32    `gorm:"column:id"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	Title         string    `gorm:"column:title"`
	OriginContent string    `gorm:"column:origin_content"`
	ParseContent  string    `gorm:"column:parse_content"`
	ImgUrl        string    `gorm:"column:img_url"`
	UserId        uint32    `gorm:"column:user_id"`
}

func (a *ArticleModel) TableName() string {
	return "article"
}

func (a *ArticleModel) Create() bool {
	if err := database.DB.Create(a).Error; err != nil {
		logrus.Warn("Create article failed", a)
		return false
	}
	return true
}

func (a *ArticleModel) Delete() bool {
	if err := database.DB.Create(a).Error; err != nil {
		logrus.Warn("Delete article failed", a)
		return false
	}
	return true
}

func (a *ArticleModel) Update() bool {
	if err := database.DB.Updates(a).Error; err != nil {
		logrus.Warn("Update article failed", a)
		return false
	}
	return true
}
