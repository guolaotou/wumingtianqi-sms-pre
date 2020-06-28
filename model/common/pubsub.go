package common
// 存放pubsub用到的json结构体

// 需要提醒的订单信息
type NeedToRemindOrder struct {
	SubscriberId int    `json:"subscriber_id"`
	City         string `json:"city"`
	Tips         string `json:"tips"`
}

// 拼接好的待发送短信，按照微信短信的格式拼接？
type Sms2Send struct {
	TelephoneNum  string `json:"telephone_num"`
	City          string `json:"city"`
	ToSendContent string `json:"to_send_content"`
}