package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"time"
	"wumingtianqi-sms-pre/model/city"
	"wumingtianqi-sms-pre/model/common"
	"wumingtianqi-sms-pre/model/order"
	"xorm.io/core"
)
//
//var Engine *xorm.Engine


func InitMysql() {
	// todo 本地数据库init
	var err error
	common.Engine, err = xorm.NewEngine("mysql", "root:00000000@(127.0.0.1:3306)/wumingtianqi")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to start  mysql: ", err.Error())
		os.Exit(1)
	}
	common.Engine.SetColumnMapper(core.SnakeMapper{})
	common.Engine.SetMaxIdleConns(1000)
	common.Engine.SetMaxOpenConns(1000)
	common.Engine.SetConnMaxLifetime(20 * time.Second)
	if syncErr := common.Engine.Sync2(new(order.UserSubscribeContent)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync UserSubscribeContent mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := common.Engine.Sync2(new(city.City)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync City mysql: ", syncErr.Error())
		os.Exit(1)
	}
}
