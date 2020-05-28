package order

// 用户订阅表
type Order struct {
	OrderId         int    `json:"order_id" xorm:"pk autoincr INT(11)"`
	UserId          int    `json:"user_id" xorm:"INT(11)"`
	RemindCity      string `json:"remind_city" xorm:"json"`          // 城市的拼音
	RemindTime      string `json:"remind_time" xorm:"VARCHAR(4)"`    // 提醒时间， 0900
	RemindPatternId int    `json:"remind_pattern_id" xorm:"INT(11)"` // 提醒模式表id
	Value           int    `json:"value" xorm:"INT(11)"`             // 数值
}