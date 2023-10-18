-- CateiruSSOのDDL

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
    -- CharGPT:
    -- スクリプトタグや変種タグが含まれる場合、BCP 47のタグの最大文字列長は最大で15文字程度になる可能性があります。
    -- これは、言語タグが2文字、地域タグが2文字、スクリプトタグが4文字、変種タグが最大で7文字まで許容されるためです。
    -- ただし、スクリプトタグや変種タグが必ずしもすべての場合に必要とされるわけではないため、必要に応じて長さを調整する必要があります。
    -- また、BCP 47のタグは将来的に変更や追加がある可能性があるため、データベースの設計には余裕を持たせておくことが望ましいです。
    `locale_id` VARCHAR(15) DEFAULT 'ja-JP' NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

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
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- ブランドテーブル
CREATE TABLE `brand` (
    `id` VARCHAR(32) NOT NULL,

    `name` TEXT NOT NULL,

    `description` TEXT DEFAULT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

CREATE TABLE `user_brand` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,

    `user_id` VARCHAR(32) NOT NULL,

    -- ブランドID
    `brand_id` VARCHAR(32) NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`brand_id`) REFERENCES `brand` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `brand_user_id` (`user_id`),
    INDEX `brand_brand_id` (`brand_id`),
    UNIQUE INDEX `brand_user_brand` (`user_id`, `brand_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- スタッフテーブル
-- ここに存在するユーザはスタッフとなる
CREATE TABLE `staff` (
    `user_id` VARCHAR(32) NOT NULL,

    -- メモ
    -- なぜスタッフなのかみたいなのを書くスペース
    `memo` TEXT DEFAULT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- ユーザー名保存
CREATE TABLE `user_name` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,

    `user_name` VARCHAR(15) NOT NULL,

    `user_id` VARCHAR(32) NOT NULL,

    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    INDEX `user_name_user_name_id_period` (`user_name`, `user_id`, `period`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- WebAuthnを保存するテーブル
CREATE TABLE `webauthn` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,

    `user_id` VARCHAR(32) NOT NULL,

    -- webauthn.Credentialのオブジェクト
    `credential` JSON NOT NULL,

    -- 登録した端末の情報
    `device` VARCHAR(31) DEFAULT NULL,
    `os` VARCHAR(31) DEFAULT NULL,
    `browser` VARCHAR(31) DEFAULT NULL,
    `is_mobile` BOOLEAN DEFAULT NULL,

    `ip` VARBINARY(16) NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `webauthn_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- パスワードを保存するテーブル
CREATE TABLE `password` (
    `user_id` VARCHAR(32) NOT NULL,

    `salt` VARBINARY(32) NOT NULL,
    `hash` VARBINARY(32) NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- アプリを使用したOTPを保存するテーブル
CREATE TABLE `otp` (
    `user_id` VARCHAR(32) NOT NULL,

    `secret` TEXT NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- OTPのバックアップコードを保存するテーブル
CREATE TABLE `otp_backup` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(32) NOT NULL,

    `code` VARCHAR(15) NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `otp_backup_user_id` (`user_id`),
    UNIQUE INDEX `otp_backup_user_code` (`user_id`, `code`)
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

    -- orgの招待を受けてアカウントを作成する場合に使用する
    `org_id` VARCHAR(32) DEFAULT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    UNIQUE INDEX `register_session_email` (`email`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- アプリを使用したOTPを新規に登録する際に使用するセッションテーブル
CREATE TABLE `register_otp_session` (
    `id` VARCHAR(31) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    `public_key` TEXT NOT NULL,
    `secret` TEXT NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- コードを入力した回数
    `retry_count` TINYINT UNSIGNED NOT NULL DEFAULT 0,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY(`id`),
    INDEX `register_otp_session_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- Emailを更新したときに確認に使用するテーブル
CREATE TABLE `email_verify_session` (
    `id` VARCHAR(31) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,
    `new_email` VARCHAR(255) NOT NULL,

    -- 認証コード
    `verify_code` CHAR(6) NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- コードを入力した回数
    `retry_count` TINYINT UNSIGNED NOT NULL DEFAULT 0,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY(`id`),
    INDEX `email_verify_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- パスワード再登録用
CREATE TABLE `reregistration_password_session` (
    `id` VARCHAR(31) NOT NULL,

    `email` VARCHAR(255) NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 疲労攻撃を回避するため有効期限と別のレコード削除期限を設ける
    `period_clear` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `completed` BOOLEAN NOT NULL DEFAULT 0,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

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
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

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
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY(`id`),
    INDEX `refresh_user_id` (`user_id`),
    UNIQUE INDEX `refresh_history_id` (`history_id`),
    UNIQUE INDEX `refresh_session_id` (`session_id`),
    INDEX `refresh_id_period` (`id`, `period`)
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
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

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
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY(`id`),
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

    -- スコープ
    `scopes` JSON NOT NULL,

    -- client_sessionのid
    `session_id` VARCHAR(31) NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY(`id`),
    INDEX `client_refresh_user_id` (`user_id`),
    INDEX `client_refresh_session_id` (`session_id`),
    INDEX `client_refresh_client_id` (`client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- WebAuthnを登録・ログインするときに保持する情報
-- webauthn.SessionData を元にしています
CREATE TABLE `webauthn_session` (
    -- ランダムにトークンを生成する
    `id` VARCHAR(31) NOT NULL,

    -- 紐付けられるユーザ
    `user_id` VARCHAR(32) DEFAULT NULL,

    -- webauthn.SessionData のjsonデータ
    `row` JSON NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 識別子
    `identifier` TINYINT NOT NULL DEFAULT 0,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(`id`),
    INDEX `webauthn_session_user_id` (`user_id`),
    INDEX `webauthn_session_identifier` (`identifier`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- パスワード変更時など、認証しているユーザーに対してさらに認証を求めるときのセッション
CREATE TABLE `certificate_session` (
    -- ランダムにトークンを生成する
    `id` VARCHAR(31) NOT NULL,

    `user_id` VARCHAR(32) NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 識別子
    `identifier` TINYINT NOT NULL DEFAULT 0,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY(`id`),
    INDEX `certificate_session_user_id` (`user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- ODICのクライアント
CREATE TABLE `client` (
    -- OAuth2.0のClient ID
    `client_id` VARCHAR(32) NOT NULL,

    -- クライアント名
    `name` VARCHAR(31) NOT NULL,
    -- 説明
    `description` TEXT DEFAULT NULL,
    -- クライアントのイメージ
    `image` TEXT DEFAULT NULL,

    -- orgで作成したものであれば、ここにorgのIDが入る
    -- 個人で作成したものはnullになる
    `org_id` VARCHAR(32) DEFAULT NULL,

    -- org_idが設定されている場合にこのフラグがtrueだと、orgのメンバーのみが利用できるようになる
    `org_member_only` BOOLEAN NOT NULL DEFAULT 0,

    -- ホワイトリストを使用するかどうか
    `is_allow` BOOLEAN NOT NULL DEFAULT 0,
    -- OAuthの認証リクエスト時にログイン、2faログインを求めることを強制する
    `prompt` ENUM('login', '2fa_login') DEFAULT NULL,

    -- 作成者
    `owner_user_id` VARCHAR(32) NOT NULL,

    -- OAuth2.0のClient Secret
    `client_secret` VARCHAR(63) NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`owner_user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`client_id`),
    INDEX `client_owner_user_id` (`owner_user_id`),
    INDEX `client_owner_user_id_client_id` (`owner_user_id`, `client_id`),
    INDEX `client_org_id` (`org_id`),
    INDEX `client_org_id_client_id` (`org_id`, `client_id`),
    INDEX `client_org_id_client_id_owner_user_id` (`org_id`, `client_id`, `owner_user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- クライアントのリダイレクトURLを設定するテーブル
CREATE TABLE `client_redirect` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,

    -- クライアントのID
    `client_id` VARCHAR(31) NOT NULL,

    `host` VARCHAR(128) NOT NULL,
    `url` TEXT NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `client_redirect_client_id` (`client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- クライアントのリファラーURLを設定するテーブル
CREATE TABLE `client_referrer` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,

    -- クライアントのID
    `client_id` VARCHAR(31) NOT NULL,

    `host` VARCHAR(128) NOT NULL,
    `url` TEXT NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `client_redirect_client_id` (`client_id`)
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

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY(`code`),
    INDEX `oauth_session_user_id` (`user_id`),
    INDEX `oauth_session_client_id` (`client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- SSOクライアントのスコープを保存するテーブル
CREATE TABLE `client_scope` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,

    `client_id` VARCHAR(31) NOT NULL,

    -- スコープ名
    -- ref. https://auth0.com/docs/get-started/apis/scopes/openid-connect-scopes
    `scope` VARCHAR(15) NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `client_scope_client_id` (`client_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- クライアントのis_allowが1のときのホワイトリストルール
CREATE TABLE `client_allow_rule` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,

    `client_id` VARCHAR(31) NOT NULL,

    -- user_idが指定されている場合、そのユーザのみを通過させる
    `user_id` VARCHAR(32) DEFAULT NULL,

    -- email_domainが指定されている場合、そのドメインと一致するユーザのみを通過させる
    `email_domain` VARCHAR(31) DEFAULT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `client_allow_rule_client_id` (`client_id`),
    UNIQUE INDEX `client_allow_rule_user_id_client_id` (`client_id`, `user_id`),
    UNIQUE INDEX `client_allow_rule_email_domain_client_id` (`client_id`, `email_domain`),
    UNIQUE INDEX `client_allow_rule_user_id_email_domain_client_id` (`client_id`, `user_id`, `email_domain`)
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
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `login_client_history_client_id` (`client_id`),
    INDEX `login_client_history_user_id` (`user_id`),
    INDEX `login_client_history_client_id_user_id` (`client_id`, `user_id`)
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
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

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

    -- 識別子
    `identifier` TINYINT NOT NULL DEFAULT 0,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

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
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

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
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`entry_id`) REFERENCES `broadcast_entry` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `broadcast_notice_entry_id` (`entry_id`),
    INDEX `broadcast_notice_user_id` (`user_id`),
    INDEX `broadcast_notice_user_id_is_read` (`user_id`, `is_read`),
    UNIQUE INDEX `broadcast_notice_entry_id_user_id` (`entry_id`, `user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- 組織
CREATE TABLE `organization` (
    `id` VARCHAR(32) NOT NULL,

    -- 組織名
    `name` VARCHAR(128) NOT NULL,
    `image` TEXT DEFAULT NULL,
    `link` TEXT DEFAULT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- 組織に所属するユーザー
CREATE TABLE `organization_user` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,

    `organization_id` VARCHAR(32) NOT NULL,
    `user_id` VARCHAR(32) NOT NULL,

    -- owner: 管理者。このユーザーは組織から脱退することができない。
    -- member: クライアントの作成・編集が可能なユーザー
    -- guest: クライアントにログインすることのみが可能なユーザー
    `role` ENUM('owner', 'member', 'guest') NOT NULL DEFAULT 'guest',

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`organization_id`) REFERENCES `organization` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY (`id`),
    INDEX `organization_user_organization_id` (`organization_id`),
    INDEX `organization_user_user_id` (`user_id`),
    UNIQUE INDEX `organization_user_organization_id_user_id` (`organization_id`, `user_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- Org招待メールのセッション
CREATE TABLE `invite_org_session` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,

    `token` VARCHAR(31) NOT NULL,

    -- 同一ユーザーに複数メールを送信できるようにUNIQUEではない
    `email` VARCHAR(255) NOT NULL,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- orgの招待の場合はorg_idが付与される
    -- 現状、orgの招待しかないのでNOT NULLにしている
    `org_id` VARCHAR(32) NOT NULL,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`org_id`) REFERENCES `organization` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY(`id`),
    UNIQUE INDEX `invite_org_session_token` (`token`),
    INDEX `invite_email_session_email` (`email`),
    INDEX `invite_email_session_org_id` (`org_id`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;

-- oauth 未ログインや `prompt=login` 時の一次保存セッション
CREATE TABLE `oauth_login_session` (
    `token` VARCHAR(31) NOT NULL,

    `client_id` VARCHAR(31) NOT NULL,

    `referrer_host` VARCHAR(128) DEFAULT NULL,

    `login_ok` BOOLEAN NOT NULL DEFAULT 0,

    -- 有効期限
    `period` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 管理用
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE,

    PRIMARY KEY(`token`)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ENGINE=InnoDB;
