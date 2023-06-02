package lib

import (
	"regexp"
	"time"
	"unicode/utf8"

	"golang.org/x/text/language"
)

var emailReg = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9]+\.)+[a-zA-Z]{2,}$`)
var userNameReg = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var otpReg = regexp.MustCompile(`^[0-9]+$`)

var allowContentType = []string{
	"image/gif",
	"image/jpeg",
	"image/png",
	"image/webp",
}

// サポートしているOIDCのスコープ
var allowScopes = []string{
	"openid",
	"profile",
	"email",
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

// 性別判定
// 0: 不明、1: 男性、2: 女性、9: 適用不能
func ValidateGender(gender string) bool {
	if gender == "0" || gender == "1" || gender == "2" || gender == "9" {
		return true
	}
	return false
}

// 誕生日 YYYY-MM-DDの形式
func ValidateBirthDate(b string) (*time.Time, bool) {
	birthDate, err := time.Parse(time.DateOnly, b)
	if err != nil {
		return nil, false
	}

	// 現在時刻より先なのはありえないのでそのときはfalse
	if time.Now().Before(birthDate) {
		return nil, false
	}
	return &birthDate, true
}

func ValidateLocale(l string) bool {
	_, err := language.Parse(l)
	return err == nil
}

func ValidateContentType(c string) bool {
	for _, allow := range allowContentType {
		if allow == c {
			return true
		}
	}
	return false
}

func ValidateScope(s string) bool {
	for _, allow := range allowScopes {
		if allow == s {
			return true
		}
	}
	return false
}
