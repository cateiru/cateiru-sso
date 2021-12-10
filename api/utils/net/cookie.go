// Cookieの作成、取得、削除を行います。
//
// Example:
//	cookie := NewCookie("example.com", true, http.SameSiteNoneMode)
//	exp := NewCookieHourExp(10)
//	cookie.Set(w, "key", "value", exp)
//
//	value, err := cookie.Get(r, "key")
//
//	err := cookie.Delete(w, r, "key")
//
package net

import (
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/api/logging"
)

type Cookie struct {
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
	SomeSite http.SameSite
}

func NewCookie(domain string, secure bool, someSite http.SameSite, httpOnly bool) *Cookie {
	path := "/"

	return &Cookie{
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
		SomeSite: someSite,
	}
}

// Cookieをセットします
func (c *Cookie) Set(w http.ResponseWriter, key string, value string, exp *CookieExp) {
	cookie := &http.Cookie{
		Name:  key,
		Value: value,

		Secure:   c.Secure,
		Path:     c.Path,
		Domain:   c.Domain,
		HttpOnly: c.HttpOnly,
		SameSite: c.SomeSite,
	}

	// 有効期限を設定する
	// IsSession = trueの場合はセッションクッキーにするため設定しない
	if !exp.IsSession {
		cookie.Expires = time.Now().Add(exp.GetTime())
		cookie.MaxAge = exp.GetNum()
	}

	logging.Sugar.Debugf("Set the cookie. key: %s, value: %s, exp: %vs", key, value, cookie.MaxAge)
	http.SetCookie(w, cookie)
}

// Cookieを削除します
func (c *Cookie) Delete(w http.ResponseWriter, req *http.Request, key string) error {
	cookie, err := req.Cookie(key)
	if err != nil {
		return err
	}

	logging.Sugar.Debugf("Delete cookie. key: %s, value: %s", key, cookie.Value)

	cookie.Expires = time.Unix(0, 0)
	cookie.MaxAge = -1

	cookie.Secure = c.Secure
	cookie.Path = c.Path
	cookie.Domain = c.Domain
	cookie.HttpOnly = c.HttpOnly
	cookie.SameSite = c.SomeSite

	http.SetCookie(w, cookie)

	return nil
}

// keyで指定した名前のcookieを返します
func GetCookie(req *http.Request, key string) (string, error) {
	cookie, err := req.Cookie(key)

	if err != nil {
		return "", err
	}
	logging.Sugar.Debugf("Get cookie. key: %s, value: %s", key, cookie.Value)

	return cookie.Value, nil
}
