package helper

import "strings"

func ParseForm(content string) map[string]string {
	result := make(map[string]string)
	for _, value := range strings.Split(content, "&") {
		kv := strings.Split(value, "=")
		result[kv[0]] = kv[1]
	}
	return result
}
