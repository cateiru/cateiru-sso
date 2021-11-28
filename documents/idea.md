# Cateiru認証アイデア

## TL;DR

- 独自SSOクライアント
  - Google SSOみたいなかんじ
- Cateiruの作ったサービスなどで使えればいいなぁ

## 定義

### Routes

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
    - `?nr`でリダイレクトせず(no redirect)、トークンを表示（terminalなどでSSOする際用）
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
    - ~~`/create/onetime`~~ (アカウント作成後、設定から追加できるようにする)
      - GET: ワンタイムパスワードのトークン取得
    - `/create/info`
      - POST: ユーザ情報（名前、テーマ、プロフィール画像、~~ワンタイムパスワード~~）
      - もし、メール認証できてもこれが有効時間以内に送られなければユーザはリセットする
  - `/login`
    - POST: メールアドレス、パスワードを送信しcookieを作成
    - `/login/onetime`
      - ワンタイムパスワードを入力（必要な場合）
    - `/login/sso`
      - POST: メールアドレス、パスワードを送信しリダイレクトURLを返す
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
  - `/pro`
    - `/pro/sso`
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
      - `/user/password/forget`
        - パスワードを忘れてしまった際に使用
        - POST: メールアドレスを入力→登録されているメールアドレスがある場合そのメールアドレスにPW再登録用のURL送付
    - `/user/onetime`
      - ~~POST: ワンタイムパスワード変更~~
      - GET: ワンタイムパスワードのURL取得
      - POST: ワンタイムパスワードの無効化、有効化
      - get /create/onetimeでトークンを取得する必要あり
    - `/user/onetime/backup`
      - GET: ワンタイムパスワードのバックアップコードを取得
    - `/user/access`
      - GET: SSOログイン履歴取得
      - POST: ログアウト処理など
    - `/user/history`
      - GET: ログイン履歴取得
  - `/logout`
    - GET: ログアウト
    - DELETE: アカウント削除
  - `/oauth`
    - `/oauth/cert`
      - POST: セッショントークンでユーザ情報取得
    - `/oauth/update`
      - POST: リフレッシュトークンでセッショントークンを再取得

### トークン

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
- `buffer_token`
  - アカウント作成するときのバッファトークン
  - 有効時間付きでこの時間内にアカウント作成しないと作成はキャンセルされる

### 識別子

- `user_id`
  - ユーザごとに割り当てられるID
  - 重複なし
- `run_id`
  - workerのrun id
  - 特に意味はない
- `forget_token`
  - パスワード忘れ時の再登録トークン

### テーブル

- 認証

    ```ts
    {
        mail: string,
        password: string, // パスワードはハッシュ化
        create_account_date: Date, // 認証後のdate
        user_id: string,

        onetime_password_secret?: string,
        onetime_password_backups?: string[]
    }
    ```

- メールアドレス認証
  - wsは、このエンティティを作成し、メールを送信する。
  - 作成時`open_new_window`はfalseでwsが閉じられたときtrueにする。`verify`はfalse
  - メール側リンクは`open_new_window`がtrueである場合そのページで続きを開始しこのエンティティを削除。falseの場合は`verify`をtrueにする（エンティティは削除しない）
  - wsで`verify`がtrueになったのを確認したらそのページで続きを開始しエンティティを削除、wsを閉じる

    どちらも、認証でき次第`認証`のテーブルにユーザを追加

    ```ts
    {
        mail_token: string, // メールのURLパラメータに付加するトークン
        create_date: Date,  // メール認証開始時間
        period_minute: number = 30, //メール認証の有効期限
        open_new_window: boolean, // そのままのウインドウで続きをやるか(false)メールのリンク先ウインドウからやるか(true)
        verify: boolean, // 認証されているか。wsでこのテーブルを読むときに確認する部分
        change_mail_mode: boolean, // メールアドレス変更時のメールアドレス認証用に使用しているか。もしその場合passwordは空

        mail: string,
        password: string, // パスワードはハッシュ化
    }
    ```

- アカウント作成Buffer

    ```ts
    {
        buffer_token: string,

        mail: string,
        password: string,

        create_date: Date,
        period_minute: int = 30;
    }
    ```

- ワンタイムパスワード認証Buffer

    ```ts
    {
        onetime_token: string,

        onetime_password_secret: string,
        user_id: string,
        mail: string,

        create_date: Date,
        period_minute: int = 30;
    }
    ```

- User情報
  - SSOで暗号化して送る情報はこれ

    ```ts
    {
        user_id: string,

        first_name: string,
        last_name: string,
        role: string, // `admin`. `pro` or `user`
        mail: string, // これがないと`認証`テーブルを検索できなくなる

        theme: string,
        avatar_url?: string,
    }
    ```

- Userログイン情報
  - SSO含めこのサイトへログインしたときに履歴を残す

    ```ts
    {
        user_id: string,

        histories: {
            access_id: string,
            date: Date,
            ip_address: string,
            is_sso: boolean,
            sso_public?: string,
        }[]
    }
    ```

- セッショントークン
  - セッションを開始したときにエンティティを作成
  - workerで定期的にperiod_hourを過ぎたエンティティを削除

    ```ts
    {
        session_token: string,
        user_id: string,
        create_date: Date,
        period_hour: number = 6,
    }
    ```

- リフレッシュトークン
  - セッショントークンを更新する場合: 前のセッショントークンを削除して新しいセッショントークンを作成。リフレッシュトークンを新しいものに更新
  - workerで定期的にperiod_hourを過ぎたエンティティを削除

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
        sso_public_key: string, // key

        name: string,
        from_url: string[],
        to_url: string[],
        login_only: boolean, // ログインのみモード: トークンは送信しない

        session_token_period?: number,
        refresh_token_period?: number,

        sso_secret_key: string
    }
    ```

- ユーザのSSO

    ```ts
    {
        user_id: string,
        sso: string[], // public_keyのlist
    }
    ```

- SSOセッショントークン
  - `login_only`がtrueの場合、このトークンを作成して一緒に暗号化して返す

    ```ts
    {
        sso_session_token: string,
        create_date: Date,
        period_hour: number;
        user_id: string,
    }
    ```

- SSOリフレッシュトークン
  - `login_only`がfalseの場合、このトークンを作成して一緒に暗号化して返す

    ```ts
    {
        sso_refresh_token: string,
        sso_session_token: string,

        create_date: Date,
        period_hour: number;
        user_id: string;
    }
    ```

- SSOのログイン履歴

    ```ts
    {
        user_id: string

        sso_refresh_tokens: string[]
    }
    ```

- Worker ログ

    ```ts
    {
        run_id: string,
        status: number, // 0は正常終了、その他はエラー
        status_message?: string,
        run_date: Date,
    }
    ```

- PW忘れ再登録用トークン

    ```ts
    {
        forget_token: string,

        mail: string,
        create_date: Date,
        period_minute: number = 10,
    }
    ```

- ワンタイムパスワード登録用

    ```ts
    {
        onetime_public_key: string,
        onetime_private_key: string.

        create_date: Date,
        period_minute: number = 10,
        user_id: string,
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
       - `buffer_token`をcookieに入れる
  3. 氏名、（プロフィール画像: 後々実装）、テーマ: ライトorダーク、ワンタイムパスワードを入力し`buffer_token`で認証しアカウント作成
- ログイン
  1. メールアドレス、パスワードを入力
  2. ワンタイムパスワードを入力
- アカウント
  - ロール: `admin`、`pro`、`user`の3種類
    - admin: 一人のみ。デプロイ時にenv `ADMIN_USER_EMAIL`と`ADMIN_TEMP_PW`、`ADMIN_ONETIME_PW`を入力
      - envで設定したメアドとパスワードを入力するとパスワードを変更するダイアログとアカウント作成3以降
    - pro: SSOのAPI設定ができるユーザ。adminのダッシュボードから追加、削除できるようにする
    - user: アカウント作成、ログイン、SSOでのログイン、ログイン履歴閲覧（なんのSSOにログインしたか）、SSOログイン停止が可能
- SSO
  - ロール`admin`と`pro`が可能
  - ダッシュボードでリダイレクト先のURLとリダイレクト元URLを動的に追加できる
    - 初回作成時はどちらも最低1個入力する必要がある
    - 作成すると、PublicTokenとSecretTokenを作成する
      - PublicToken: アクセス元がSSOへリダイレクトする際に付与する。サービス認識用
      - SecretToken: SSOから返ってきたトークン（PASETO）を復号化するトークン。サーバーなどの利用者の目には見えないところで使用
