package service

import (
	"DJ-Blog/internal/model"
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
)

type User struct {
	model.UserModel
}

func CreateUser(username, password string) bool {
	user := &model.UserModel{
		Username: username,
		Password: password,
	}
	if err := database.DB.Create(user).Error; err != nil {
		logrus.Error("Create user failed", err)
		return false
	}
	return true
}

func GetUserById(id int32) (*User, bool) {
	user := &User{}
	if err := database.DB.Model(&model.UserModel{}).Where("id = ?", id).First(user).Error; err != nil {
		return nil, false
	}
	return user, true
}

func GetUserByUsername(username string) (*User, bool) {
	user := &User{}
	if err := database.DB.Model(&model.UserModel{}).Where("username = ?", username).First(user).Error; err != nil {
		return nil, false
	}
	return user, true
}

func IsExistUser(username string) bool {
	var count int64
	database.DB.Model(&model.UserModel{}).Where("username = ?", username).Count(&count)
	return count > 0
}
