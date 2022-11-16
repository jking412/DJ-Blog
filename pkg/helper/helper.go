package helper

import (
	"strings"
	"unicode/utf8"
)

func ParseForm(content string) map[string]string {
	result := make(map[string]string)
	for _, value := range strings.Split(content, "&") {
		kv := strings.Split(value, "=")
		result[kv[0]] = kv[1]
	}
	return result
}

func GetUTF8StringLength(str string) int {
	return utf8.RuneCountInString(str)
}
