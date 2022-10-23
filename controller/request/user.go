package request

import "github.com/thedevsaddam/govalidator"

type UserRegisterReq struct {
	Username string `valid:"username"`
	Password string `valid:"password"`
}

type UserLoginReq struct {
	Username string `valid:"username"`
	Password string `valid:"password"`
}

func ValidateUserRegisterReq(data interface{}) map[string][]string {
	rules := govalidator.MapData{
		"username": []string{"required", "min:2", "max:20"},
		"password": []string{"required", "min:6", "max:20"},
	}

	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项",
			"min:用户名长度需至少 2 个字符",
			"max:用户名长度不能超过 20 个字符",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需至少 6 个字符",
			"max:密码长度不能超过 20 个字符",
		},
	}
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func ValidateUserLoginReq(data interface{}) map[string][]string {
	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required", "min:6", "max:20"},
	}

	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需至少 6 个字符",
			"max:密码长度不能超过 20 个字符",
		},
	}
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}
	return govalidator.New(opts).ValidateStruct()
}
