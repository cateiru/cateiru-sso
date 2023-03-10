-- TODO:
-- - パスポートハッシュなどの型

-- ユーザテーブル
CREATE TABLE `user` (
    -- ユーザIDはULIDを使用して一意にする
    `id` VARCHAR(32) NOT NULL,

    -- ユーザ名はユーザごとに一意なIDとなる
    -- ログイン時にメールアドレスの代替としてログインできる
    -- アカウント登録時に、デフォルトはUUIDからランダムな文字列を作って入れる
    -- 検索する際にはutf8mb4_general_ciのコレクションを使用する
    `user_name` VARCHAR(15) NOT NULL DEFAULT (LEFT(UUID(), 8)) COLLATE utf8_general_ci,

    -- Email
    `email` VARCHAR(255) NOT NULL,

    -- 名前
    `family_name` TEXT DEFAULT NULL,
    `middle_name` TEXT DEFAULT NULL,
    `given_name` TEXT DEFAULT NULL,

    -- 性別
    -- 0: 不明、1: 男性、2: 女性、9: 適用不能
    `gender` CHAR(1) NOT NULL DEFAULT '0',

    `birthdate` DATE DEFAULT NULL,
    `avatar` TEXT DEFAULT NULL,

    -- ロケールID
    -- デフォルトは日本(ja_JP)
    `locale_id` CHAR(5) DEFAULT 'ja_JP' NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE INDEX `user_user_name` (`user_name`),
    UNIQUE INDEX `user_email` (`email`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- ユーザ設定テーブル
CREATE TABLE `setting` (
    `user_id` VARCHAR(32) NOT NULL,

    -- 通知設定
    `notice_email` BOOLEAN NOT NULL DEFAULT 0,
    `notice_webpush` BOOLEAN NOT NULL DEFAULT 0,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- ブランドテーブル
CREATE TABLE `brand` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(32) NOT NULL,

    -- ブランド名
    `brand` VARCHAR(31) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `brand_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- スタッフテーブル
-- ここに存在するユーザはスタッフとなる
CREATE TABLE `staff` (
    `user_id` VARCHAR(32) NOT NULL,

    -- メモ
    -- なぜスタッフなのかみたいなのを書くスペース
    `memo` TEXT DEFAULT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- passkeyを保存するテーブル
CREATE TABLE `passkey` (
    `user_id` VARCHAR(32) NOT NULL,

    -- WebauthnID
    -- このIDを使用してpasskeyを認証する
    -- 64byteのランダムな文字列
    `webauthn_user_id` VARBINARY(64) NOT NULL,

    -- webauthn.Credentialのオブジェクト
    `credential` JSON NOT NULL,

    -- authenticatorData.flagsのBackupState値
    -- これが1の場合はpasskeyが複数デバイス感で共有される可能性がある
    -- ref. https://www.docswell.com/s/ydnjp/KWDLDZ-2022-10-14-141235#p20
    `flag_backup_state` BOOLEAN NOT NULL DEFAULT 0,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- iCloudのPasskeyなどは複数のApple端末で共有できるため、
-- Passkeyでログインした端末を記録しておくためのテーブル
CREATE TABLE `passkey_login_device` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(32) NOT NULL,

    -- 使用した端末のUA
    `device` VARCHAR(31) DEFAULT NULL,
    `os` VARCHAR(31) DEFAULT NULL,
    `browser` VARCHAR(31) DEFAULT NULL,
    `is_mobile` BOOLEAN DEFAULT NULL,

    -- passkeyを登録したデバイスかどうか
    `is_register_device` BOOLEAN NOT NULL DEFAULT 0,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `passkey_login_device_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- パスポートを保存するテーブル
CREATE TABLE `password` (
    `user_id` VARCHAR(32) NOT NULL,

    `salt` VARBINARY(32) NOT NULL,
    `hash` VARBINARY(32) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- アプリを使用したOTPを保存するテーブル
CREATE TABLE `otp` (
    `user_id` VARCHAR(32) NOT NULL,

    -- TODO: サイズの最適化をしたい
    `secret` VARCHAR(31),

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- OTPのバックアップコードを保存するテーブル
CREATE TABLE `otp_backup` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(32) NOT NULL,

    `code` VARCHAR(15) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `otp_backup_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- アカウント登録時に使用するセッションを保存するテーブル
CREATE TABLE `register_session` (
    `id` VARCHAR(31) NOT NULL,

    `email` VARCHAR(255) NOT NULL,
    `email_verified` BOOLEAN NOT NULL DEFAULT 0,

    -- メール送信回数
    `send_count` TINYINT UNSIGNED NOT NULL DEFAULT 1,

    -- 認証コード
    `verify_code` CHAR(6) NOT NULL,

    -- コードを入力した回数
    `retry_count` TINYINT UNSIGNED NOT NULL DEFAULT 0,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    UNIQUE INDEX `register_session_email` (`email`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- アプリを使用したOTPを新規に登録する際に使用するセッションテーブル
CREATE TABLE `register_otp_session` (
    `id` VARCHAR(31) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- TODO: 文字サイズを最適化したい
    `public_key` TEXT NOT NULL,
    `secret` TEXT NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- コードを入力した回数
    `retry_count` TINYINT UNSIGNED NOT NULL DEFAULT 0,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    INDEX `register_otp_session_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- Emailを更新したときに確認に使用するテーブル
CREATE TABLE `email_verify_session` (
    `id` VARCHAR(31) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- 認証コード
    `verify_code` CHAR(6) NOT NULL,

    -- メール送信回数
    `send_count` TINYINT UNSIGNED NOT NULL DEFAULT 1,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- コードを入力した回数
    `retry_count` TINYINT UNSIGNED NOT NULL DEFAULT 0,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    INDEX `email_verify_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- パスワード再登録用
CREATE TABLE `reregistration_password_session` (
    `id` VARCHAR(31) NOT NULL,

    `email` VARCHAR(255) NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    INDEX `reregistration_password_session_email` (`email`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- セッション維持用のテーブル
CREATE TABLE `session` (
    -- ランダムにトークンを生成する
    `id` VARCHAR(31) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    INDEX `session_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- セッショントークンを更新するためのリフレッシュトークン用テーブル
-- 同時ログインでは、このトークンのみcookieに入れっぱなしにしておく
CREATE TABLE `refresh` (
    -- ランダムにトークンを生成する
    `id` VARCHAR(63) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- ログイン履歴と紐づけるためのID
    `history_id` VARBINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),

    -- sessionのid
    -- 複数ログインを可能にするためNULLABLE
    `session_id` VARCHAR(31) DEFAULT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    INDEX `refresh_user_id` (`user_id`),
    UNIQUE INDEX `refresh_history_id` (`history_id`),
    UNIQUE INDEX `refresh_session_id` (`session_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- パスワードによる認証は成功して次にOTPを求める場合のセッションを保存するテーブル
CREATE TABLE `otp_session` (
    `id` VARCHAR(31) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- コードを入力した回数
    `retry_count` TINYINT UNSIGNED NOT NULL DEFAULT 0,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    INDEX `otp_session_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- SSOクライアントのセッショントークンテーブル
CREATE TABLE `client_session` (
    -- ランダムにトークンを生成する
    `id` VARCHAR(31) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- クライアントのID
    `client_id` VARCHAR(31) NOT NULL,
    `login_client_id` INT UNSIGNED NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    -- TODO: もっと突き詰める
    INDEX `client_session_user_id` (`user_id`),
    INDEX `client_session_client_id` (`client_id`),
    INDEX `client_session_login_client_id` (`login_client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- SSOクライアントのリフレッシュトークンテーブル
CREATE TABLE `client_refresh` (
    -- ランダムにトークンを生成する
    `id` VARCHAR(63) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- クライアントのID
    `client_id` VARCHAR(31) NOT NULL,
    `login_client_id` INT UNSIGNED NOT NULL,

    -- client_sessionのid
    `session_id` VARCHAR(31) NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    -- TODO: もっと突き詰める
    INDEX `client_refresh_user_id` (`user_id`),
    INDEX `client_refresh_session_id` (`session_id`),
    INDEX `client_refresh_client_id` (`client_id`),
    INDEX `client_refresh_login_client_id` (`login_client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- WebAuthnを登録・ログインするときに保持する情報
-- webauthn.SessionData を元にしています
CREATE TABLE `webauthn_session` (
    -- ランダムにトークンを生成する
    `id` VARCHAR(31) NOT NULL,

    -- 紐付けられるユーザ
    -- ログインの場合ではこれにユーザIDが入る
    `user_id` VARCHAR(32) DEFAULT NULL,

    -- WebAuthnID
    -- いわゆるユーザID
    `webauthn_user_id` VARBINARY(64) NOT NULL,
    `user_display_name` TEXT NOT NULL,

    -- チャレンジ
    -- チャレンジは（16 バイト以上の）ランダムな情報のバッファーであること
    `challenge` TEXT NOT NULL,

    `user_verification` VARCHAR(15) NOT NULL DEFAULT 'preferred',

    -- webauthn.SessionData のjsonデータ
    `row` JSON NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    INDEX `webauthn_register_session_webauthn_user_id` (`webauthn_user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- OAuthで接続したときのセッション
-- nonceとかを保存しておく
CREATE TABLE `oauth_session` (
    `code` VARCHAR(63) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- クライアントのID
    `client_id` VARCHAR(31) NOT NULL,

    -- CSRF, XSRFで使用される`state`を格納するやつ
    `state` VARCHAR(31) DEFAULT NULL,
    `nonce` VARCHAR(31) DEFAULT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`code`),
    INDEX `oauth_session_user_id` (`user_id`),
    INDEX `oauth_session_client_id` (`client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- ODICのクライアント
CREATE TABLE `client` (
    -- OAuth2.0のClient ID
    -- 公開されるIDであり、それ単体では使用できないのでv1で良い
    `client_id` VARCHAR(31) NOT NULL,

    -- クライアント名
    `name` VARCHAR(31) NOT NULL,
    -- 説明
    `description` TEXT DEFAULT NULL,
    -- クライアントのイメージ
    `image` TEXT DEFAULT NULL,

    -- ホワイトリストを使用するかどうか
    `is_allow` BOOLEAN NOT NULL DEFAULT 0,
    -- 二段階認証をしたユーザのみに限定するかどうか
    `indispensable_2fa` BOOLEAN NOT NULL DEFAULT 0,
    -- prompt_loginが1の場合、OAuthの認証リクエスト時にログイン、クイズを求めることを強制する
    `prompt` ENUM('login', 'quiz') DEFAULT NULL,

    `owner_user_id` VARCHAR(32) NOT NULL,

    -- OAuth2.0のClient Secret
    `client_secret` VARCHAR(63) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`client_id`),
    INDEX `client_owner_user_id` (`owner_user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- SSOクライアントのスコープを保存するテーブル
CREATE TABLE `client_scope` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,

    `client_id` VARCHAR(31) NOT NULL,

    -- スコープ名
    -- ref. https://auth0.com/docs/get-started/apis/scopes/openid-connect-scopes
    `scope` VARCHAR(15) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `client_scope_client_id` (`client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- ログインしたSSOクライアントのスコープを保存するテーブル
-- クライアントが途中でスコープを変更してもログイン履歴にはログイン時に求めたスコープとなる
CREATE TABLE `login_client_scope` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,

    `login_client_id` INT UNSIGNED NOT NULL,

    -- スコープ名
    -- ref. https://auth0.com/docs/get-started/apis/scopes/openid-connect-scopes
    `scope` VARCHAR(15) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `login_client_scope_login_client_id` (`login_client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- クライアントのis_allowが1のときのホワイトリストルール
CREATE TABLE `client_allow_rule` (
    `client_id` VARCHAR(31) NOT NULL,

    -- user_idが指定されている場合、そのユーザのみを通過させる
    `user_id` VARCHAR(32) DEFAULT NULL,

    -- email_domainが指定されている場合、そのドメインと一致するユーザのみを通過させる
    `email_domain` VARCHAR(31) DEFAULT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- クライアントの認証方法でrequire_quizが1の場合のクイズ
-- reCAPTCHAのようなやつ
-- 転売対策とかに有効
CREATE TABLE `client_quiz` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `client_id` VARCHAR(31) NOT NULL,

    `title` TEXT NOT NULL,
    -- 複数の回答方法や選択肢がある場合があるので
    -- 答えは正規表現で記述する
    `answer_regexp` TEXT NOT NULL,

    -- 選択肢
    `choices` JSON DEFAULT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `client_quiz_client_id` (`client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- 現在ログインしているSSOクライアントテーブル
CREATE TABLE `login_client` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `client_id` VARCHAR(31) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- 過去にログインしたSSOクライアントのテーブル
CREATE TABLE `login_client_history` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `client_id` VARCHAR(31) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- 使用した端末のUA
    `device` VARCHAR(31) DEFAULT NULL,
    `os` VARCHAR(31) DEFAULT NULL,
    `browser` VARCHAR(31) DEFAULT NULL,
    `is_mobile` BOOLEAN DEFAULT NULL,

    -- INET6_ATON、INET6_NTOAを使用して格納する
    `ip` VARBINARY(16) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `login_client_history_client_id` (`client_id`),
    INDEX `login_client_history_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- ログイン履歴
CREATE TABLE `login_history` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(32) NOT NULL,

    -- リフレッシュトークンと紐づけたID
    -- refreshテーブルに参照することでユーザがどの端末でログインしているかを調べることができる
    `refresh_id` VARBINARY(16) NOT NULL,

    -- 使用した端末のUA
    `device` VARCHAR(31) DEFAULT NULL,
    `os` VARCHAR(31) DEFAULT NULL,
    `browser` VARCHAR(31) DEFAULT NULL,
    `is_mobile` BOOLEAN DEFAULT NULL,

    -- INET6_ATON、INET6_NTOAを使用して格納する
    `ip` VARBINARY(16) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `login_history_user_id` (`user_id`),
    INDEX `login_history_refresh_id` (`refresh_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- ログインを試みた履歴
CREATE TABLE `login_try_history` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(32) NOT NULL,

    -- 使用した端末のUA
    `device` VARCHAR(31) DEFAULT NULL,
    `os` VARCHAR(31) DEFAULT NULL,
    `browser` VARCHAR(31) DEFAULT NULL,
    `is_mobile` BOOLEAN DEFAULT NULL,

    -- INET6_ATON、INET6_NTOAを使用して格納する
    `ip` VARBINARY(16) NOT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `login_try_history_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- 全ユーザー一斉通知用のエントリ
CREATE TABLE `broadcast_entry` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `create_user_id` VARCHAR(32) NOT NULL,

    `title` TEXT NOT NULL,
    `body` TEXT DEFAULT NULL,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `broadcast_entry_create_user_id` (`create_user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- 全ユーザー一斉通知のユーザごとの既読状況を保存するテーブル
CREATE TABLE `broadcast_notice` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `entry_id` INT UNSIGNED NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- 既読状況
    `is_read` BOOLEAN NOT NULL DEFAULT 0,

    -- 管理用
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `broadcast_notice_entry_id` (`entry_id`),
    INDEX `broadcast_notice_user_id` (`user_id`),
    INDEX `broadcast_notice_user_id_is_read` (`user_id`, `is_read`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;
