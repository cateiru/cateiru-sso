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
