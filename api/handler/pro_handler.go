package handler

// import (
// 	"net/http"

// 	"github.com/cateiru/cateiru-sso/api/core/pro"
// 	"github.com/cateiru/cateiru-sso/api/utils/net"
// )

// // SSOの管理
// func ProSSOHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		proSSOGetHandler(w, r)
// 	case http.MethodPost:
// 		proSSOPostHandler(w, r)
// 	case http.MethodDelete:
// 		proSSODeleteHandler(w, r)
// 	default:
// 		RootHandler(w, r)
// 	}
// }

// // 自分のSSO情報取得
// // pro以上のユーザのみ
// func proSSOGetHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := pro.GetSSOHandler(w, r); err != nil {
// 		net.ResponseError(w, err)
// 	}
// }

// // SSO追加
// // pro以上のユーザのみ
// func proSSOPostHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := pro.AddSSOHandler(w, r); err != nil {
// 		net.ResponseError(w, err)
// 	}
// }

// // SSO削除
// // pro以上のユーザのみ
// func proSSODeleteHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := pro.DeleteSSOHandler(w, r); err != nil {
// 		net.ResponseError(w, err)
// 	}
// }
