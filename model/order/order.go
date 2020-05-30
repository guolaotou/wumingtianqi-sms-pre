package order

// 用户订阅表
type Order struct {
	OrderId         int    `json:"order_id" xorm:"pk autoincr INT(11)"`
	UserId          int    `json:"user_id" xorm:"INT(11)"`
	RemindCity      string `json:"remind_city" xorm:"json"`          // 城市的拼音
	RemindTime      string `json:"remind_time" xorm:"VARCHAR(4)"`    // 提醒时间， 0900
	//RemindPatternId int    `json:"remind_pattern_id" xorm:"INT(11)"` // 提醒模式表id
	//Value           int    `json:"value" xorm:"INT(11)"`             // 数值
}

type OrderDetail struct {
	Id              int `json:"id" xorm:"pk autoincr INT(11)"`
	OrderId         int `json:"order_id" xorm:"INT(11)"`
	RemindPatternId int `json:"remind_pattern_id" xorm:"INT(11)"`
	Value           int `json:"value" xorm:"INT(11)"` // 数值
}

/*
remindPattern:
[
{
"remind_pattern": 1,
"value": 1
}
]
 */

/*
Order
OrderId UserId RemindCity RemindTime RemindPatternId Value
1 149 Beijing 0900 1 20
 */
// todo 1 order表改，可以同时多个提醒；
// todo lib写逻辑：
/*
1. 用户获取所有提醒模式
2. 用户配置提醒模式数值，存到order表里
3. lib: 7种提醒模式分别做处理
3.1 todo 测试用例：测以上7种（晚上可以做？）
todo Then 整理代码结构、er图、数据库表设计
 */