DROP DATABASE IF EXISTS uestc_acm_training;
CREATE SCHEMA IF NOT EXISTS uestc_acm_training DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
USE uestc_acm_training;
DROP TABLE IF EXISTS `configurations`;
CREATE TABLE IF NOT EXISTS `configurations` (
  `key` varchar(128) PRIMARY KEY,
  `value` varchar(128)
) ENGINE = InnoDB;
DROP TABLE IF EXISTS `seasons`;
CREATE TABLE IF NOT EXISTS `seasons` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL
) ENGINE = InnoDB;
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `season_id` BIGINT NOT NULL,
  `name` VARCHAR(24) NOT NULL UNIQUE,
  `email` VARCHAR(100) NOT NULL UNIQUE,
  `phone` VARCHAR(45) NOT NULL UNIQUE,
  -- sha256sum
  `password` VARCHAR(130) NOT NULL UNIQUE,
  `rating` INT NOT NULL DEFAULT 1500
) ENGINE = InnoDB;
CREATE TABLE IF NOT EXISTS `contests` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `season_id` BIGINT NOT NULL,
  `name` VARCHAR(128) NOT NULL,
  `description` VARCHAR(128) DEFAULT NULL,
  `url` VARCHAR(1024) DEFAULT NULL COMMENT 'the contest link, if it is online',
  `start_time` BIGINT NOT NULL COMMENT 'the start time of the contest, as unix timestamp',
  `duration` BIGINT NOT NULL DEFAULT 1500 COMMENT 'the duration of the contest, unit in seconds'
) ENGINE = InnoDB;
CREATE TABLE IF NOT EXISTS `teams` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `season_id` BIGINT NOT NULL,
  `name` VARCHAR(32) NOT NULL,
  `captain` BIGINT NOT NULL COMMENT 'the id of the captain user',
  -- Since we only three members in the team, keep the table simple as an internal usage system.
  `member1` BIGINT DEFAULT NULL COMMENT 'the id of the first member',
  `member2` BIGINT DEFAULT NULL COMMENT 'the id of the second member'
) ENGINE = InnoDB;