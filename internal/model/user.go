package model

import "time"

type UserModel struct {
	Id        uint32    `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	AvatarUrl string    `gorm:"column:avatar_url"`
}

func (UserModel) TableName() string {
	return "user"
}
