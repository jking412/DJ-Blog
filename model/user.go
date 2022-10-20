package model

import (
	"DJ-Blog/pkg/database"
	"DJ-Blog/pkg/encrypt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username  string `gorm:"column:username;type:varchar(255);not null"`
	Password  string `gorm:"column:password;type:varchar(255);default:''"`
	Email     string `gorm:"column:email;type:varchar(255);default:''"`
	AvatarUrl string `gorm:"column:avatar_url;type:varchar(255);default:''"`
	Posts     []Post `gorm:"foreignKey:UserId"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	if err := database.DB.Create(u).Error; err != nil {
		logrus.Error("User create failed")
		return err
	}
	return nil
}

func (u *User) GetUserById() (*User, error) {
	var user *User
	if err := database.DB.Where("id = ?", u.Id).First(&user).Error; err != nil {
		logrus.Warn("Get user by id failed")
		return user, err
	}
	return user, nil
}

func (u *User) GetUserByUsername() (*User, error) {
	var user *User
	err := database.DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		logrus.Warn("Get user by username failed")
	}
	return user, err
}

func (u *User) IsExistUsername() bool {
	var user User
	if err := database.DB.Where("username = ?", u.Username).First(&user).Error; err != nil {
		return false
	}
	return true
}

func (u *User) CheckPassword() bool {
	var user *User
	if err := database.DB.Where("username = ?", u.Username).First(&user).Error; err != nil {
		logrus.Warn("Get user by username failed")
		return false
	}
	return encrypt.CheckPassword(u.Password, user.Password)
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password == "" {
		return nil
	}
	if !encrypt.IsEncrypt(u.Password) {
		u.Password = encrypt.EncryptPassword(u.Password)
	}
	return nil
}
