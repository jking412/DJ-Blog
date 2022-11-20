package validate

import (
	"DJ-Blog/internal/service"
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"strings"
)

func init() {
	// field是结构体中valid后的字段
	// rule是 rules中定义的规则，如 min:3
	// message是 rules中定义的错误信息，如 min:用户名长度需至少 3 个字符
	// value是结构体中的字段值
	govalidator.AddCustomRule("not_exist", func(field string, rule string, message string, value interface{}) error {

		validateField := strings.TrimPrefix(rule, "not_exist:")

		if validateField == "username" && service.IsExistUser(value.(string)) {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%s : %s已存在", validateField, value.(string))
		}

		return nil
	})

	govalidator.AddCustomRule("exist", func(field string, rule string, message string, value interface{}) error {

		validateField := strings.TrimPrefix(rule, "exist:")

		if validateField == "username" && !service.IsExistUser(value.(string)) {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%s : %s不存在", validateField, value.(string))
		}

		if validateField == "articleId" && !service.IsExistArticleById(value.(uint32)) {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%s : %d不存在", validateField, value.(uint32))
		}

		return nil
	})

}
