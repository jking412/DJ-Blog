package model

import (
	"DJ-Blog/pkg/database"
	"DJ-Blog/pkg/encrypt"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	Id        uint32    `gorm:"column:id;primaryKey;autoIncrement" example:"1"`
	CreatedAt time.Time `gorm:"column:created_at" example:"2021-01-01 00:00:00"`
	UpdatedAt time.Time `gorm:"column:updated_at" example:"2021-01-01 00:00:00"`
	Username  string    `gorm:"column:username;index" example:"admin"`
	Password  string    `gorm:"column:password" example:"123456"`
	AvatarUrl string    `gorm:"column:avatar_url" example:"https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"`
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

func (u *UserModel) String() string {
	return fmt.Sprintf("UserModel{Id:%d, CreatedAt:%s, UpdatedAt:%s, Username:%s, Password:%s, AvatarUrl:%s}\n",
		u.Id,
		u.CreatedAt,
		u.UpdatedAt,
		u.Username,
		u.Password,
		u.AvatarUrl)
}
