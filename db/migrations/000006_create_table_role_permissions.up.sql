CREATE TABLE IF NOT EXISTS `role_permissions` (
	`id` INTEGER PRIMARY KEY AUTO_INCREMENT,
	`role_id` INTEGER NOT NULL,
	`permission_id` INTEGER NOT NULL,
	FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`),
	FOREIGN KEY (`permission_id`) REFERENCES `permissions`(`id`)
);