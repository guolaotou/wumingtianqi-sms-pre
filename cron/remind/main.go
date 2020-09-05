package main

import (
	"fmt"
	"time"
	"github.com/robfig/cron"
)

// 1. 获取5分钟后提醒提醒用户列表
// model: user, order, Q_to_remind


// 2. 拼接提醒信息
// model: user, order, Q_to_remind, Q_to_send
// lib 短信模块


// 调用remind相关的后台（定时）任务
// go run cron/remind/main.go
func f()  {
	fmt.Println(time.Now())
}

func dingshirenwu1() {
	// 定时任务更新天气信息，调用跑起来？更新北京的天气
	// 发布订阅
	//fmt.Println("duandian1")
	//c := time.Tick(10 * time.Second)
	//for {
	//	<- c
	//	go f()
	//}
}

//// 返回一个支持至 秒 级别的 cron
//func newWithSeconds() *cron.Cron {
//	secondParser := cron.NewParser(cron.Second | cron.Minute |
//		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
//	return cron.New(cron.WithParser(secondParser), cron.WithChain())
//}

func dingshirenwu2() {
	fmt.Println("duandian2")
	//c := newWithSeconds()
	c := cron.New()
	err := c.AddFunc("*/2 * * * * *", f)
	if err != nil {
		fmt.Println(err.Error())
	}
	go c.Start()
	defer c.Stop()

	select {}
}

// 运行方法1：
//  mac build: /usr/local/go/bin/go build -o ~/go/src/wumingtianqi-remind.out -v wumingtianqi/cron/remind
//  mac run: ~/go/src/wumingtianqi-remind.out
// 运行方法2：
// go run cron/remind/main.go
func main() {  // 定时任务参考 https://www.bookstack.cn/read/topgoer/58c7a1319bdc2491.md
	go dingshirenwu1() // 调用天气抓取，存取天气信息；更新天气信息；第一步完成北京的
	go dingshirenwu2() // 调用 order，then拼接remind信息
	// 发送提醒，先用邮件；再用短信（618买短信？买服务器）
	select {}


}

