

-- RemindPattern 假数据
select * from `wumingtianqi`.`remind_pattern`;
INSERT INTO `wumingtianqi`.`remind_pattern` (`remind_text`, `remind_object`, `met_classification`, `value`, `format_text`, `vip`, `tag`, `priority_display`, `priority_remind`) VALUES ('突然降雨', '天气现象', '天气现象', -999, '{text_day}{text_night}，注意带伞', 0, '通勤', 1, 1);

INSERT INTO `wumingtianqi`.`remind_pattern` (`remind_text`, `remind_object`, `met_classification`, `value`, `format_text`, `vip`, `tag`, `priority_display`, `priority_remind`) VALUES ('突然升温', '最高温度', '温度', 5, '{text_day}{text_night}，注意带伞', 0, '通勤', 2, 1);
INSERT INTO `wumingtianqi`.`remind_pattern` (`remind_text`, `remind_object`, `met_classification`, `value`, `format_text`, `vip`, `tag`, `priority_display`, `priority_remind`) VALUES ('突然降温', '最高温度', '温度', 5, '{text_day}{text_night}，注意带伞', 0, '通勤', 3, 1);
INSERT INTO `wumingtianqi`.`remind_pattern` (`remind_text`, `remind_object`, `met_classification`, `value`, `format_text`, `vip`, `tag`, `priority_display`, `priority_remind`) VALUES ('空气质量变差', '天气现象', '天气现象', -999, '{text_day}{text_night}，注意带伞', 1, '通勤', 4, 1);
INSERT INTO `wumingtianqi`.`remind_pattern` (`remind_text`, `remind_object`, `met_classification`, `value`, `format_text`, `vip`, `tag`, `priority_display`, `priority_remind`) VALUES ('9点突然升温', '天气现象', '天气现象', -999, '{text_day}{text_night}，注意带伞', 2, '通勤', 5, 1);
INSERT INTO `wumingtianqi`.`remind_pattern` (`remind_text`, `remind_object`, `met_classification`, `value`, `format_text`, `vip`, `tag`, `priority_display`, `priority_remind`) VALUES ('高温预警', '最高温度', '温度', 40, '{text_day}{text_night}，注意带伞', 1, '更多配置', 1, 1);
INSERT INTO `wumingtianqi`.`remind_pattern` (`remind_text`, `remind_object`, `met_classification`, `value`, `format_text`, `vip`, `tag`, `priority_display`, `priority_remind`) VALUES ('低温预警', '最低温度', '温度', -10, '{text_day}{text_night}，注意带伞', 1, '更多配置', 2, 1);

-- Order 假数据
INSERT INTO `wumingtianqi`.`order` (`user_id`, `remind_city`, `remind_time`) VALUES ('1', 'Beijing', '0900');

-- OrderDetail假数据
INSERT INTO `wumingtianqi`.`order_detail` (`order_id`, `remind_pattern_id`, `value`) VALUES ('1', '2', '5');
INSERT INTO `wumingtianqi`.`order_detail` (`order_id`, `remind_pattern_id`, `value`) VALUES ('1', '6', '40');
