package service

import (
	"DJ-Blog/internal/model"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestCreateUser(t *testing.T) {

	// 用于构造随机的测试数据
	//user := &model.UserModel{
	//	Username:  faker.Username(),
	//	Password:  faker.Password(),
	//	AvatarUrl: faker.URL(),
	//}

	// 用于测试的用户
	user := &model.UserModel{
		Username:  "admin",
		Password:  "123456",
		AvatarUrl: faker.URL(),
	}

	serviceUser := CreateUser(user)

	if serviceUser.Id == 0 {
		t.Error("创建用户失败")
	}

	t.Log(serviceUser)
}

func TestDeleteUserByUsername(t *testing.T) {

	if !DeleteUserByUsername("admin") {
		t.Error("删除用户失败")
	}
	TestCreateUser(t)
}

func TestGetUserById(t *testing.T) {
	// 选取一个存在的id
	id := 1

	user, ok := GetUserById(uint32(id))
	if !ok {
		t.Error("获取用户失败")
	}
	fmt.Print(user)
}

func TestGetUserByIdWithArticle(t *testing.T) {

	id := 1
	user, ok := GetUserByIdWithArticle(uint32(id))
	if !ok {
		t.Error("获取用户失败")
	}
	fmt.Println(user)
}

func TestGetUserByUsername(t *testing.T) {

	username := "admin"
	user, ok := GetUserByUsername(username)
	if !ok {
		t.Error("获取用户失败")
	}
	fmt.Println(user)
}

func TestIsExistUser(t *testing.T) {

	username := "admin"
	if !IsExistUser(username) {
		t.Log("用户不存在")
	}
	t.Log("用户已存在")
}
