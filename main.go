package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Model struct {
	UserId int    `json:"user_id"`
	Value  string `json:"value"`
}

const (
	BufferSms  = 5000 // 发送短信channel的缓冲器
	tickerTime = 2 * time.Second
)


// todo 把这个放到util函数里，不要干扰主流程；
// todo 几个长运行的函数放到某文件里，然后用main调用
// 内部函数：获取指定时间的时分，默认北京时间(4位数字的str格式)
func getLocalHourMin4Str() string {
	durationNum, _ := time.ParseDuration(strconv.Itoa(28800) + "s") // 时区偏移量（北京时间）
	localDate := time.Now().UTC().Add(durationNum)
	localDateStr := localDate.Format("1504")
	return localDateStr
}

func getSpecifyDurationHourMin(duration time.Duration) string {
	durationNum, _ := time.ParseDuration(strconv.Itoa(28800) + "s") // 时区偏移量
	localDate := time.Now().UTC().Add(durationNum).Add(duration)    // 北京时间加上指定偏移时间
	localDateStr := localDate.Format("1504")
	return localDateStr
}

// Q1，用来存订单
var queue1 = map[string][]Model{}

// Q2，用来存要发短信的订单
var queue2 = map[string][]Model{}

// Q3，所有要发的短信内容放进这个队列
var queue3 = map[string][]Model{}

// 造数据，存入Q1
func makeData() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1440; i++ {
		keyStr := getSpecifyDurationHourMin(time.Minute * time.Duration(i)) // 时分作为map的key

		//randValue := rand.Intn(3)   // [0,3)的随机值，返回值为int，为了造数据用，代表某时间有0， 1， 2个订单
		randValue := 10000
		for j := 0; j < randValue; j++ {
			model := Model{ // 实例化一个订单，并填入造的数据
				UserId: j,
				Value:  strconv.Itoa(j),
			}
			queue1[keyStr] = append(queue1[keyStr], model)
		}
		fmt.Println("i: ", i)
	}
}

// todo 拆解上面子模块：每天的订单在10W的量，需要能够手动加快时间的流动，模拟订单在一天内的消耗过程
// todo 我的思维导图需要有个地方存放；

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

// 循环取Q2(queue2)，制作短信内容，然后放到Q3； 这个是一个一直执行的线程
func prepareSmsContent(c chan Model, cTmp <-chan string) {
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

// 循环取Q3，发送短信，同时控制每秒能够发送的数量
func func3() {

}

// 发送短信的func
func SmsFunc() {

}

var wg sync.WaitGroup

// todo 命名

func main() {
	makeData() // 模拟每天凌晨跑出来今日的订单

	c := make(chan Model, BufferSms)
	cTmp := make(chan string, BufferSms)

	go prepareSmsContent(c, cTmp) // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	//go prepareSmsContent(c, cTmp)  // 消费者
	wg.Add(1)

	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可
	//以释放相关资源。
	ticker1 := time.NewTicker(tickerTime)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
			if queue1[getLocalHourMin4Str()] != nil {
				fmt.Println(getLocalHourMin4Str(), "数据来啦")
				fmt.Printf("%p", queue1[getLocalHourMin4Str()])

				// 预分配足够多的元素切片
				lenData := len(queue1[getLocalHourMin4Str()])
				copyData := make([]Model, lenData)
				// 将数据复制到新的切片空间中
				copy(copyData, queue1[getLocalHourMin4Str()])
				go FilterOrders(copyData, c, cTmp, getLocalHourMin4Str()) // 生产者
				time.Sleep(tickerTime)
			}
		}
	}(ticker1)

	wg.Wait()
}
