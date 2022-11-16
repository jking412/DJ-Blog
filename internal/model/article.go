package model

import "time"

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
