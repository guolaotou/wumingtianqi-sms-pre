package main

import (
	"fmt"
	"wumingtianqi-sms-pre/config"
	"wumingtianqi-sms-pre/model"
	"wumingtianqi-sms-pre/model/city"
)

// go run main.go
func main() {
	cfg, _ := config.LoadConfig()
	fmt.Println(cfg.Log)
	model.InitMysql()
	model.InitPubSub()
	cityModel, err := city.GetAllCity()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("cityModel", cityModel)

	// 业务开始

	select {}
}
