package remind

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
	"strconv"
	"wumingtianqi-sms-pre/model/common"
	"wumingtianqi-sms-pre/model/user"
)

// 读取需要提醒的订单队列，拼接短信
func PubSubOrder() {  // 订阅"需要提醒的订单" then发布what？？
	messages, err := common.PubSub.Subscribe(context.Background(), "Topic.needToRemindOrder")
	if err != nil {
		panic(err)
	}
	go func() {
		for msg := range messages {
			log.Printf("received order message: %s, payload: %s", msg.UUID, string(msg.Payload))

			// 1.解析Payload
			needToRemindOrder := new(common.NeedToRemindOrder)
			err := json.Unmarshal(msg.Payload, needToRemindOrder)
			if err != nil {
				log.Println("err needToRemindOrder", err.Error())
			}

			// 2.拼接短信，然后放到短信队列里
			// 2.1根据SubscriberId找到电话
			user, has, err := user.QueryById(needToRemindOrder.SubscriberId)
			if err != nil || !has {
				log.Printf("get user %s error", strconv.Itoa(needToRemindOrder.SubscriberId))
			} else {
				log.Printf("user %s not found", strconv.Itoa(needToRemindOrder.SubscriberId))
			}
			// 2.2 拼接完整短信  // todo libs/order/order.go pattern那控制字符数，不要在后面控制了；测出字符边界在哪
			var toSendContent string  // todo 根据提醒时间，控制字符是明日还是今日
			// todo 找个地方split整段文字为12个字符一组的
			sms2Send := common.Sms2Send{
				TelephoneNum:  user.TelephoneNum,
				City : needToRemindOrder.City,
				ToSendContent: toSendContent,
			}
			sms2SendJson, err := json.Marshal(sms2Send)
			if err != nil {
				log.Println(err.Error())
				panic(err)
			}

			//msg2 := message.NewMessage(watermill.NewUUID(), []byte("Hello, world"))  // 封装用户信息和带拼接的短信
			msg2 := message.NewMessage(watermill.NewUUID(), sms2SendJson)  // 封装用户信息和带拼接的短信
			if err := common.PubSub.Publish("Topic.Sms2Send", msg2); err != nil {
				panic(err)
			}
			msg.Ack()
		}
	}()
	//go SpliceSms(messages)
}

//func SpliceSms(messages <- chan *message.Message) {  // 消费 then生产
//	for msg := range messages {
//		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
//		// todo 拼接短信，然后放到短信队列里
//		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world"))  // 封装用户信息和带拼接的短信
//		if err := common.PubSub.Publish("Topic.Sms2Send", msg); err != nil {
//			panic(err)
//		}
//		msg.Ack()
//	}
//}

// 读取需要发送的短信队列，发送短信
func PubSubSms() {  // todo 以后是不是放到短信模块里
	messages, err := common.PubSub.Subscribe(context.Background(), "Topic.Sms2Send")
	if err != nil {
		panic(err)
	}
	go func() {
		for msg := range messages {
			log.Printf("received sms message: %s, payload: %s", msg.UUID, string(msg.Payload))
			sms2Send := new(common.Sms2Send)
			err := json.Unmarshal(msg.Payload, sms2Send)
			if err != nil {
				fmt.Println("err sms2Send", err.Error())
			}
			fmt.Println("sms2Send", sms2Send)
			// todo 调用发送短信函数，想想控制速率的地方
			// 这里sleep可以么？
			// todo next 把FakeWeather改成真实数据，结构改下，改成weather里包裹城市和时间的
			// todo 腾讯云短信服务搞上
			msg.Ack()
		}
	}()
}
