package common
// 存放pubsub用到的json结构体

// 需要提醒的订单信息
type NeedToRemindOrder struct {
	CityCode       string `json:"city_code"`
	SubscriberName string `json:"subscriber_name"`
	TelephoneNum   string `json:"telephone_num"`
	Creator        int    `json:"creator"`
	Tips           string `json:"tips"`
}

// 拼接好的待发送短信，按照微信短信的格式拼接？
type Sms2Send struct {
	TelephoneNum  string `json:"telephone_num"`
	City          string `json:"city"`
	ToSendContent string `json:"to_send_content"`
}