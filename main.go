package main

import (
	"fmt"
	"sync"
	"time"
	"wumingtianqi-sms-pre/config"
	"wumingtianqi-sms-pre/handler/sms_pre"
	"wumingtianqi-sms-pre/model"
	"wumingtianqi-sms-pre/model/city"
	. "wumingtianqi-sms-pre/model/order"
	"wumingtianqi-sms-pre/testing"
	"wumingtianqi-sms-pre/utils"
)

func main() {
	cfg, _ := config.LoadConfig()
	fmt.Println(cfg.Log)
	model.InitMysql()
	cityModel, err := city.GetAllCity()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("cityModel", cityModel)
	//fmt.Println("cityModel", cityModel[0].PinYin)
	//city, err := city.GetOneCity()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("city", city)

	//res, err := weather.CityWeatherDailyGet(cityModel[0].PinYin)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(res)
	//}
	testing.MakeData() // 模拟每天凌晨跑出来今日的订单
	c := make(chan Model, utils.BufferSms)
	cTmp := make(chan string, utils.BufferSms)

	for i := 0; i < utils.PrepareSmsContentGoroutineNum; i++ { // 准备短信内容
		go sms_pre.PrepareSmsContent(c, cTmp) // 消费者

	}

	var wg sync.WaitGroup
	wg.Add(1)

	//NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 就向该通道发送当时的时间。它会调
	//整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源。
	ticker1 := time.NewTicker(utils.TickerTime)
	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
			if Queue1[utils.GetLocalHourMin4Str()] != nil {
				fmt.Println(utils.GetLocalHourMin4Str(), "数据来啦")
				fmt.Printf("%p", Queue1[utils.GetLocalHourMin4Str()])

				// 预分配足够多的元素切片
				lenData := len(Queue1[utils.GetLocalHourMin4Str()])
				copyData := make([]Model, lenData)
				// 将数据复制到新的切片空间中
				copy(copyData, Queue1[utils.GetLocalHourMin4Str()])
				go sms_pre.FilterOrders(copyData, c, cTmp, utils.GetLocalHourMin4Str()) // 生产者
				time.Sleep(utils.TickerTime)
			}
		}
	}(ticker1)

	wg.Wait()
}
