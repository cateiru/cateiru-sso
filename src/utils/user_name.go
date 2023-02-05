package utils

import "strings"

// ユーザ名は、[A-Za-z0-9_]{3,15}にマッチする必要があります
func CheckUserName(userName string) bool {
	// 3文字以上15文字以内
	length := len(userName)
	if length < 3 || length > 15 {
		return false
	}

	for _, char := range userName {
		if 'A' <= char && char <= 'Z' {
			// A-Z
			continue
		} else if 'a' <= char && char <= 'z' {
			// a-z
			continue
		} else if char >= '0' && char <= '9' {
			// numbers
			continue
		} else if char == '_' {
			// under line
			continue
		} else {
			return false
		}
	}

	return true
}

// ユーザ名をフォーマットする
func FormantUserName(userName string) string {
	return strings.ToLower(userName)
}
