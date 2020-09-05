package sms_pre

import (
	"fmt"
	"runtime"
	"time"
	. "wumingtianqi/model/order"
)

// 消费者&生产者: 循环取Q2(queue2)，制作短信内容，然后放到Q3； 这个是一个一直执行的线程
func PrepareSmsContent(c chan Model, cTmp <-chan string) {
	fmt.Println("\n消费begin")
	for true {
		fmt.Println("\n消费begin2")
		v := <-c

		v2 := <-cTmp
		fmt.Println("消费啦", v2)
		time.Sleep(3 * time.Second)
		fmt.Println(v)
		fmt.Println("goroutine数量:", runtime.NumGoroutine())

		// todo 放到Q3,这个需要缓冲区10000，测试并发？情况下，服务能均匀的每秒消费30个，这个弄完就可以重构啦；
	}
	return
}
