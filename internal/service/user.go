package service

import (
	"DJ-Blog/internal/model"
	"DJ-Blog/pkg/database"
)

type User struct {
	model.UserModel
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
