package src

// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#CodeIDToken
type AuthorizationCodeFlowClaims struct {
	IDTokenClaimsBase
	StandardClaims

	// Access Token のハッシュ値
	AtHash string `json:"at_hash,omitempty"`
}

// `nonce` は IDTokenClaimsBase で定義されているため省略
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#HybridIDToken
type HybridFlowClaims struct {
	IDTokenClaimsBase
	StandardClaims

	// Access Token のハッシュ値
	AtHash string `json:"at_hash,omitempty"`

	// Code のハッシュ値
	CHash string `json:"c_hash,omitempty"`
}

// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#SelfIssuedResponse
type SelfIssuesClaims struct {
	IDTokenClaimsBase
	StandardClaims

	// Self-Issued OpenID Provider から発行された ID Token の署名検証に使われる公開鍵
	SubJWK string `json:"sub_jwk"`
}

// OpenIDConnectの ID Tokenのクレームベース定義
// ref. https://openid.net/specs/openid-connect-core-1_0.html#IDToken
type IDTokenClaimsBase struct {
	// レスポンスを返した Issuer の Issuer Identifier
	Iss string `json:"iss"`

	// Subject Identifier. Client に利用される前提で, Issuer のローカルでユニークであり再利用されない End-User の識別子
	Sub string `json:"sub"`

	// ID Token の想定されるオーディエンス (Audience)
	// RFC 3339 で表現される数値
	Aud string `json:"aud"`

	// ID Token の有効期限
	Exp int64 `json:"exp"`

	// JWT 発行時刻
	// リクエストに max_age が含まれていた場合, この Claim は必須である
	// RFC 3339 で表現される数値
	Iat int64 `json:"iat,omitempty"`

	// End-User の認証が発生した時刻
	AuthTime int64 `json:"auth_time"`

	// Client セッションと ID Token を紐づける文字列値. リプレイアタック防止のために用いられる. Authentication Request で指定されたままの値を ID Token に含める
	Nonce string `json:"nonce,omitempty"`

	// Authentication Context Class Reference. 実施された認証処理が満たす Authentication Context Class を表す Authentication Context Class Reference 値を示す文字列
	Acr string `json:"acr,omitempty"`

	// Authentication Methods References. 認証時に用いられた認証方式を示す識別子文字列の JSON 配列
	Amr string `json:"amr"`

	// ID Token 発行対象である認可された関係者 (authorized party).
	Azp string `json:"azp"`
}

// OIDCのスタンダードクレーム
// `sub` は IDTokenClaimsBase で定義されているため省略
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#StandardClaims
type StandardClaims struct {
	Name                string  `json:"name"`
	GivenName           string  `json:"given_name"`
	FamilyName          string  `json:"family_name"`
	MiddleName          string  `json:"middle_name"`
	Nickname            string  `json:"nickname"`
	PreferredUsername   string  `json:"preferred_username"`
	Profile             string  `json:"profile"`
	Picture             string  `json:"picture"`
	Website             string  `json:"website"`
	Email               string  `json:"email"`
	EmailVerified       bool    `json:"email_verified"`
	Gender              string  `json:"gender"`
	BirthDate           string  `json:"birthdate"`
	ZoneInfo            string  `json:"zoneinfo"`
	Locale              string  `json:"locale"`
	PhoneNumber         string  `json:"phone_number"`
	PhoneNumberVerified bool    `json:"phone_number_verified"`
	Address             Address `json:"address"`
	UpdatedAt           string  `json:"updated_at"`
}

// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#AddressClaim
type Address struct {
	// 表示や宛名ラベルで使用するために整形された完全な郵送先住所
	Formatted string `json:"formatted"`

	// 家番号, ストリート名, 私書箱, 複数行からなる拡張された住所情報を含むことができる
	StreetAddress string `json:"street_address"`

	// 市町村などを表す city, locality を示す要素
	Locality string `json:"locality"`

	// 都道府県などを表す state, province, prefecture, region を示す要素
	Region string `json:"region"`

	// 郵便番号を表す ZIP code, postal code を示す要素
	PostalCode string `json:"postal_code"`

	// 国名を示す要素
	Country string `json:"country"`
}

// 成功時のパラメータ
// oidc ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#AuthResponse
// oauth2 ref. https://openid-foundation-japan.github.io/rfc6749.ja.html#code-authz-resp
type SuccessfulAuthenticationResponse struct {
	// 認可コードは認可サーバーによって許可される
	Code string `json:"code"`

	// クライアントからの認可リクエストに state パラメーターが含まれていた場合は必須 (REQUIRED). クライアントから受け取った値をそのまま返す
	State string `json:"state"`
}

// エラーレスポンスのパラメータ
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#AuthError
type AuthenticationErrorResponse struct {
	// エラーコード
	Error string `json:"error"`

	// 人間が読める ASCII エンコードされたエラーの説明文
	ErrorDescription string `json:"error_description,omitempty"`

	// エラーについての追加情報を含むWebページのURI
	ErrorURI string `json:"error_uri,omitempty"`

	// OAuth 2.0 の state の値
	State string `json:"state,omitempty"`
}
