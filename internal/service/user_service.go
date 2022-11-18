package service

import (
	"DJ-Blog/internal/model"
	"DJ-Blog/pkg/database"
	"github.com/sirupsen/logrus"
	"time"
)

type User struct {
	model.UserModel `json:"user,omitempty"`
}

func CreateUser(user *model.UserModel) *User {

	//if err := database.DB.Create(user).Error; err != nil {
	//	logrus.Error("Create user failed", err)
	//	return false
	//}

	err := database.DB.Exec("INSERT INTO user (username, password, created_at, updated_at, avatar_url) VALUES (?, ?, ?, ?, ?)",
		user.Username,
		user.Password,
		time.Now(),
		time.Now(),
		user.AvatarUrl).Error

	serviceUser := &User{
		UserModel: *user,
	}

	if err != nil {
		logrus.Error("Create user failed", err)
		return serviceUser
	}

	database.DB.Raw("SELECT LAST_INSERT_ID()").Scan(&serviceUser.Id)

	return serviceUser
}

func DeleteUserByUsername(username string) bool {
	if err := database.DB.Where("username = ?", username).Delete(&model.UserModel{}).Error; err != nil {
		logrus.Error("Delete user failed", err)
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
