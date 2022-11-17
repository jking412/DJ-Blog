package service

import (
	"DJ-Blog/internal/model"
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestCreateUser(t *testing.T) {

	user := &model.UserModel{
		Username:  faker.Username(),
		Password:  faker.Password(),
		AvatarUrl: faker.URL(),
	}

	serviceUser := CreateUser(user)

	if serviceUser.Id == 0 {
		t.Error("创建用户失败")
	}

	t.Log(serviceUser)
}
