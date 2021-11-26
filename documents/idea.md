# Cateiru認証アイデア

## TL;DR

- 独自SSOクライアント
  - Google SSOみたいなかんじ
- Cateiruの作ったサービスなどで使えればいいなぁ

## 定義

- Routes
  - `sso.cateiru.com`
    - `/`
      - ルートページ
    - `/create`
      - アカウント作成ページ
      - `?m=[token]`を付加しメール確認で送られる
    - `/dashboard`
      - ダッシュボード
    - `/login`
      - ログイン
    - `/oauth`
      - `/oauth/login`
        - OAuthを使用したSSOログイン
      - ~~`/oauth/cert`~~
        - トークンの検証API
  - `api.sso.cateiru.com`
    - `/create`
      - POST: アカウント作成（メールアドレス、パスワード）
      - `/create/verify`
        - WS: メールアドレス確認待機WS
        - POST: メールアドレスから開いたときにトークンを送信してcookie作成
      - `/create/accept`
        - POST: WS接続中メールアドレスで確認できた場合にWSを閉じてここにトークンを送信しcookie作成
      - `/create/onetime`
        - POST: ワンタイムパスワード作成
      - `/create/info`
        - POST: ユーザ情報（名前、テーマ、プロフィール画像）
    - `/login`
      - POST: メールアドレス、パスワード、ワンタイムパスワードを送信しcookieを作成
      - `/login/checksso`
        - POST: sso_public_keyを送信し、keyが存在するかチェックと、URLを取得
      - `/login/sso`
        - POST: メールアドレス、パスワード、ワンタイムパスワードを送信しリダイレクト
    - `/me`
      - ユーザ情報参照
    - `/admin`
      - `/admin/pro`
        - GET: proユーザ一覧を取得
        - POST: proユーザ追加
        - DELETE: `?id=[user id]`proユーザ削除
      - `/admin/user`
        - GET: 全ユーザ情報取得
        - DELETE: `?id=[user id]`ユーザ削除
      - `/admin/ban`
        - POST: 特定ユーザBan
      - `/admin/status`
        - GET: Workerのログとmail_tokenのDB情報などのWorkerで操作するDB情報を取得
        - POST: mail_token削除など
    - `/sso`
      - GET: SSO情報取得
      - POST: SSO追加
      - DELETE: `?id=[id]`SSO削除
      - （各SSOはidを作成して管理）
    - `/user`
      - `/user/mail`
        - GET: メールアドレス取得
        - POST: メールアドレス更新
      - `/user/password`
        - POST: パスワード変更
      - `/user/onetime`
        - POST: ワンタイムパスワード変更
      - `/user/access`
        - GET: SSOログイン履歴取得
        - POST: ログアウト処理など
      - `/user/history`
        - GET: ログイン履歴取得
    - `/logout`
      - GET: ログアウト
      - DELETE: アカウント削除
    - `/oauth/cert`
      - POST: トークン検証
- トークン
  - `mail_token`
    - メールアドレス確認用トークン。メールアドレスに`?m=[token]`とパラメータをつけて送信する
    - POST `/create/verify`でcookie作成し、トークンは削除
  - `mail_accept_token`
    - 元ページでWSがつながっている場合に認証が完了したときにWS経由で返すトークン
    - POST `/create/accept`でそのトークンを送ることでcookieを作成
  - `session_token`
    - 認証用セッショントークン
  - `user_token`
    - ユーザトークン。session_tokenを作成するたびにトークンが更新される
    - 有効期限: 48h
  - `sso_public_key`
    - アクセス元がSSOへリダイレクトする際に付与する。サービス認識用
  - `sso_secret_key`
    - SSOから返ってきたトークン（PASETO）を復号化するトークン。サーバーなどの利用者の目には見えないところで使用
  - `sso_private_key`
    - 内部でユーザデータを暗号化するトークン
  - `sso_token`
    - SSOのセッショントークン
    - 有効期限はダッシュボードで決められる
  - `sso_refresh_token`
    - SSOのリフレッシュトークン
    - 有効期限はダッシュボードできめられる

### テーブル

- 認証

    ```ts
    {
        mail: string,
        password: string, // パスワードはハッシュ化
        create_account_date: Date,
        user_id: string,

        onetime_password_key?: string,
    }
    ```

- メールアドレス認証

    ```ts
    {
        mail_token: string,
        create_date: Date,
        period_minute: number = 10,

        mail: string,
        password: string, // パスワードはハッシュ化
    }
    ```

- 元ページでWSがつながっている場合に認証が完了したときにWS経由で返すトークンのやつ

    ```ts
    {
        mail_accept_token: string,
        mail_token: string,
        create_date: Date,
        period_minute: number = 10,
    }
    ```

- User情報

    ```ts
    {
        user_id: string,
        sso_user_id: string,

        first_name: string,
        last_name: string,

        mail: string,

        theme: string,
        avatar_url?: string,
    }
    ```

- Userログイン情報

    ```ts
    {
        user_id: string,

        histories: {
            access_id: string,
            date: Date,
            ip_address: string,
        }[]
    }
    ```

- セッショントークン

    ```ts
    {
        session_token: string,
        user_id: string,
        create_date: Date,
        period_hour: number = 6,
    }
    ```

- リフレッシュトークン

    ```ts
    {
        refresh_token: string,
        session_token: string,

        user_id: string,
        create_date: Date,
        period_hour: number = 48,
    }
    ```

- SSO

    ```ts
    {
        sso_public_key: string,
        sso_secret_key: string,
    }
    ```

- ユーザのSSO

    ```ts
    {
        user_id: string,
        sso: {
            id: string,
            name: string,
            from_url: string[],
            to_url: string[],
            session_token_period: number,
            refresh_token_period: number,
            sso_public_key: string,
            sso_secret_key: string
        }[],
    }
    ```

- SSOセッショントークン

    ```ts
    {
        sso_session_token: string,
        create_date: Date,
        period_hour: number;
        user_id: string,
    }
    ```

- SSOリフレッシュトークン

    ```ts
    {
        sso_refresh_token: string,
        sso_session_token: string,
        create_date: Date,
        period_hour: number;
        user_id: string;
    }
    ```

## 機能

- アカウント作成
  1. メールアドレス、パスワードを入力
  2. メールアドレスを確認
     - そのページに戻る、もしくはメールに送られてきたURLから続きをやるをユーザは選択できる
     - メールアドレスを確認になった場合WSでトンネルを作成
       - もし、WSが閉じられた場合はメールのURLから続きを開始する
       - WSが閉じられていない場合、メールのURLを踏むと「このウィンドウは消して元のウインドウに戻って」とダイアログ
  3. ワンタイムパスワードを入力し有効化
  4. 氏名、（プロフィール画像: 後々実装）、テーマ: ライトorダークを入力
- ログイン
  1. メールアドレス、パスワードを入力
  2. ワンタイムパスワードを入力
- アカウント
  - ロール: `admin`、`pro`、`user`の3種類
    - admin: 一人のみ。デプロイ時にenv `ADMIN_USER_EMAIL`と`ADMIN_TEMP_PW`を入力
      - envで設定したメアドとパスワードを入力するとパスワードを変更するダイアログとアカウント作成3以降
    - pro: SSOのAPI設定ができるユーザ。adminのダッシュボードから追加、削除できるようにする
    - user: アカウント作成、ログイン、SSOでのログイン、ログイン履歴閲覧（なんのSSOにログインしているか）、SSOログイン停止が可能
- SSO
  - ロール`admin`と`pro`が可能
  - ダッシュボードでリダイレクト先のURLとリダイレクト元URLを動的に追加できる

    ```ts
    {
        id: string,
        name: string,
        from_url: string[],
        to_url: string[],
        session_token_period: number,
        refresh_token_period: number,
        sso_public_key: string,
        sso_secret_key: string
    }
    ```

    - 初回作成時はどちらも最低1個入力する必要がある
    - 作成すると、PublicTokenとSecretTokenを作成する
      - PublicToken: アクセス元がSSOへリダイレクトする際に付与する。サービス認識用
      - SecretToken: SSOから返ってきたトークン（PASETO）を復号化するトークン。サーバーなどの利用者の目には見えないところで使用
