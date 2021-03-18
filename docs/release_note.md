### v0.0.1.1[2021-03-04]
**`Query Parameter`:**
**`功能修改`:**
1.发送提醒模块，增加用户的次数限制

2.invitation规则改成
  （1）查invitation_user_map是否用过该邀请码，若用过，直接报错
  （2）用户当前vip等级为邀请码的等级，直接续期；否则，直接覆盖

3.登录时/定时任务发送提醒时 校验vip过期时间

4.提醒的时候校验剩余次数是否足够，保证并发情况下的数据准确性

5.新增"城市列表"接口

6.新增docs/release_note.md; 新增docs/schema.sql

**`确保自测`:**
1. 天气提醒代码可用
2. 新invitation规则确保接口可用
3. 老登录web接口可用
4. go run main.go 可用 （如果是上线阶段，这个应该在改表之后运行）

**`上线流程`：**
1. 改表
```sql
ALTER TABLE `user_info_flexible` CHANGE today_edit_chance_remaining today_tel_remind_remaining int(3) DEFAULT '0' COMMENT '短信提醒当天剩余次数';
ALTER TABLE `user_info_flexible` ADD COLUMN last_remind_time int(11) DEFAULT '20000101' COMMENT '上次提醒时间' AFTER today_tel_remind_remaining;
ALTER TABLE vip_rights_map DROP today_edit_chance_max;
```

2. 改数据
`user_info_flexible`表中所有数据的`today_tel_remind_remaining`字段改成相应vip等级的值

3.部署运行
部署订单模块；确保supervisor重启
部署web；确保supervisor重启

### v0.0.1.2[2021-03-12]
**`Query Parameter`:**
**`功能修改`:**
1.城市列表接口固定返回的顺序
2.代码优化


### v0.0.1.3[2021-03-18]
**`Query Parameter`:**
**`功能修改`:**
1.城市表新增code列，用作城市的唯一标识
2.抓取天气的代码，用code做唯一标识（改day_weather表）
3.前后端查询城市天气列表，用code做城市唯一标识
4.抓取提醒逻辑，用code做唯一标识


**`上线流程`：**
1. 改表
```sql
ALTER TABLE `city` ADD INDEX IDX_city_code ( `code` );
ALTER TABLE `day_weather` CHANGE city_pin_yin city_code varchar(30) NOT NULL;
DROP INDEX `IDX_city_code` ON city;
```

2.更新城市数据
清空city表，重跑脚本load_city_csv2mysql.py
```sql
TRUNCATE `wumingtianqi`.`city`;
```

3.改数据
清空 day_weather表
order表字段为remind_city的值改成对应的city_code
```sql
TRUNCATE `wumingtianqi`.`day_weather`;
SELECT * FROM wumingtianqi.day_weather limit 9999;
SELECT * FROM wumingtianqi.order limit 9999;
```

4.部署运行
部署订单模块；确保supervisor重启
部署web；确保supervisor重启

