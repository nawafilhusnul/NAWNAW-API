CREATE TABLE IF NOT EXISTS `users` (
    `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
    `email` VARCHAR(50) UNIQUE NOT NULL,
    `phone` VARCHAR(20) UNIQUE NOT NULL,
    `password` TEXT NOT NULL,
    `created_at` TIMESTAMP DEFAULT NOW(),
    `created_by` INTEGER NOT NULL,
    `updated_at` TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
    `updated_by` INTEGER NOT NULL,
    `is_deleted` BOOLEAN DEFAULT FALSE,
    `is_activated` BOOLEAN DEFAULT FALSE,
    `deleted_at` TIMESTAMP NULL,
    `deleted_by` INTEGER NULL
);