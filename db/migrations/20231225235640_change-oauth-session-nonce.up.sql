ALTER TABLE `oauth_session` ADD COLUMN `nonce` varchar(31) COLLATE utf8mb4_bin DEFAULT null AFTER `client_id`;
ALTER TABLE `oauth_session` DROP COLUMN `state`;
