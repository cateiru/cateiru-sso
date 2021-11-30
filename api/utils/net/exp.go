// Cookieの有効時間を設定します
//
// Example:
//	exp := NewCookieHourExp(10) // 10時間
//	exp2 := NewCookieHourExp(2) // 2時間
//
//	exp3 := NewCookieMinutsExp(10) // 10分
//	exp4 := NewCookieMinutsExp(3) // 3分
//
package net

import "time"

type CookieExp struct {
	time int
	unit time.Duration
}

// 1時間（60分）単位の時間を作成
func NewCookieHourExp(hour int) *CookieExp {
	return &CookieExp{
		time: hour,
		unit: time.Hour,
	}
}

// 分単位の時間を作成
func NewCookieMinutsExp(minuts int) *CookieExp {
	return &CookieExp{
		time: minuts,
		unit: time.Minute,
	}
}

// 時間をtime.Durationで返す
func (c *CookieExp) GetTime() time.Duration {
	return c.unit * time.Duration(c.time)
}

// 時間をミリ秒のintで返す
func (c *CookieExp) GetNum() int {
	switch c.unit {
	case time.Hour:
		return 3600 * c.time
	case time.Minute:
		return 60 * c.time
	default:
		return c.time
	}
}
