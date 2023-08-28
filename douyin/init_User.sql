DROP DATABASE IF EXISTS douyin;
CREATE DATABASE IF NOT EXISTS douyin ;

USE douyin;

-- Users table
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
		`user_id` bigint(20) NOT NULL,
    `username` VARCHAR(32) UNIQUE NOT NULL,
    `password` VARCHAR(256) NOT NULL, -- hashed password
    `name` VARCHAR(256),-- 删去了not null
    `avatar` VARCHAR(255),
    `background_image` VARCHAR(255),
    `signature` TEXT,
    `follow_count` BIGINT DEFAULT 0,
    `follower_count` BIGINT DEFAULT 0,
    `total_favorited` BIGINT DEFAULT 0,
    `work_count` BIGINT DEFAULT 0,
    `favorite_count` BIGINT DEFAULT 0,
    `token` VARCHAR(255) UNIQUE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

# 待添加功能备忘录：
# 登录密码长度要求