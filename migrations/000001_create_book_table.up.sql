CREATE TABLE IF NOT EXISTS `books` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `subtitle` VARCHAR(255),
    `isbn` VARCHAR(13),
    `authors` TEXT NOT NULL,
    `categories` TEXT NOT NULL,
    `language` VARCHAR(2) NOT NULL,
    `cover` VARCHAR(255),
    `publisher` VARCHAR(255),
    `published_at` DATE,
    `pages` INT NOT NULL,
    `read_pages` INT NOT NULL,
    `description` TEXT,
    `reading_status` VARCHAR(7),
    `edition` VARCHAR(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);