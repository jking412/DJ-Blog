package service

import (
	"DJ-Blog/internal/model"
	"DJ-Blog/pkg/database"
)

type UserService struct {
}

func NewUserService() *UserService {
	return new(UserService)
}

type User struct {
	model.UserModel
}

func (u *UserService) GetUserById(id int32) (*User, bool) {
	user := &User{}
	if err := database.DB.Model(&model.UserModel{}).Where("id = ?", id).First(user).Error; err != nil {
		return nil, false
	}
	return user, true
}

func (u *UserService) GetUserByUsername(username string) (*User, bool) {
	user := &User{}
	if err := database.DB.Model(&model.UserModel{}).Where("username = ?", username).First(user).Error; err != nil {
		return nil, false
	}
	return user, true
}
