// 環境変数 DEPLOY_MODE を読み取り、DEPLOY_MODEに格納します
package utils

import "os"

var DEPLOY_MODE string

func init() {
	deployMode := os.Getenv("DEPLOY_MODE")

	if len(deployMode) == 0 {
		deployMode = "develop"
	}

	DEPLOY_MODE = deployMode
}
