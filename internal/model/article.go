package model

import "time"

type ArticleModel struct {
	Id            uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt     time.Time `gorm:"column:created_at;not null"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null"`
	Title         string    `gorm:"column:title;type:varchar(100);not null"`
	OriginContent string    `gorm:"column:origin_content;type:longtext;not null"`
	ParseContent  string    `gorm:"column:parse_content;type:longtext;not null"`
	ImgUrl        string    `gorm:"column:img_url;type:varchar(255);not null"`
}

func (ArticleModel) TableName() string {
	return "article"
}
