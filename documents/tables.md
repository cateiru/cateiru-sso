# 各テーブルの使用方法

[tables.go](../api/database/tables.go)

## Certification

認証テーブル

- メールアドレスをkeyとしてパスコード、ワンタイムパスワード、ワンタイムパスワードのバックアップコード、ユーザIDを保存
- 認証する場合はここを参照する

## MailCertification

メールアドレス認証用テーブル

- ユーザがはじめにアカウントを作成する際にメールアドレスの認証を行う。このときに、メールアドレスをパスコードを入力後このテーブルにmailTokenをkeyとして保存、該当ユーザにmailTokenをパラメータにもつメールを送信する
- このテーブルを参照することで認証されたかがわかる

## CreateAccountBuffer

メールアドレスの認証が済んでいるが、名前、その他ユーザ設定が完了してないユーザのデータの一時保管場所

- メールアドレス認証後のやつ
- `/create/verify`、`/create/accept`でPOSTのときにbufferTokenをkeyとして作成しトークンをcookieに入れる

## PWForget

パスワード忘れによる再登録用テーブル

- パスワード再発行は、該当メールアドレスに再設定用のURLを送付するメールアドレスの認証と似た仕組みを取る
- ForgetTokenをkeyにしてURLパラメータに付与してメールに送信する

## OnetimePassword

ワンタイムパスワード設定は、

1. secretとURLを作成する
2. URLをユーザに与えパスコードを生成してもらう
3. パスコードをsecretと検証し正しいか確認する
4. secretを保存

と、いうロジックをたどるためGET `/user/onetime`したときにこのテーブルにsecretを確認し、返ってきたパスコードをこのsecretで検証する。
URL内にsecret情報が含まれているため別にテーブルに保存する必要はないと思われるがユーザが独自にsecretを作って勝手にPOSTしてしまうのを防ぐ

## OnetimePasswordValidate

ログイン時にメールアドレスとパスコードを入力した後にワンタイムパスワードを求めるかどうかのやつ

## User

ユーザ情報

## UserLoginHistories

アカウントログイン履歴。OAuthでのログイン履歴も入る

## SSOLogins

現在、ログインしているSSO

## SessionInfo

当サイトのセッショントークン

## RefreshInfo

当サイトのリフレッシュトークン。セッションが変わるごとにセッショントークンと一緒に変更

## SSOService

SSOたちの情報。publicKeyがkeyになる

## SSOSession

SSOで取得できるセッショントークン

## SSORefreshToken

SSOで取得できるリフレッシュトークン

## WorkerLog

workerが走ったログ
