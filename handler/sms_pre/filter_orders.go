package sms_pre


import (
	"fmt"
	. "wumingtianqi-sms-pre/model/sms_pre"
)

// 生产者：接收所有订单，然后判断这些函数是否需要发送短信，如果需要，则放到chan里
func FilterOrders(orderList []Model, c chan<- Model, cTmp chan string, cTmpValue string) { // 生产者
	fmt.Println("\nlen", len(orderList))
	if len(orderList) == 10000 { // 假设订单数量为1就是需要发短信
		// 加到队列Q2
		// queue2
		for i := 0; i < 10; i++ {
			fmt.Println("生产啦", cTmpValue)
			c <- orderList[i]
			cTmp <- cTmpValue
			fmt.Println("生产结束啦", cTmpValue)
		}
	}
	//os.Exit(1)
}