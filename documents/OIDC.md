# CateiruSSO OpenID Connect 仕様

## 1. はじめに

CateiruSSOは[OpenID Connect](https://openid.net/connect/)を使用してSSOを実装しています。

また、CateiruSSOを使用してSSOサービスを作成する場合**Pro**ユーザが必要です。

[cateiru-sso-example](https://cateiru-sso-example.vercel.app/)（[github](https://github.com/cateiru/cateiru-sso-example)）で試すことが可能です。

## 2. Flow

### 2.2. サービスを作成する

- SSOサービスを作成する場合**Pro**ユーザが必要です。
- 詳しくは、 @cateiru に連絡してください。

1. [ダッシュボード - CateiruSSO](https://sso.cateiru.com/dashboard)にアクセスします。
2. `+`を押してサービス名、送信元URL、リダイレクトURLを入力してサービスを作成します。
    - サービスアイコンは作成後に追加できます。
    - ClientID、Token Secretは作成後に確認できます。
    - サービス名、送信元UR、リダイレクトURLは後で編集ができます。
    - 送信元URL、リダイレクトURLはSSL/TLSのURL（`https`）が必要です。
        - ローカル用のURLは`http://localhost`が使用可能です。
        - 使用しない場合（例: gcloudの認証）はどちらも`direct`にします。
3. ClientID、Token Secretをコピーします。
    - Token Secretは必ず公開しないようにしてください。

### 2.3. リダイレクトを設定する

- ログインボタンを押したときにこのURLへ飛ぶようにします。
- エンドポイントは、`https://sso.cateiru.com/sso/login` です。
- `client_id`のクエリは2.2で作成したサービスのClientID、redirect_uriはサービス作成時に設定したリダイレクトURIを設定してください。

```ts
const ENDPOINT = "https://sso.cateiru.com/sso/login"
const clientId = ""
const redirectURL = ""

const url = `${ENDPOINT}?scope=openid&response_type=code&client_id=${clientId}&redirect_uri=${redirectURL}&prompt=consent`

res.redirect(url)
```

### 2.4. リダイレクト先を設定する

- サービス作成時に設定したリダイレクトURL先を設定します
- CateiruSSOは`code`クエリにトークンを付与してリダイレクトします。
  - トークンは、Token Secretを一緒にトークンエンドポイント（`https://api.sso.cateiru.com/oauth/token`）へ送信し、IDTokenを取得します。
  - Token Secretは`Authorization`ヘッダに`Basic [secret]`を付与します。

```ts
const TOKEN_ENDPOINT = "https://api.sso.cateiru.com/oauth/token"
const redirectURL = ""
const code = ""
const tokenSecret = ""

    const res = await request("GET", `${TOKEN_ENDPOINT}?grant_type=authorization_code&code=${code}&redirect_uri=${redirectURL}`, {
    headers: {
    authorization: `Basic ${tokenSecret}`
    }
})
```

- トークンエンドポイントのレスポンスは以下の通りです。

  ```json
  {
    "access_token": "", // codeで帰ってくるトークンと一緒
    "token_type": "", // Bearer
    "refresh_token": "", // リフレッシュトークン。これを使用することで新しいaccess_tokenを取得できます
    "expires_in": "", // access_tokenの有効期限(秒)
    "id_token": "", // IDToken. JWTです
  }
  ```

### 2.5. IDTokenを公開鍵を使用して検証します

- JWTの公開鍵は`https://api.sso.cateiru.com/oauth/jwt/key`で取得できます。
  - CateiruSSOのデプロイごとに値が変わってしまうため、JWTを検証する度にGETすることをおすすめします。
- JWTの暗号化アルゴリズムは`RS256`です。
- 詳しくは[example](https://github.com/cateiru/cateiru-sso-example/blob/main/utils/jwt.ts)を参照してください。
