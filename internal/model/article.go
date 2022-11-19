package model

import (
	"DJ-Blog/pkg/database"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type ArticleModel struct {
	Id            uint32    `gorm:"column:id" json:"id,omitempty"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updatedAt"`
	Title         string    `gorm:"column:title;type:varchar(255);index:,class:FULLTEXT" json:"title,omitempty"`
	OriginContent string    `gorm:"column:origin_content;type:longtext;index:,class:FULLTEXT" json:"originContent,omitempty"`
	ParseContent  string    `gorm:"column:parse_content;type:longtext" json:"parseContent,omitempty"`
	ImgUrl        string    `gorm:"column:img_url;type:varchar(255)" json:"imgUrl,omitempty"`
	UserId        uint32    `gorm:"column:user_id" json:"userId,omitempty"`
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

func (a *ArticleModel) String() string {
	return fmt.Sprintf("ArticleModel{Id:%d, CreatedAt:%s, UpdatedAt:%s, Title:%s, OriginContent:%s, ParseContent:%s, ImgUrl:%s, UserId:%d}\n",
		a.Id,
		a.CreatedAt,
		a.UpdatedAt,
		a.Title,
		a.OriginContent,
		a.ParseContent,
		a.ImgUrl,
		a.UserId)
}
