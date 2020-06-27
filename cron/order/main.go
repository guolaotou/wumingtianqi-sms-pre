package main

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
	"time"
	"wumingtianqi-sms-pre/config"
	"wumingtianqi-sms-pre/libs/order"
	"wumingtianqi-sms-pre/libs/remind"
	"wumingtianqi-sms-pre/model"
)

// 这个模块先用来测试pubsub
// 参考：https://studygolang.com/articles/26894
// github: https://github.com/ThreeDotsLabs/watermill/blob/master/_examples/pubsubs/go-channel/main.go
// go run cron/order/main.go

func main() {
	// 1. 调用lib函数: 每1分钟查看一次订单，将符合条件的放到队列里
	// 2. 调用lib函数，读取队列，如果有订单就按照订单拼接短信，放到短信队列里
	// 3. 调用lib函数，读取短信队列，如果有需要发短信的，就调用发短信的接口
	if _, err := config.LoadConfig(); err != nil {
		fmt.Println(err.Error())
	}
	model.InitMysql()
	model.InitPubSub()  // 如果主main也调用过了这个函数，重复调用会不会重复建立。
	go order.CronOrder()  // 每1分钟查看一次订单，将符合条件的放到队列里
	remind.PubSubOrder()  // 读取需要提醒的订单队列，拼接短信 todo 思考需要需要开启goroutine
	// 建立sms包，发送短信（按照腾讯云的样式）；然后接收Topic.Sms2Send，发送信息；
	// 上面这步需要把splicePattern1里的假天气数据找地方写成真的
	fmt.Println("here")
	select {}

	//pubSub := common.PubSub
	//
	//messages, err := pubSub.Subscribe(context.Background(), "example.topic")
	//if err != nil {
	//	panic(err)
	//}
	//go process(messages)
	//
	//publishMessages(pubSub)
	//select{}
}

func publishMessages(publisher message.Publisher) {
	for {
		newUUID := watermill.NewUUID()
		fmt.Println("watermill.NewUUID()", newUUID)
		msg := message.NewMessage(newUUID, []byte("Hello, world"))
		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(6 * time.Second)
	}
}

func process(messages <- chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		msg.Ack()
	}
}