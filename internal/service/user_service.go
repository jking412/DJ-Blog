package service

import (
	"DJ-Blog/internal/model"
	"DJ-Blog/pkg/database"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type User struct {
	model.UserModel `json:"user,omitempty"`
	Articles        []Article `json:"articles,omitempty" gorm:"foreignKey:UserId"`
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

func GetUserById(id uint32) (*User, bool) {
	user := &User{}
	if err := database.DB.Model(&model.UserModel{}).Where("id = ?", id).First(user).Error; err != nil {
		return nil, false
	}
	return user, true
}

func GetUserByIdWithArticle(id uint32) (*User, bool) {

	user := &User{}
	if err := database.DB.Model(&model.UserModel{}).Where("id = ?", id).First(user).Error; err != nil {
		return nil, false
	}

	if err := database.DB.Model(&model.ArticleModel{}).Where("user_id = ?", id).Find(&user.Articles).Error; err != nil {
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

func (u *User) String() string {
	// toString
	userStr := fmt.Sprintf("User{Id:%d, Username:%s, Password:%s, CreatedAt:%s, UpdatedAt:%s, AvatarUrl:%s}",
		u.Id,
		u.Username,
		u.Password,
		u.CreatedAt,
		u.UpdatedAt,
		u.AvatarUrl)

	articleStr := ""
	for _, article := range u.Articles {
		articleStr += article.String()
	}

	return userStr + articleStr + "\n"
}
