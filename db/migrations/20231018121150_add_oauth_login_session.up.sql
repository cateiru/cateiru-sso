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