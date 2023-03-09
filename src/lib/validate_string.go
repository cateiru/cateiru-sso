package lib

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var emailReg = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9]+\.)+[a-zA-Z]{2,}$`)
var userNameReg = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var otpReg = regexp.MustCompile(`^[0-9]+$`)

// iCloudなどでPasskeyを共有する際に、UAのOSで判定するため
// AppleのOSは予め定義しておく
var appleOS = []string{
	"macOS",
	"iOS",
	"iPadOS",
}

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

// ユーザ名の検証
// 3文字以上15文字以下
// 使用できる文字は、アルファベット、数字、_のみ
func ValidateUsername(n string) bool {
	nameLength := len(n)
	if nameLength < 3 || nameLength > 15 {
		return false
	}

	return userNameReg.MatchString(n)
}

// OTPの検証
func ValidateOTPCode(o string) bool {
	if len(o) != 6 {
		return false
	}

	return otpReg.MatchString(o)
}

// Passkeyの自動ログイン判定用
func ValidateOS(os string, currentOS string) bool {
	formattedOS := strings.ToLower(os)
	formattedCurrentOS := strings.ToLower(currentOS)

	if formattedOS == formattedCurrentOS {
		return true
	}

	isAppleOS := false
	isAppleCurrentOS := false
	for _, apple := range appleOS {
		formattedAppleOS := strings.ToLower(apple)
		if formattedOS == formattedAppleOS {
			isAppleOS = true
		}
		if formattedCurrentOS == formattedAppleOS {
			isAppleCurrentOS = true
		}
	}
	if isAppleOS && isAppleCurrentOS {
		return true
	}

	return false
}
