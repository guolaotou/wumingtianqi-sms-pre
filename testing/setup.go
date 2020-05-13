package testing

import (
	"fmt"
	"wumingtianqi-sms-pre/config"
	"wumingtianqi-sms-pre/model"
)

func Setup() {
	if _, err := config.LoadConfig("../../"); err != nil {
		fmt.Println(err.Error())
	}
	model.InitMysql()
}