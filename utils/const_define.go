package utils

import "time"

const (
	BufferSms                     = 5000 // 发送短信channel的缓冲器
	TickerTime                    = 2 * time.Second
	SmsOneGoroutine               = 10 * time.Second
	PrepareSmsContentGoroutineNum = 10
)
