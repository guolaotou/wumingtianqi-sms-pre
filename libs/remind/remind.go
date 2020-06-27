package remind

import (
	"context"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
	"wumingtianqi-sms-pre/model/common"
)

// 读取需要提醒的订单队列，拼接短信
func PubSubOrder() {
	messages, err := common.PubSub.Subscribe(context.Background(), "Topic.needToRemindOrder")
	if err != nil {
		panic(err)
	}
	go SpliceSms(messages)
}

func SpliceSms(messages <- chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		// todo 拼接短信，然后放到短信队列里
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world"))  // 封装用户信息和带拼接的短信
		if err := common.PubSub.Publish("Topic.Sms2Send", msg); err != nil {
			panic(err)
		}
		msg.Ack()
	}
}
