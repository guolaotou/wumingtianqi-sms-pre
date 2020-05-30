package remind

// todo cron要用的"拼接提醒信息" 的model在这里写：queue
// 	参考model/order/sms_pre.go queue1

// 1. 获取5分钟后提醒提醒订单列表 - 存储容器
// model: user, order, Q_to_remind
var QToRemind = map[string][]int{}  // {"0900": [1,2,3]}  代表9点有3个订单1、2、3要提醒
// todo 最大的main改？ 思考是不是挪到这里
// 还可以参考正规cron写法，vanguard代码


// 2. 拼接提醒信息
// model: user, order, RemindPattern， Q_to_remind, Q_to_send
type RemindToSend struct {
	UserId     int    `json:"user_id"`
	RemindText string `json:"remind_text"`
}
var QToSend = map[string][]RemindToSend{}

// lib 短信模块

