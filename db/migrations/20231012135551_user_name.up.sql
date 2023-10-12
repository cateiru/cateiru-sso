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
