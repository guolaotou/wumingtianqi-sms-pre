package order

import (
	_ "github.com/go-sql-driver/mysql"
	"wumingtianqi/model/common"
)


type RemindConditionJson struct {
}

type UserSubscribeContent struct {
	OrderId     int `json:"order_id" xorm:"pk autoincr INT(11)"`
	UserId      int `json:"user_id" xorm:"INT(11)"`
	SubscribeId int `json:"subscribe_id" xorm:"INT(11)"`
	RemindTime  int `json:"remind_time" xorm:"INT(4)"`
	RemindCity        string              `json:"remind_city" xorm:"json"`      // 城市的拼音
	RemindCondition   RemindConditionJson `json:"remind_condition" xorm:"json"` // todo 同上：alter table wumingtianqi.user_subscribe_content change remind_condition remind_condition  json;
	RemindTemplate    int                 `json:"remind_template" xorm:"INT(11)"`
	RemindSequence    []int               `json:"remind_sequence" xorm:"list"`            // todo 改？
	RemindWay         []int               `json:"remind_way" xorm:"VARCHAR(11)"`          // todo 改？
	RemindContactInfo []int               `json:"remind_contact_info" xorm:"VARCHAR(12)"` // todo 改？
}

func GetAll() ([]UserSubscribeContent, error) {
	cityList := make([]UserSubscribeContent, 0)
	err := common.Engine.Find(&cityList)
	return cityList, err
}

// 订阅 todo，然后写代码控制订阅的模块（较多）