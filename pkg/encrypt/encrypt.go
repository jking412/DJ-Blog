package encrypt

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logrus.Warn("加密密码失败", err.Error())
		return password
	}
	return string(encryptPassword)
}

func CheckPassword(password string, encryptPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	if err != nil {
		logrus.Warn("密码错误", err.Error())
		return false
	}
	return true
}

func IsEncrypt(password string) bool {
	return len(password) == 60
}
