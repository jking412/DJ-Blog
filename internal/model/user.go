package model

import "time"

type UserModel struct {
	Id        uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
	Username  string    `gorm:"column:username;type:varchar(100);not null"`
	Password  string    `gorm:"column:password;type:varchar(100);not null"`
	AvatarUrl string    `gorm:"column:avatar_url;type:varchar(255);default:''"`
}

func (UserModel) TableName() string {
	return "user"
}
