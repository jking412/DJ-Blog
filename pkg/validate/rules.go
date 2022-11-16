package validate

import (
	"DJ-Blog/internal/service"
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"strings"
)

func Init() {
	// field是结构体中valid后的字段
	// rule是 rules中定义的规则，如 min:3
	// message是 rules中定义的错误信息，如 min:用户名长度需至少 3 个字符
	// value是结构体中的字段值
	govalidator.AddCustomRule("not_exist", func(field string, rule string, message string, value interface{}) error {

		valueStr := value.(string)
		validateField := strings.TrimPrefix(rule, "not_exist")
		if validateField == "username" {
			service.IsExistUser(valueStr)
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v : %v已存在", validateField, valueStr)
		}

		return nil
	})
}
