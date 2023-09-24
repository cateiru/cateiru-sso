ALTER TABLE `oauth_session` ADD COLUMN `nonce` VARCHAR(31) DEFAULT NULL AFTER `period`;
