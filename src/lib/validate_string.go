package lib

import (
	"regexp"
	"unicode/utf8"
)

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

// パスワードの形式（主に長さが正しいか）を検証する
// 13文字以上、256文字以下
func ValidatePassword(p string) bool {
	pLen := len(p)
	if pLen < 13 || pLen > 256 {
		return false
	}

	// ASCII以外はfalse
	if !(utf8.ValidString(p) && utf8.RuneCountInString(p) == len(p)) {
		return false
	}

	// 繰り返し文字列チェッカー
	chars := map[rune]int{}
	for _, c := range p {
		chars[c]++
	}
	maxCharLen := 0
	for _, c := range chars {
		if maxCharLen < c {
			maxCharLen = c
		}
	}
	// パスワード全体中から一番多い繰り返し文字を引いた値が5未満の場合エラー
	return (pLen - maxCharLen) >= 5
}
