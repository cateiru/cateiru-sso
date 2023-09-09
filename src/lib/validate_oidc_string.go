package lib

// サポートしているOIDCのスコープ
var AllowScopes = []string{
	"openid",
	"profile",
	"email",
}

func ValidateScope(s string) bool {
	for _, allow := range AllowScopes {
		if allow == s {
			return true
		}
	}
	return false
}
