package request

import "github.com/thedevsaddam/govalidator"

type UserRegisterReq struct {
	Username  string `valid:"username" json:"username,omitempty"`
	Password  string `valid:"password" json:"password,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
}

type UserLoginReq struct {
	Username string `valid:"username" json:"username,omitempty"`
	Password string `valid:"password" json:"password,omitempty"`
}

// ValidateUserRegisterReq 验证用户注册请求
func ValidateUserRegisterReq(data interface{}) map[string][]string {
	// 定义验证规则
	rules := govalidator.MapData{
		"username": []string{"required", "min:2", "max:20", "not_exist:username"},
		"password": []string{"required", "min:6", "max:20"},
	}

	// 定义验证错误信息
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
		"username": []string{"required", "min:2", "max:20", "exist:username"},
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
