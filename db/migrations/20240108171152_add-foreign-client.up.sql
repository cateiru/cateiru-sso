ALTER TABLE `client_refresh` ADD CONSTRAINT `client_refresh_ibfk_3` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `client_session` ADD KEY `client_id` (`client_id`);
ALTER TABLE `client_session` ADD CONSTRAINT `client_session_ibfk_2` FOREIGN KEY (`client_id`) REFERENCES `client` (`client_id`) ON DELETE CASCADE ON UPDATE CASCADE;
