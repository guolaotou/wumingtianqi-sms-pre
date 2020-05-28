package remind

// todo 提醒模式表
// todo order表（原user_subscribe_content改一下）
// 20200528
// todo 之后整合一下model层下的东西：order, user, city, weather,remind

// 提醒模式表
type RemindPattern struct {
	Id             int    `json:"id" xorm:"pk autoincr INT(11)"`
	RemindObjectId int    `json:"remind_object_id" xorm:"INT(11)"` // 提醒对象对照表id
	Variety        string `json:"variety" xorm:"VARCHAR(30)"`      // 变化动词
	IsValue        int    `json:"is_value" xorm:"INT(2)"`          // 是否配置数值
	Unit           string `json:"unit" xorm:"VARCHAR(30)"`         // 单位
	FormatText     string `json:"format_text" xorm:"VARCHAR(60)"`  // 提醒用语
	Tips           string `json:"tips" xorm:"VARCHAR(40)"`
	Priority       int    `json:"priority" xorm:"INT(11)"`
}