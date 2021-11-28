package utils

import "os"

// ワンタイムパスワードを発行する組織名
var ONETIME_PASSWORD_ISSUER string = os.Getenv("ONETIME_PASSWORD_ISSUER")

// ワンタイムパスワードのsecret
var ONETIME_PASSWORD_SECRET []byte = []byte(os.Getenv("ONETIME_PASSWORD_SECRET"))

// デプロイモード
var DEPLOY_MODE string = os.Getenv("DEPLOY_MODE")

// reCHAPTCHAのsecret
var RECHAPTCHA_SECRET string = os.Getenv("RECHAPTCHA_SECRET")
