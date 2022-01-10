# API仕様

## `/`

- Root Handler
- とくになにもなし

## `/create`

- POST
  - 一時的なアカウントの作成
  - Req
    - application/json

        ```ts
        {
            mail: string,
            password: string,
            re_captcha: string, // reCHAPTCHAのトークン
        }
        ```

  - Res
    - application/json

        ```ts
        {
            client_check_token: string,
        }
        ```

## `/create/verify`

- GET
  - Websocketに変わる
  - メール認証の可否をリアルタイムに取得する
  - Req
    - `?cct=`: `POST /create`で帰ってきたClientCheckTokenを指定する
  - Res
    - 認証された場合、Websocketで`true`という文字列が返る
    - Websocketのcloseはclient側で行い、サーバ側で終了した場合はエラーである。
- POST
  - メール認証する
  - Req
    - application/json

    ```ts
    {
        mail_token: string, // mail tokenはメールアドレスにクエリパラメータ付きURLとして送信される
    }
    ```

  - Res
    - application/json

    ```ts
    {
        keep_this_page: boolean, // Websocketがすでに閉じている = 元のウィンドウが消されている場合はtrueが帰ってくる。
        buffer_token: string, // ? 何故あるのか忘れた。 keep_this_pageがtrueのときのみある
        client_check_token: string, // keep_this_pageがfalseで元ウィンドウが開かれていた場合、ユーザにどちらのウィンドウで操作するかを選択してもらう。メールから開いたウィンドウで続きをする場合、このtokenを使用して`HEAD /create/verify`にアクセスしてcookieをセットする
    }
    ```

- HEAD
  - WebsocketでVerify=trueになった場合、このエンドポイントを叩きcookieをセットする
  - Req
    - `?token=client_check_token`
  - Cookie
    - `buffer-token`: `/create/info`で使用する

## `/create/info`

- POST
  - ユーザの名前、ユーザ名、テーマ、アバターなどを設定する
  - TODO: ユーザ名は一意にしたい

  - Req
    - application/json

    ```ts
    {
        first_name: string,
        last_name: string,

        user_name: string,

        theme: string,
        avatar_url: string,
    }
    ```

  - Cookie
    - `session-token`: ログインセッション用トークン。同一セッション内のみ有効
    - `refresh-token`: ユーザを識別するトークン。セッションごとに値が更新される。有効期限: 7日

## `/login`

- POST
  - メールアドレス、パスワード、（ワンタイムパスワード）を使用してログインする
  - OTPが必要な場合`otp_token`をセットする
  - Req
    - application/json

    ```ts
    {
        mail: string,
        password: string,
        re_chaptcha: string,
    }
    ```

  - Res
    - application/json or null

    ```ts
    {
        is_otp: boolean, // OTPが必要かどうか
        otp_id: string, // `otp_token` cookieにセットされているやつ
    }
    ```

  - Cookie
    - OTPが必要ない場合のみセットする
      - `session-token`
      - `refresh-token`

## `/login/onetime`

- POST
  - `POST /login`でOTPが必要な場合、このエンドポイントを使用してOTPに検証を行う
  - Req
    - application/json

    ```ts
    {
        passcode: string, // OTPのパスコード
    }
    ```

  - Cookie
    - `session-token`
    - `refresh-token`

## `/me`

- Get
  - ユーザ情報を取得する
  - ログインしている必要がある
  - Res
    - application/json

    ```ts
    {
        first_name: string,
        last_name: string,
        user_name: string,

        role: string[],
        mail: string,

        theme: string,
        avatar_url: string,

        user_id: string,
    }
    ```

## `/pro/sso`

- GET
  - 自分が権限を持っているSSOの情報を取得する
  - Roleがpro以上の権限を持つユーザ限定
  - Res
    - application/json

    ```ts
    {
        sso_publickey: string,

        sso_secretkey: string,
        sso_privatekey: string,

        name: string,
        login_only: boolean, // SSOをログインのみに使用するかどうか。trueの場合、session, refresh tokenは無いため一度のみのログインになる

        from_url: string[], // ssoにリダイレクトする元のURL。1つor複数
        to_url: string[], // ssoログイン後にリダイレクトするURL。from_urlが1つの場合は1つ、複数の場合は同じ個数である必要があります。

        session_token_period?: number, // login_onlyがfalseの場合
        refresh_token_period?: number, // login_onlyがfalseの場合

        user_id: string,
    }
    ```

- POST
  - SSOを追加する
  - Roleがpro以上の権限を持つユーザ限定
  - Req
    - application/json

    ```ts
    {
        name: string,

        from_url: string[], // ssoにリダイレクトする元のURL。1つor複数
        to_url: string[], // ssoログイン後にリダイレクトするURL。from_urlが1つの場合は1つ、複数の場合は同じ個数である必要があります。

        login_only: boolean, // SSOをログインのみに使用するかどうか。trueの場合、session, refresh tokenは無いため一度のみのログインになる

        session_token_period?: number, // 分。login_onlyがfalseの場合
        refresh_token_period?: number, // 分。login_onlyがfalseの場合
    }
    ```

  - Res
    - application/json

    ```ts
    {
        public_key: string,
        secret_key: string,
        private_key: string,
    }
    ```

- DELETE
  - 指定したSSOを削除する
  - Roleがpro以上の権限を持つユーザ限定&そのSSOの作成者限定
  - Req
    - `?key=`: public_keyを指定する

## `/password/forget`

- POST
  - パスコードを忘れた場合の変更用
  - 指定したメールアドレスに再登録用のURLを送信する
  - Req
    - application/json

    ```ts
    {
        mail: string, // そのメールアドレスを使用しているアカウントがすでに存在する必要がある
    }
    ```

## `/password/forget/accept`

- POST
  - パスコードを忘れた場合の再登録メールのURL認証用
  - Req
    - application/json

    ```ts
    {
        forget_token: string,
        new_password: string,
    }
    ```

## `/user/mail`

- GET
  - ユーザのメールアドレスを取得する
  - ログインしている必要がある
  - Resp
    - application/json

    ```ts
    {
        mail: string,
    }
    ```

- POST
  - メールアドレスを変更する
  - 一度認証メールを送信する
  - Req
    - application/json

    ```ts
    {
        // `change`: new_mailのアドレス先に確認メールを送信する。new_mailが必須
        // `verify`: メールの認証をする。mail_tokenが必須
        type: string,

        new_mail?: string,
        mail_token?: string,
    }
    ```

## `/user/password`

- POST
  - パスコードを変更する
  - Req
    - application/json

    ```ts
    {
        new_password: string,
        old_password: string,
    }
    ```

## `/user/otp`

- GET
  - OTPのトークンURLを取得する
  - Res
    - application/json

    ```ts
    {
        id: string,
        otp_token: string,
    }
    ```

- POST
  - OTPを新規設定、削除を行う
  - Req
    - application/json

    ```ts
    {
        // `enable`: OTPの新規設定
        // `disable`: OTPの削除。idはいらない
        type: string,
        passcode: string,
        id?: string, // `GET /user/otp`で取得したID
    }
    ```

  - Res
    - application/json or null
    - `type = enable`のみ

    ```ts
    {
        backups: string[],
    }
    ```

## `/user/otp/backup`

- GET
  - OTPのbackup codeを返す
  - ログインしている必要がある
  - Res
    - application/json

    ```ts
    {
        codes: string[],
    }
    ```

TODO: ~~ ## `/user/history/access` ~~

## `/user/history/login`

- GET
  - ユーザのログイン履歴を取得する
  - ログインしている必要がある
  - Req
    - `?limit=`を指定すると~~最新から~~その数分取得できる (TODO: 日付でorderbyする)
  - Res
    - application/json

    ```ts
    {
        access_id: string,
        date: Date,
        ip_address: string,
        user_agent: string,
        is_sso: string,
        sso_publicKey?: string,
        user_id: string,
    }[]
    ```

## `/logout`

- GET
  - ログアウトする
  - ログインしている必要がある
- DELETE
  - アカウントを削除する
  - ログインしている必要がある
