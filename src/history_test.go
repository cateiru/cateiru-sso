package src_test

import "testing"

func TestHistoryClientLoginHandler(t *testing.T) {
	t.Run("成功: ログインしているクライアントが返る", func(t *testing.T) {})

	t.Run("成功: クライアントが存在しないのは返らない", func(t *testing.T) {})

	t.Run("成功: 何もログインしていないときは空", func(t *testing.T) {})
}

func TestHistoryClientHandler(t *testing.T) {
	t.Run("成功: ログイン履歴が返る", func(t *testing.T) {})

	t.Run("成功: 個数を指定できる", func(t *testing.T) {})

	t.Run("成功: クライアントが存在しないのは返らない", func(t *testing.T) {})

	t.Run("成功: 何もログインしていないときは空", func(t *testing.T) {})
}

func TestHistoryLoginDeviceHandler(t *testing.T) {
	t.Run("成功: ログインしているデバイスを取得できる", func(t *testing.T) {})

	t.Run("成功: 履歴はあるが、リフレッシュトークンが存在しない場合は返さない", func(t *testing.T) {})
}

func TestHistoryLoginHistoryHandler(t *testing.T) {
	t.Run("成功: ログイン履歴を取得できる", func(t *testing.T) {})

	t.Run("成功: 個数を指定できる", func(t *testing.T) {})
}

func TestHistoryLoginTryHistoryHandler(t *testing.T) {
	t.Run("成功: ログイントライ履歴を取得できる", func(t *testing.T) {})

	t.Run("成功: 個数を指定できる", func(t *testing.T) {})
}
