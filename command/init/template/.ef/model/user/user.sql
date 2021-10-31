CREATE TABLE IF NOT EXISTS `user` (
  `id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT "雪花id",
  `useranme` VARCHAR(20) NOT NULL DEFAULT "" COMMENT "用户名",
  `gender` TINYINT(1) NOT NULL DEFAULT 0 COMMENT "性别,0-未知,1-男,2-女",
  `marriage` TINYINT(1) NOT NULL DEFAULT 0 COMMENT "是否已婚,0-未知,1-未婚,2-已婚",
  PRIMARY KEY(`id`),
  INDEX `idx_username` (`username`)
) COMMENT = "用户表";