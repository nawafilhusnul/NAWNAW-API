CREATE TABLE IF NOT EXISTS `user_roles` (
	`id` INTEGER PRIMARY KEY AUTO_INCREMENT,
	`user_id` INTEGER NOT NULL,
	`role_id` INTEGER NOT NULL,
	FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
	FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`)
);