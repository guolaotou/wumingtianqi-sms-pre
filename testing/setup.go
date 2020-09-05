package testing

import (
	"fmt"
	"wumingtianqi/config"
	"wumingtianqi/model"
)

func Setup() {
	if _, err := config.LoadConfig("../../"); err != nil {
		fmt.Println(err.Error())
	}
	model.InitMysql()
}