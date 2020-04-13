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
	RemindCity        CityJson        `json:"remind_city" xorm:"json"`      // todo 初始化表时，因为xorm的限制，该字段会被初始化为text字段；需要进入mysql，手动将该字段更改为json格式：alter table wumingtianqi.user_subscribe_content change remind_city remind_city  json;
	RemindCondition   RemindCondition `json:"remind_condition" xorm:"json"` // todo 同上：alter table wumingtianqi.user_subscribe_content change remind_condition remind_condition  json;
	RemindTemplate    int             `json:"remind_template" xorm:"INT(11)"`
	RemindSequence    []int           `json:"remind_sequence" xorm:"list"`
	RemindWay         []int           `json:"remind_way" xorm:"VARCHAR(11)"`
	RemindContactInfo []int           `json:"remind_contact_info" xorm:"VARCHAR(12)"`
}

type City struct {
	Id       int    `json:"id" xorm:"pk autoincr INT(11)"`
	Province string `json:"province" xorm:"VARCHAR(20)"`
	City     string `json:"city" xorm:"VARCHAR(20)"`
	District string `json:"district" xorm:"VARCHAR(20)"`
	PinYin   string `json:"pin_yin" xorm:"VARCHAR(30)"`
	Abbr     string `json:"abbr" xorm:"VARCHAR(60)"`
}

// todo 天气表
// typora里新增天气表的详设

var Engine *xorm.Engine

func MysqlInit() {
	// todo 本地数据库init
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:00000000@(127.0.0.1:3306)/wumingtianqi")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to start  mysql: ", err.Error())
		os.Exit(1)
	}
	Engine.SetColumnMapper(core.SnakeMapper{})
	Engine.SetMaxIdleConns(1000)
	Engine.SetMaxOpenConns(1000)
	Engine.SetConnMaxLifetime(20 * time.Second)
	if syncErr := Engine.Sync2(new(UserSubscribeContent)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync UserSubscribeContent mysql: ", syncErr.Error())
		os.Exit(1)
	}
	if syncErr := Engine.Sync2(new(City)); syncErr != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync City mysql: ", syncErr.Error())
		os.Exit(1)
	}
}
