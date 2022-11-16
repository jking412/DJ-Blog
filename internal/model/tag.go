package model

import (
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
)

type TagModel struct {
	Id   uint32 `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name"`
}

func (t *TagModel) TableName() string {
	return "article_tag"
}

func (t *TagModel) Create() bool {
	if err := database.DB.Create(t).Error; err != nil {
		logrus.Warn("Create tag failed", t)
		return false
	}
	return true
}

func (t *TagModel) Delete() bool {
	if err := database.DB.Create(t).Error; err != nil {
		logrus.Warn("Delete tag failed", t)
		return false
	}
	return true
}

func (t *TagModel) Update() bool {
	if err := database.DB.Updates(t).Error; err != nil {
		logrus.Warn("Update tag failed", t)
		return false
	}
	return true
}