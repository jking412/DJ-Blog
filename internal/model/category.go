package model

import (
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
)

type CategoryModel struct {
	Id   uint32 `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name"`
}

func (c *CategoryModel) TableName() string {
	return "article_category"
}

func (c *CategoryModel) Create() bool {
	if err := database.DB.Create(c).Error; err != nil {
		logrus.Warn("Create category failed", c)
		return false
	}
	return true
}

func (c *CategoryModel) Delete() bool {
	if err := database.DB.Create(c).Error; err != nil {
		logrus.Warn("Delete category failed", c)
		return false
	}
	return true
}

func (c *CategoryModel) Update() bool {
	if err := database.DB.Updates(c).Error; err != nil {
		logrus.Warn("Update category failed", c)
		return false
	}
	return true
}