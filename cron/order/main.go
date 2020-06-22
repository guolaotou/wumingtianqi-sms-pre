package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
	"time"
)

// 这个模块先用来测试pubsub
// go run cron/order/main.go
func main() {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	messages, err := pubSub.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process(messages)
	go process(messages)

	publishMessages(pubSub)
}

func publishMessages(publisher message.Publisher) {
	for {
		aa := watermill.NewUUID()
		fmt.Println("watermill.NewUUID()", aa)
		msg := message.NewMessage(aa, []byte("Hello, world"))
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