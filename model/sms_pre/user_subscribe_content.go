package sms_pre

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"time"
	"xorm.io/core"
)

type CityJson struct {
	Province string `json:"province"`
	City     string `json:"city"`
	Region   string `json:"region"`
}

type RemindCondition struct {
}

type UserSubscribeContent struct {
	OrderId           int             `json:"order_id" xorm:"pk autoincr INT(11)"`
	UserId            int             `json:"user_id" xorm:"INT(11)"`
	SubscribeId       int             `json:"subscribe_id" xorm:"INT(11)"`
	RemindTime        int             `json:"remind_time" xorm:"INT(4)"`
	RemindCity        CityJson        `json:"remind_city" xorm:"json"` // todo 试一试
	RemindCondition   RemindCondition `json:"remind_condition" xorm:"json"`
	RemindTemplate    int             `json:"remind_template xorm:"INT(11)`
	RemindSequence    []int           `json:"remind_sequence" xorm:"list"`
	RemindWay         []int           `json:"remind_way" xorm:"VARCHAR(11)"`
	RemindContactInfo []int           `json:"remind_contact_info" xorm:"VARCHAR(12)"`
}

var engine *xorm.Engine

func MysqlInit() {
	// todo 本地数据库init
	engine, err := xorm.NewEngine("mysql", "root:00000000@(127.0.0.1:3306)/giftcenter")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to start  mysql: ", err.Error())
		os.Exit(1)
	}
	engine.SetColumnMapper(core.SnakeMapper{})
	engine.SetMaxIdleConns(1000)
	engine.SetMaxOpenConns(1000)
	engine.SetConnMaxLifetime(20 * time.Second)
	if syncErr := engine.Sync2(new(UserSubscribeContent)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync UserSubscribeContent mysql: ", syncErr.Error())
		os.Exit(1)
	}
}
