USE uestc_acm_training;
INSERT INTO `configurations` (`key`, `value`)
VALUES('current_season_id', 1);
INSERT INTO `seasons` (`id`, `name`)
VALUES(1, 'default season');
INSERT INTO `users` (
    `id`,
    `season_id`,
    `name`,
    `email`,
    `phone`,
    `password`,
    `rating`
  )
VALUES(
    1,
    1,
    'administrator',
    'admin@acm.uestc.edu.cn',
    '13122222222',
    'f3d6e24c7d6976f1a6267e2c9a294caa545e7b381a55d4600bee9fdca5a432c7af3404f0cacaccfb03f513d0f7ad913155b174f26698499a5b2381bc2f925ad9',
    1500
  );