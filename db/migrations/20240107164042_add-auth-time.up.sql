ALTER TABLE `oauth_session` ADD COLUMN `auth_time` datetime NOT NULL DEFAULT current_timestamp AFTER `nonce`;
