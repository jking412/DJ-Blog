package core

import (
	"DJ-Blog/boot"
	"DJ-Blog/internal/http/request"
	"DJ-Blog/internal/model"
	"DJ-Blog/internal/service"
	"DJ-Blog/pkg/helper"
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestUserRegister(t *testing.T) {
	boot.Initialize()

	// 成功的注册
	var username, password string

	username = faker.Username()
	password = faker.Password()
	ul := helper.GetUTF8StringLength(username)
	pl := helper.GetUTF8StringLength(password)
	if ul < 6 {
		username = username + username
	}
	if pl < 6 {
		password = password + password
	}
	if ul > 20 {
		username = username[:20]
	}
	if pl > 20 {
		password = password[:20]
	}

	errs, ok := UserRegister(&request.UserRegisterReq{
		Username: username,
		Password: password,
	})

	if !ok {
		t.Error(errs)
	}

	// 不合法的用户名
	// 用户名过短
	errs, ok = UserRegister(&request.UserRegisterReq{
		Username: "1",
		Password: password,
	})
	if ok {
		t.Error("过短的用户名被错误的通过了")
	}
	t.Log(errs)

	// 用户名过长
	errs, ok = UserRegister(&request.UserRegisterReq{
		Username: "123456789012345678901",
		Password: password,
	})
	if ok {
		t.Error("过长的用户名被错误的通过了")
	}
	t.Log(errs)

	// 用户名为空
	errs, ok = UserRegister(&request.UserRegisterReq{
		Username: "",
		Password: password,
	})
	if ok {
		t.Error("空的用户名被错误的通过了")
	}
	t.Log(errs)
	//	TODO: 重复的用户名，用户名中包含特殊字符，用户名中包含空格

	// TODO:不合法的密码

}

func TestUserLogin(t *testing.T) {
	boot.Initialize()

	var username, password string

	username = "admin"
	password = "123456"

	if !service.IsExistUser(username) {
		service.CreateUser(&model.UserModel{
			Username: username,
			Password: password,
		})
	}

	errs, ok := UserLogin(&request.UserLoginReq{
		Username: username,
		Password: password,
	})

	if !ok {
		t.Error(errs)
	}
}


