package remind

import "wumingtianqi-sms-pre/model/common"

// todo 提醒模式表
// todo order表（原user_subscribe_content改一下）
// 20200528
// todo 之后整合一下model层下的东西：order, user, city, weather,remind

// 提醒模式表
type RemindPattern struct {
	Id                int    `json:"id" xorm:"pk autoincr INT(11)"`
	RemindText        string `json:"remind_text" xorm:"VARCHAR(20)"`        // 前端配置文案
	RemindObject      string `json:"remind_object" xorm:"VARCHAR(20)"`      // 提醒对象对照表id
	MetClassification string `json:"met_classification" xorm:"VARCHAR(20)"` // 气象分类
	Value             int    `json:"value" xorm:"INT(10)"`                  // 默认配置数值
	FormatText        string `json:"format_text" xorm:"VARCHAR(60)"`        // 提醒用语
	Vip               int    `json:"vip" xorm:"INT(3)"`                     // vip等级
	Tag               string `json:"tag" xorm:"VARCHAR(10)"`                 // 标签：通勤/ 更多配置
	PriorityDisplay   int    `json:"priority_display" xorm:"INT(3)"`        // 前端默认显示优先级
	PriorityRemind    int    `json:"priority_remind" xorm:"INT(3)"`         // 提醒默认优先级
	//Tips           string `json:"tips" xorm:"VARCHAR(40)"`         // 温馨提醒
}

func (m *RemindPattern) QueryOneById(id int) (*RemindPattern, bool, error) {
	has, err := common.Engine.Where("id=?", id).Get(m)
	return m, has, err
}

// todo 造数据
// todo 提醒对象表造！！
//  & 天气对象:
// 天气对象weather表，应该配套有子表：把每一个提醒对象都贴上标签
/* eg:
WeatherObject name isValue unit options
high 最高气温 true 度 []
low 最低气温   true 度 []
quality 空气质量 False "" ["重度污染"，"中度污染", "轻度污染", "优", "良"]
wind_direction 风向 False "" ["东", "南", "西", "北", "东南", "西南", "东北", "西北"]

// Rain 降雨 False "" [1,2,3] // 这个对应天气的代码

phenomenon 天气现象 False "" [1,2,3,4]
ipad 画图，提醒对象几个表，用er图勾勒出来：带着假数据，关键！）
 */



// 提醒模式表 RemindPattern 穷举100种以内
/* eg
Id RemindObjectId Variety FormatText Tips Priority
1 high 增加 {RemindObjectName}{Variety}{value}{unit} 注意防范 1
2 quality 变差 {RemindObjectName}{Variety}  注意防范 2
// 3 phenomenon 突变
4 phenomenon 突然下雨
4 phenomenon 突然下雪
// 5 isSuddenRain 突变 明日有{weather_text} 记得带伞 1

应该用RemindObjectId穷举switch case




 */

