package main

import (
	"fmt"
	"wumingtianqi-sms-pre/config"
	"wumingtianqi-sms-pre/model"
	"wumingtianqi-sms-pre/model/city"
	"wumingtianqi-sms-pre/web"

	//"github.com/lithammer/shortuuid/v3"
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
	web.ListenHttp()
	//_ := shortuuid.New() // Cekw67uyMpBGZLRP2HFVbe

	select {}
}
