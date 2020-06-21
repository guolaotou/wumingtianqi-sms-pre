package main

import (
	"fmt"
	"wumingtianqi-sms-pre/config"
	"wumingtianqi-sms-pre/libs/weather"
	"wumingtianqi-sms-pre/model"
)

// go run cron/weather/main.go
func main() {
	fmt.Println("duandian1")
	if _, err := config.LoadConfig(); err != nil {
		fmt.Println(err.Error())
	}
	model.InitMysql()
	// 不断更新天气信息
	for true {
		weather.UpdateWeatherDaily()
	}
}
