package common_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src/core/common"
	ua "github.com/mileusna/useragent"
	"github.com/stretchr/testify/require"
)

// ログイン履歴が正しく残されているか
// func TestLoginHistory(t *testing.T) {
// 	config.TestInit(t)

// 	ctx := context.Background()

// 	db, err := database.NewDatabase(ctx)
// 	require.NoError(t, err)
// 	defer db.Close()

// 	userId := utils.CreateID(30)
// 	ip := "198.51.100.0"
// 	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

// 	err = common.SetLoginHistory(ctx, db, userId, ip, userAgent)
// 	require.NoError(t, err)

// 	goretry.Retry(t, func() bool {
// 		loginHistories, err := models.GetAllLoginHistory(ctx, db, userId)
// 		require.NoError(t, err)

// 		return len(loginHistories) == 1 && loginHistories[0].IpAddress == ip
// 	}, "ログイン履歴が格納できている")
// }

func TestUserAgent(t *testing.T) {
	userAgents := []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1b4948 Mobile/14F89 Safari/603.2.4",
		"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
		"Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36",
		"Mozilla/5.0 (Android 4.3; Mobile; rv:54.0) Gecko/54.0 Firefox/54.0",
		"Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.91 Mobile Safari/537.36 OPR/42.9.2246.119956",
		"Opera/9.80 (Android; Opera Mini/28.0.2254/66.318; U; en) Presto/2.12.423 Version/12.16",
	}

	for _, s := range userAgents {
		parse := ua.Parse(s)

		target, err := common.UserAgentToJson(s)
		require.NoError(t, err)

		ex, err := common.ParseUserAgentJson(target)
		require.NoError(t, err)

		require.Equal(t, ex.Name, parse.Name)
		require.Equal(t, ex.Version, parse.Version)
		// 他は省略
	}
}

func TestClientHints(t *testing.T) {
	r := &http.Request{
		Header: http.Header{
			"Sec-Ch-Ua":          {`" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"`},
			"Sec-Ch-Ua-Mobile":   {"?0"},
			"Sec-Ch-Ua-Platform": {"\"Windows\""},
		},
	}

	result, err := common.UACHToJson(r)
	require.NoError(t, err)

	var ua common.UserAgent

	err = json.Unmarshal(result, &ua)

	require.Equal(t, ua.OS, "Windows")
	require.Equal(t, ua.Name, "Chrome")
	require.Equal(t, ua.Version, "100")
	require.Equal(t, ua.Desktop, true)
	require.Equal(t, ua.Mobile, false)
}

func TestUserData(t *testing.T) {
	r := &http.Request{
		Header: http.Header{
			"Sec-Ch-Ua":          {`" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"`},
			"Sec-Ch-Ua-Mobile":   {"?0"},
			"Sec-Ch-Ua-Platform": {"Windows"},
		},
	}

	re, err := common.ParseUserData(r)
	require.NoError(t, err)

	require.NotEmpty(t, re)

	r = &http.Request{
		Header: http.Header{
			"User-Agent": {`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36`},
		},
	}

	re, err = common.ParseUserData(r)
	require.NoError(t, err)

	require.NotEmpty(t, re)
}
