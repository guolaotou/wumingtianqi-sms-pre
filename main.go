package main

import (
	"fmt"
	"log"
	"wumingtianqi/config"
	"wumingtianqi/model"
	"wumingtianqi/web"

	//"github.com/lithammer/shortuuid/v3"
)

// go run main.go
func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	cfg, _ := config.LoadConfig()
	fmt.Println(cfg.Log)
	model.InitMysql()
	model.InitPubSub()
	log.Println("init finished")

	// 业务开始
	web.ListenHttp()
	log.Println("web started")
	//_ := shortuuid.New() // Cekw67uyMpBGZLRP2HFVbe

	select {}
}
