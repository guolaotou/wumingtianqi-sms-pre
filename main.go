package main

import (
	"fmt"
	"log"
	"wumingtianqi/config"
	"wumingtianqi/model"
	"wumingtianqi/model/city"
	"wumingtianqi/web"

	//"github.com/lithammer/shortuuid/v3"
)

// go run main.go
func main() {
	log.SetFlags(log.Ldate|log.Llongfile)
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
