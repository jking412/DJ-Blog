package request

import "github.com/thedevsaddam/govalidator"

type UserRegisterReq struct {
	// example 是 swagger 文档中的示例的关键字，用于生成 swagger 文档。
	// 这里的值并没有使用一个真的示例，而是用了这个数据的限制作为示例。
	// TODO: 利用这个example值，做自动化的测试，限制条件之间用分号分隔，如：min:2,max:20
	Username  string `valid:"username" json:"username,omitempty" example:"required;min:2;max:20"`
	Password  string `valid:"password" json:"password,omitempty" example:"required;min:6;max:20"`
	AvatarUrl string `json:"avatar_url,omitempty"`
}

type UserLoginReq struct {
	Username string `valid:"username" json:"username,omitempty" example:"required;min:2;max:20"`
	Password string `valid:"password" json:"password,omitempty" example:"required;min:6;max:20"`
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
