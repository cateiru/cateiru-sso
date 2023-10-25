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
