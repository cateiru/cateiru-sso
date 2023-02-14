package lib

import "regexp"

var emailReg = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9]+\.)+[a-zA-Z]{2,}$`)

// Emailの形式が正しいかを検証する
// 1~255文字まで & 正規表現
func ValidateEmail(e string) bool {
	eLen := len(e)
	if eLen <= 0 || eLen > 256 {
		return false
	}
	return emailReg.MatchString(e)
}
