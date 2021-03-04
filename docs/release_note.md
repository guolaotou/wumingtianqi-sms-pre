### v0.0.1.1[2021-03-04]
**`Query Parameter`:**
**`功能修改`:**
1.发送提醒模块，增加用户的次数限制

2.invitation规则改成
  （1）查invitation_user_map是否用过该邀请码，若用过，直接报错
  （2）用户当前vip等级为邀请码的等级，直接续期；否则，直接覆盖

3.登录时/定时任务发送提醒时 校验vip过期时间

4.提醒的时候校验剩余次数是否足够，保证并发情况下的数据准确性

5.新增docs/release_note.md; 新增docs/schema.sql

**`确保自测`:**
1. 天气提醒代码可用
2. 新invitation规则确保接口可用
3. 老登录web接口可用
4. go run main.go 可用 （如果是上线阶段，这个应该在改表之后运行）

**`上线流程`：**
1. 改表
```sql
ALTER TABLE `user_info_flexible ` CHANGE today_edit_chance_remaining today_tel_remind_remaining int(3) DEFAULT '0' COMMENT '短信提醒当天剩余次数';
ALTER TABLE `user_info_flexible` ADD COLUMN last_remind_time int(11) DEFAULT '20000101' COMMENT '上次提醒时间' AFTER today_tel_remind_remaining;
ALTER TABLE vip_rights_map DROP today_edit_chance_max;
```

2. 改数据
`user_info_flexible`表中所有数据的`today_tel_remind_remaining`字段改成相应vip等级的值

4.部署运行
部署订单模块；确保supervisor重启
部署web；确保supervisor重启
