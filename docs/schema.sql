--
-- Table structure for table `user_info_flexible`
--
DROP TABLE IF EXISTS `user_info_flexible`;
CREATE TABLE `user_info_flexible` (
  `user_id` int(11) NOT NULL,
  `invitation_code` varchar(100) DEFAULT '' COMMENT '邀请码',
  `vip_level` int(3) DEFAULT '0',
  `wechat_order_remaining` int(3) DEFAULT '0' COMMENT '微信订单剩余配置数',
  `tel_order_remaining` int(3) DEFAULT '0' COMMENT '手机号订单剩余配置数',
  `today_edit_chance_remaining` int(3) DEFAULT '10' COMMENT '当天剩余编辑次数',
  `coin` int(20) DEFAULT '0',
  `diamond` int(11) DEFAULT '0',
  `expiration_time` int(11) DEFAULT '20000101',
  `creator` int(11) DEFAULT '-1',
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE `user_info_flexible` CHANGE today_edit_chance_remaining today_tel_remind_remaining int(3) DEFAULT '0' COMMENT '短信提醒当天剩余次数';
ALTER TABLE `user_info_flexible` ADD COLUMN last_remind_time int(11) DEFAULT '20000101' COMMENT '上次提醒时间' AFTER today_tel_remind_remaining;


--
-- Table structure for table `vip_rights_map`
--
CREATE TABLE `vip_rights_map` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vip_level` int(3) DEFAULT NULL,
  `wechat_order_max` int(3) DEFAULT '3' COMMENT '微信订单最大配置数（-1代表无限）',
  `tel_order_max` int(3) DEFAULT '0' COMMENT '手机号订单最大配置数（-1代表无限）',
  `today_edit_chance_max` int(3) DEFAULT '10' COMMENT '每天可编辑次数',
  `remind_pattern_id_list` text COMMENT '提醒模式id列表',
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

ALTER TABLE vip_rights_map DROP today_edit_chance_max;

--
-- Table structure for table `user_invitation_map`
--
CREATE TABLE `user_invitation_map` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `invitation_code` varchar(100) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_user_invitation_map_user_invitation` (`user_id`,`invitation_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


--
-- Table structure for table `city`
--
CREATE TABLE `city` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `province` varchar(20) DEFAULT NULL,
  `city` varchar(20) DEFAULT NULL,
  `district` varchar(20) DEFAULT NULL,
  `pin_yin` varchar(30) DEFAULT NULL,
  `abbr` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_city_pin_yin` (`pin_yin`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
ALTER TABLE `city` ADD COLUMN code varchar(32) DEFAULT 'A' COMMENT 'xinzhi code' AFTER abbr;
ALTER TABLE `city` ADD INDEX IDX_city_code ( `code` )
DROP INDEX `IDX_city_code` ON city;

--
-- Table structure for table `day_weather`
--
CREATE TABLE `day_weather` (
  `city_pin_yin` varchar(30) NOT NULL,
  `date_id` int(11) NOT NULL,
  `text_day` varchar(30) DEFAULT NULL,
  `code_day` int(11) DEFAULT NULL,
  `text_night` varchar(30) DEFAULT NULL,
  `code_night` int(11) DEFAULT NULL,
  `high` int(11) DEFAULT NULL,
  `low` int(11) DEFAULT NULL,
  `wind_direction` varchar(30) DEFAULT NULL,
  `wind_scale` int(11) DEFAULT NULL,
  `wind_speed` int(11) DEFAULT NULL,
  `humidity` int(11) DEFAULT NULL,
  PRIMARY KEY (`city_pin_yin`,`date_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
ALTER TABLE `day_weather` CHANGE city_pin_yin city_code varchar(30) NOT NULL;
