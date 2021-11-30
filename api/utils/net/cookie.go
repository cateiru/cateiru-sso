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

func NewCookie(domain string, secure bool, someSite http.SameSite) *Cookie {
	path := "/"

	return &Cookie{
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: true,
		SomeSite: someSite,
	}
}

// Cookieをセットします
func (c *Cookie) Set(w http.ResponseWriter, key string, value string, exp *CookieExp) {
	expires := time.Now().Add(exp.GetTime())
	maxAge := exp.GetNum()

	logging.Sugar.Debugf("Set the cookie. key: %s, value: %s, exp: %vs", key, value, maxAge)

	cookie := &http.Cookie{
		Name:  key,
		Value: value,

		Expires: expires,
		MaxAge:  maxAge,

		Secure:   c.Secure,
		Path:     c.Path,
		Domain:   c.Domain,
		HttpOnly: c.HttpOnly,
		SameSite: c.SomeSite,
	}

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
func (c *Cookie) Get(req *http.Request, key string) (string, error) {
	cookie, err := req.Cookie(key)

	if err != nil {
		return "", err
	}
	logging.Sugar.Debugf("Get cookie. key: %s, value: %s", key, cookie.Value)

	return cookie.Value, nil
}
