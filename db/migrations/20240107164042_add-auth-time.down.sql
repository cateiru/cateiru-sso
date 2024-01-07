ALTER TABLE `oauth_session` CHANGE COLUMN `nonce` `nonce` varchar(31) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT null;
ALTER TABLE `operation_history` CHANGE COLUMN `user_id` `user_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL;
ALTER TABLE `operation_history` CHANGE COLUMN `device` `device` varchar(31) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT null;
ALTER TABLE `operation_history` CHANGE COLUMN `os` `os` varchar(31) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT null;
ALTER TABLE `operation_history` CHANGE COLUMN `browser` `browser` varchar(31) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT null;
ALTER TABLE `oauth_session` DROP COLUMN `auth_time`;
