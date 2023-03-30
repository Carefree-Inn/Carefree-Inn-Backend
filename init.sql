DROP DATABASE IF EXISTS  `Inn`;
CREATE DATABASE `Inn` CHARACTER SET=`utf8mb4`;

USE `Inn`;

CREATE TABLE `user`(
    `uuid` VARCHAR(20) NOT NULL PRIMARY KEY,
    `account` VARCHAR(15) NOT NULL UNIQUE KEY,
    `password` VARCHAR(15) NOT NULL,
    `username` VARCHAR(30) NOT NULL,
    `avatar` VARCHAR(100),
    `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `post` (
    `post_id` INT NOT NULL AUTO_INCREMENT,
    `uuid` VARCHAR(20) NOT NULL,
    `title` VARCHAR(200) NOT NULL,
    `content` TEXT,
    `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_time` DATETIME NULL DEFAULT NULL,
    PRIMARY KEY(`post_id`,`uuid`),
    FOREIGN KEY(`uuid`) REFERENCES `user`(`uuid`)
);

CREATE TABLE `tag` (
    `tag_id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `tag_name` VARCHAR(300) NOT NULL UNIQUE KEY,
    `reference` INT DEFAULT 0,
    `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_time` DATETIME NULL DEFAULT NULL
);

CREATE TABLE `post_tag`(
    `post_id` INT NOT NULL,
    `tag_id` INT NOT NULL,
    PRIMARY KEY(`post_id`,tag_id),
    FOREIGN KEY(`post_id`) REFERENCES `post`(`post_id`),
    FOREIGN KEY(`tag_id`) REFERENCES `tag`(`tag_id`)
);

CREATE TABLE `comment` (
    `comment_id` INT NOT NULL AUTO_INCREMENT,
    `uuid` VARCHAR(20) NOT NULL,
    `post_id` INT NOT NULL,
    `reply_comment` INT DEFAULT NULL,
    `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_time` DATETIME NULL DEFAULT NULL,
    PRIMARY KEY(`comment_id`,`uuid`,`post_id`),
    FOREIGN KEY(`uuid`) REFERENCES `user`(`uuid`),
    FOREIGN KEY(`post_id`) REFERENCES `post`(`post_id`),
    FOREIGN KEY(`reply_comment`) REFERENCES `comment`(`comment_id`)
);

CREATE TABLE `like` (
    `like_id` INT NOT NULL AUTO_INCREMENT,
    `uuid` VARCHAR(20) NOT NULL,
    `post_id` INT NOT NULL,
    `create_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`like_id`,`uuid`,`post_id`),
    FOREIGN KEY(`uuid`) REFERENCES `user`(`uuid`),
    FOREIGN KEY(`post_id`) REFERENCES `post`(`post_id`)
);