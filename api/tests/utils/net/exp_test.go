package net_test

import (
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/stretchr/testify/require"
)

// 時間
func TestExpHour(t *testing.T) {
	hours := []int{1, 10, 24, 48, 100}

	for _, hour := range hours {
		exp := net.NewCookieHourExp(hour)

		require.Equal(t, exp.GetTime(), time.Duration(hour)*time.Hour, "GetTimeの時間が違う")
		require.Equal(t, exp.GetNum(), hour*60*60, "GetNumの時間が違う")
	}
}

// 分
func TestExpMinute(t *testing.T) {
	minutes := []int{1, 10, 60, 24, 48, 100, 3600}

	for _, minute := range minutes {
		exp := net.NewCookieMinutsExp(minute)

		require.Equal(t, exp.GetTime(), time.Duration(minute)*time.Minute, "GetTimeの時間が違う")
		require.Equal(t, exp.GetNum(), minute*60, "GetNumの時間が違う")
	}
}
