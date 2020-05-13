package testing

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	. "wumingtianqi-sms-pre/model/order"
	"wumingtianqi-sms-pre/utils"
)

// 造数据，存入Q1
func MakeData() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1440; i++ {
		keyStr := utils.GetSpecifyDurationHourMin(time.Minute * time.Duration(i)) // 时分作为map的key

		//randValue := rand.Intn(3)   // [0,3)的随机值，返回值为int，为了造数据用，代表某时间有0， 1， 2个订单
		randValue := 10000
		for j := 0; j < randValue; j++ {
			model := Model{ // 实例化一个订单，并填入造的数据
				UserId: j,
				Value:  strconv.Itoa(j),
			}
			Queue1[keyStr] = append(Queue1[keyStr], model)
		}
		fmt.Println("i: ", i)
	}
}
// todo 拆解上面子模块：每天的订单在10W的量，需要能够手动加快时间的流动，模拟订单在一天内的消耗过程
