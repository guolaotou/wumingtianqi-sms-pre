package sms_pre

import (
	"time"
	"wumingtianqi/utils"
)

// 消费者：循环取Q3，发送短信，同时控制每秒能够发送的数量
func SendSmsProducer() {
	// 10秒钟只发一个短信
	time.Sleep(utils.SmsOneGoroutine)
}
