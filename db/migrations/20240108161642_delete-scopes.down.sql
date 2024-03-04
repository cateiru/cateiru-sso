ALTER TABLE `client_refresh` ADD COLUMN `scopes` json NOT NULL AFTER `client_id`;
