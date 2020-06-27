package main

import (
	"fmt"
	"wumingtianqi-sms-pre/config"
	"wumingtianqi-sms-pre/libs/weather"
	"wumingtianqi-sms-pre/model"
)

// 运行方法1：
//  mac build: /usr/local/go/bin/go build -o ~/go/src/wumingtianqi-weather.out -v wumingtianqi-sms-pre/cron/weather
//  mac run: ~/go/src/wumingtianqi-weather.out
// 运行方法2：
// go run cron/weather/main.go
func main() {
	if _, err := config.LoadConfig(); err != nil {
		fmt.Println(err.Error())
	}
	model.InitMysql()
	// 不断更新天气信息
	for true {
		weather.UpdateWeatherDaily()
	}
}
