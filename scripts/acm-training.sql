DROP DATABASE IF EXISTS uestc_acm_training;
CREATE SCHEMA IF NOT EXISTS uestc_acm_training DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
USE uestc_acm_training;
-- user table
DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(24) NOT NULL UNIQUE,
  `email` VARCHAR(100) NOT NULL UNIQUE,
  `phone` VARCHAR(45) NOT NULL UNIQUE,
  -- sha256sum
  `password` VARCHAR(130) NOT NULL UNIQUE,
  `rating` INT NOT NULL DEFAULT 1500
) ENGINE = InnoDB;
INSERT INTO `users` (
    `id`,
    `name`,
    `email`,
    `phone`,
    `password`,
    `rating`
  )
VALUES(
    1,
    'administrator',
    'admin@acm.uestc.edu.cn',
    '13122222222',
    'f3d6e24c7d6976f1a6267e2c9a294caa545e7b381a55d4600bee9fdca5a432c7af3404f0cacaccfb03f513d0f7ad913155b174f26698499a5b2381bc2f925ad9',
    1500
  );
CREATE TABLE IF NOT EXISTS `contests` (
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(128) NOT NULL,
    `description` VARCHAR(128) DEFAULT NULL,
    `url` VARCHAR(1024) DEFAULT NULL COMMENT 'the contest link, if it is online',
    `start_time` BIGINT NOT NULL COMMENT 'the start time of the contest, as unix timestamp',
    `duration` BIGINT NOT NULL DEFAULT 1500 COMMENT 'the duration of the contest, unit in seconds'
  ) ENGINE = InnoDB;
CREATE TABLE IF NOT EXISTS `teams` (
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(32) NOT NULL,
    `captain` BIGINT NOT NULL COMMENT 'the id of the captain user',
    -- Since we only three members in the team, keep the table simple as an internal usage system.
    `member1` BIGINT DEFAULT NULL COMMENT 'the id of the first member',
    `member2` BIGINT DEFAULT NULL COMMENT 'the id of the second member'
  ) ENGINE = InnoDB;