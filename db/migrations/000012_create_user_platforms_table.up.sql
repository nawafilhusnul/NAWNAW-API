CREATE TABLE IF NOT EXISTS `user_platforms` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT NOT NULL,
  `platform_id` INT NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`platform_id`) REFERENCES `platforms` (`id`)
);