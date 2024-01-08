ALTER TABLE `oauth_session` ADD COLUMN `nonce` varchar(31) DEFAULT null AFTER `client_id`;
ALTER TABLE `oauth_session` DROP COLUMN `state`;
