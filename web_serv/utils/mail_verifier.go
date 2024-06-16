package utils

import "regexp"

// 校验邮箱
func VerifyEmail(email string) bool {
	// 匹配电子邮箱
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)

	return reg.MatchString(email)
}
