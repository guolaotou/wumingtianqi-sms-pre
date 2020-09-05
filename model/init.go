package model

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"time"
	"wumingtianqi/config"
	"wumingtianqi/model/city"
	"wumingtianqi/model/common"
	"wumingtianqi/model/order"
	"wumingtianqi/model/remind"
	"wumingtianqi/model/user"
	"wumingtianqi/model/weather"
	"xorm.io/core"
)
func InitMysql() {
	// todo 本地数据库init
	var err error
	mysqlConfig := config.GlobalConfig.Main.Mysql
	common.Engine, err = xorm.NewEngine("mysql", mysqlConfig)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to start  mysql: ", err.Error())
		os.Exit(1)
	}
	common.Engine.SetColumnMapper(core.SnakeMapper{})
	common.Engine.SetMaxIdleConns(1000)
	common.Engine.SetMaxOpenConns(1000)
	common.Engine.SetConnMaxLifetime(20 * time.Second)
	fmt.Println("duandian2", mysqlConfig)

	if syncErr := common.Engine.Sync2(new(order.UserSubscribeContent)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync UserSubscribeContent mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := common.Engine.Sync2(new(city.City)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync City mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := common.Engine.Sync2(new(weather.DayWeather)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync DayWeather mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := common.Engine.Sync2(new(order.RemindCondition)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync RemindCondition mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := common.Engine.Sync2(new(remind.RemindPattern)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync RemindPattern mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := common.Engine.Sync2(new(order.Order)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync Order mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := common.Engine.Sync2(new(order.OrderDetail)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync OrderDetail mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := common.Engine.Sync2(new(user.UserToRemind)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync UserToRemind mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := common.Engine.Sync2(new(user.UserInfo)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync UserInfo mysql: ", syncErr.Error())
		os.Exit(1)
	}
}

func InitPubSub() {
	common.PubSub = gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
}