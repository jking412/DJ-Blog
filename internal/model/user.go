package model

import (
	"DJ-Blog/pkg/database"
	"DJ-Blog/pkg/encrypt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	Id        uint32    `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	AvatarUrl string    `gorm:"column:avatar_url"`
}

func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserModel) Create() bool {
	if err := database.DB.Create(u).Error; err != nil {
		logrus.Warn("Create user failed", u)
		return false
	}
	return true
}

func (u *UserModel) Delete() bool {
	if err := database.DB.Create(u).Error; err != nil {
		logrus.Warn("Delete user failed", u)
		return false
	}
	return true
}

func (u *UserModel) BeforeSave(tx *gorm.DB) error {
	if encrypt.IsEncrypt(u.Password) {
		return nil
	}

	u.Password = encrypt.EncryptPassword(u.Password)
	return nil
}
