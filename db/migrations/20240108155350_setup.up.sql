CREATE TABLE `brand` (
  `id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `name` text COLLATE utf8mb4_bin NOT NULL,
  `description` text COLLATE utf8mb4_bin,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `broadcast_entry` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `create_user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `title` text COLLATE utf8mb4_bin NOT NULL,
  `body` text COLLATE utf8mb4_bin,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `broadcast_entry_create_user_id` (`create_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `broadcast_notice` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `entry_id` int unsigned NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `is_read` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `broadcast_notice_entry_id_user_id` (`entry_id`,`user_id`),
  KEY `broadcast_notice_entry_id` (`entry_id`),
  KEY `broadcast_notice_user_id` (`user_id`),
  KEY `broadcast_notice_user_id_is_read` (`user_id`,`is_read`),
  CONSTRAINT `broadcast_notice_ibfk_1` FOREIGN KEY (`entry_id`) REFERENCES `broadcast_entry` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `broadcast_notice_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `certificate_session` (
  `id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `identifier` tinyint NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `certificate_session_user_id` (`user_id`),
  CONSTRAINT `certificate_session_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `client` (
  `client_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `description` text COLLATE utf8mb4_bin,
  `image` text COLLATE utf8mb4_bin,
  `org_id` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `org_member_only` tinyint(1) NOT NULL DEFAULT '0',
  `is_allow` tinyint(1) NOT NULL DEFAULT '0',
  `prompt` enum('login','2fa_login') COLLATE utf8mb4_bin DEFAULT NULL,
  `owner_user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `client_secret` varchar(63) COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`client_id`),
  KEY `client_owner_user_id` (`owner_user_id`),
  KEY `client_owner_user_id_client_id` (`owner_user_id`,`client_id`),
  KEY `client_org_id` (`org_id`),
  KEY `client_org_id_client_id` (`org_id`,`client_id`),
  KEY `client_org_id_client_id_owner_user_id` (`org_id`,`client_id`,`owner_user_id`),
  CONSTRAINT `client_ibfk_1` FOREIGN KEY (`owner_user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `client_allow_rule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `email_domain` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `client_allow_rule_user_id_client_id` (`client_id`,`user_id`),
  UNIQUE KEY `client_allow_rule_email_domain_client_id` (`client_id`,`email_domain`),
  UNIQUE KEY `client_allow_rule_user_id_email_domain_client_id` (`client_id`,`user_id`,`email_domain`),
  KEY `client_allow_rule_client_id` (`client_id`),
  CONSTRAINT `client_allow_rule_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `client_redirect` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `host` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `url` text COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `client_redirect_client_id` (`client_id`),
  CONSTRAINT `client_redirect_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `client_referrer` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `host` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `url` text COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `client_redirect_client_id` (`client_id`),
  CONSTRAINT `client_referrer_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `client_refresh` (
  `id` varchar(63) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `client_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `scopes` json NOT NULL,
  `session_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `client_refresh_user_id` (`user_id`),
  KEY `client_refresh_session_id` (`session_id`),
  KEY `client_refresh_client_id` (`client_id`),
  CONSTRAINT `client_refresh_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `client_refresh_ibfk_2` FOREIGN KEY (`session_id`) REFERENCES `client_session` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `client_scope` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `scope` varchar(15) COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `client_scope_client_id` (`client_id`),
  CONSTRAINT `client_scope_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `client_session` (
  `id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `client_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `client_session_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `email_verify_session` (
  `id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `new_email` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `verify_code` char(6) COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `retry_count` tinyint unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `email_verify_user_id` (`user_id`),
  CONSTRAINT `email_verify_session_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `invite_org_session` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `token` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `org_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `invite_org_session_token` (`token`),
  KEY `invite_email_session_email` (`email`),
  KEY `invite_email_session_org_id` (`org_id`),
  CONSTRAINT `invite_org_session_ibfk_1` FOREIGN KEY (`org_id`) REFERENCES `organization` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `login_client_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `device` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `os` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `browser` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `is_mobile` tinyint(1) DEFAULT NULL,
  `ip` varbinary(16) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `login_client_history_client_id` (`client_id`),
  KEY `login_client_history_user_id` (`user_id`),
  KEY `login_client_history_client_id_user_id` (`client_id`,`user_id`),
  CONSTRAINT `login_client_history_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `login_client_history_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `login_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `refresh_id` varbinary(16) NOT NULL,
  `device` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `os` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `browser` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `is_mobile` tinyint(1) DEFAULT NULL,
  `ip` varbinary(16) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `login_history_user_id` (`user_id`),
  KEY `login_history_refresh_id` (`refresh_id`),
  CONSTRAINT `login_history_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `login_try_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `device` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `os` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `browser` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `is_mobile` tinyint(1) DEFAULT NULL,
  `ip` varbinary(16) NOT NULL,
  `identifier` tinyint NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `login_try_history_user_id` (`user_id`),
  CONSTRAINT `login_try_history_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `oauth_login_session` (
  `token` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `client_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `referrer_host` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL,
  `login_ok` tinyint(1) NOT NULL DEFAULT '0',
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`token`),
  KEY `client_id` (`client_id`),
  CONSTRAINT `oauth_login_session_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `oauth_session` (
  `code` varchar(63) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `client_id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `nonce` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `auth_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`code`),
  KEY `oauth_session_user_id` (`user_id`),
  KEY `oauth_session_client_id` (`client_id`),
  CONSTRAINT `oauth_session_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `oauth_session_ibfk_2` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `operation_history` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `device` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `os` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `browser` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `is_mobile` tinyint(1) DEFAULT NULL,
  `ip` varbinary(16) NOT NULL,
  `identifier` tinyint NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `operation_history_user_id` (`user_id`),
  CONSTRAINT `operation_history_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `organization` (
  `id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `image` text COLLATE utf8mb4_bin,
  `link` text COLLATE utf8mb4_bin,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `organization_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `organization_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `role` enum('owner','member','guest') COLLATE utf8mb4_bin NOT NULL DEFAULT 'guest',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `organization_user_organization_id_user_id` (`organization_id`,`user_id`),
  KEY `organization_user_organization_id` (`organization_id`),
  KEY `organization_user_user_id` (`user_id`),
  CONSTRAINT `organization_user_ibfk_1` FOREIGN KEY (`organization_id`) REFERENCES `organization` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `organization_user_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `otp` (
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `secret` text COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `otp_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `otp_backup` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `code` varchar(15) COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `otp_backup_user_code` (`user_id`,`code`),
  KEY `otp_backup_user_id` (`user_id`),
  CONSTRAINT `otp_backup_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `otp_session` (
  `id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `retry_count` tinyint unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `otp_session_user_id` (`user_id`),
  CONSTRAINT `otp_session_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `password` (
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `salt` varbinary(32) NOT NULL,
  `hash` varbinary(32) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `password_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `refresh` (
  `id` varchar(63) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `history_id` varbinary(16) NOT NULL DEFAULT (uuid_to_bin(uuid())),
  `session_id` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `refresh_history_id` (`history_id`),
  UNIQUE KEY `refresh_session_id` (`session_id`),
  KEY `refresh_user_id` (`user_id`),
  KEY `refresh_id_period` (`id`,`period`),
  CONSTRAINT `refresh_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `register_otp_session` (
  `id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `public_key` text COLLATE utf8mb4_bin NOT NULL,
  `secret` text COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `retry_count` tinyint unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `register_otp_session_user_id` (`user_id`),
  CONSTRAINT `register_otp_session_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `register_session` (
  `id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `email_verified` tinyint(1) NOT NULL DEFAULT '0',
  `send_count` tinyint unsigned NOT NULL DEFAULT '1',
  `verify_code` char(6) COLLATE utf8mb4_bin NOT NULL,
  `retry_count` tinyint unsigned NOT NULL DEFAULT '0',
  `org_id` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `register_session_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `reregistration_password_session` (
  `id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `period_clear` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `completed` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `reregistration_password_session_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `session` (
  `id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `session_user_id` (`user_id`),
  CONSTRAINT `session_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `setting` (
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `notice_email` tinyint(1) NOT NULL DEFAULT '0',
  `notice_webpush` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `setting_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `staff` (
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `memo` text COLLATE utf8mb4_bin,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `staff_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `user` (
  `id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `user_name` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `family_name` text COLLATE utf8mb4_bin,
  `middle_name` text COLLATE utf8mb4_bin,
  `given_name` text COLLATE utf8mb4_bin,
  `gender` char(1) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `birthdate` date DEFAULT NULL,
  `avatar` text COLLATE utf8mb4_bin,
  `locale_id` varchar(15) COLLATE utf8mb4_bin NOT NULL DEFAULT 'ja-JP',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_user_name` (`user_name`),
  UNIQUE KEY `user_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `user_brand` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `brand_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `brand_user_brand` (`user_id`,`brand_id`),
  KEY `brand_user_id` (`user_id`),
  KEY `brand_brand_id` (`brand_id`),
  CONSTRAINT `user_brand_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_brand_ibfk_2` FOREIGN KEY (`brand_id`) REFERENCES `brand` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `user_name` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_name_user_name_id_period` (`user_name`,`user_id`,`period`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `webauthn` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `credential` json NOT NULL,
  `device` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `os` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `browser` varchar(31) COLLATE utf8mb4_bin DEFAULT NULL,
  `is_mobile` tinyint(1) DEFAULT NULL,
  `ip` varbinary(16) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `webauthn_user_id` (`user_id`),
  CONSTRAINT `webauthn_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
CREATE TABLE `webauthn_session` (
  `id` varchar(31) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `row` json NOT NULL,
  `period` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `identifier` tinyint NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `webauthn_session_user_id` (`user_id`),
  KEY `webauthn_session_identifier` (`identifier`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
