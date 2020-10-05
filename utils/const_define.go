package utils

import "time"

const (
	BufferSms                     = 5000 // 发送短信channel的缓冲器
	TickerTime                    = 2 * time.Second
	SmsOneGoroutine               = 10 * time.Second
	PrepareSmsContentGoroutineNum = 10
)

const (
	VIP0 = 0
	VIP1 = 1
	VIP2 = 2
	VIP3 = 3
	VIPMaster = -1
)

const (
	InvitationLevel1 = 1  // 一等邀请码（mine生成10个）VIP3，有效期99年，初始10000个钻石
	InvitationLevel2 = 2  // 二等邀请码（mine生成5个，每个可以给20个人用） 默认VIP2，有效期99年，默认100钻石
	InvitationLevel3 = 3
)