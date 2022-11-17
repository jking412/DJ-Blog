package core

import (
	"DJ-Blog/boot"
	"DJ-Blog/internal/http/request"
	"DJ-Blog/pkg/helper"
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestUserRegister(t *testing.T) {
	boot.Initialize()

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
}

func TestUserLogin(t *testing.T) {

}
