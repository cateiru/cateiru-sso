package src

import (
	"net"

	"github.com/labstack/echo/v4"
)

// Fastly のエッジサーバーのIPアドレス一覧
// TODO: 実行時に毎回APIを叩いて更新できるようにする
// ref. `curl "https://api.fastly.com/public-ip-list" | jq ".addresses"`
var fastlyIpAddresses []string = []string{
	"23.235.32.0/20",
	"43.249.72.0/22",
	"103.244.50.0/24",
	"103.245.222.0/23",
	"103.245.224.0/24",
	"104.156.80.0/20",
	"140.248.64.0/18",
	"140.248.128.0/17",
	"146.75.0.0/17",
	"151.101.0.0/16",
	"157.52.64.0/18",
	"167.82.0.0/17",
	"167.82.128.0/20",
	"167.82.160.0/20",
	"167.82.224.0/20",
	"172.111.64.0/18",
	"185.31.16.0/22",
	"199.27.72.0/21",
	"199.232.0.0/16",
}

func FastlyTrust() []echo.TrustOption {
	options := []echo.TrustOption{
		echo.TrustLoopback(false),   // e.g. ipv4 start with 127.
		echo.TrustLinkLocal(false),  // e.g. ipv4 start with 169.254
		echo.TrustPrivateNet(false), // e.g. ipv4 start with 10. or 192.168
	}

	for _, fastlyIp := range fastlyIpAddresses {
		_, ipNet, err := net.ParseCIDR(fastlyIp)
		if err != nil {
			panic(err)
		}

		options = append(options, echo.TrustIPRange(ipNet))
	}

	return options
}
