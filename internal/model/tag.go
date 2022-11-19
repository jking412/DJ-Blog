package model

import (
	"DJ-Blog/pkg/database"
	"fmt"
	"github.com/sirupsen/logrus"
)

type TagModel struct {
	Id   uint32 `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;type:varchar(255)"`
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

func (t *TagModel) String() string {
	return fmt.Sprintf("TagModel{Id:%d,Name:%s}\n", t.Id, t.Name)
}
