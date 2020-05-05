package order

import (
	_ "github.com/go-sql-driver/mysql"
	"wumingtianqi-sms-pre/model/common"
)

//type CityJson struct {
//	Province string `json:"province"`
//	City     string `json:"city"`
//	Region   string `json:"region"`
//}

type RemindConditionJson struct {
}

type UserSubscribeContent struct {
	OrderId     int `json:"order_id" xorm:"pk autoincr INT(11)"`
	UserId      int `json:"user_id" xorm:"INT(11)"`
	SubscribeId int `json:"subscribe_id" xorm:"INT(11)"`
	RemindTime  int `json:"remind_time" xorm:"INT(4)"`
	//RemindCity        CityJson        `json:"remind_city" xorm:"json"`      // todo 初始化表时，因为xorm的限制，该字段会被初始化为text字段；需要进入mysql，手动将该字段更改为json格式：alter table wumingtianqi.user_subscribe_content change remind_city remind_city  json;
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

func init() {
	println("1111")
	//if syncErr := model.Engine.Sync2(new(UserSubscribeContent)); syncErr != nil {
	//	_, _ = fmt.Fprintln(os.Stderr, "Failed to sync UserSubscribeContent mysql: ", syncErr.Error())
	//	os.Exit(1)
	//}
}

//func init(){
//	if syncErr := model.Engine.Sync2(new(city.City)); syncErr != nil {
//		_, _ = fmt.Fprintln(os.Stderr, "Failed to sync City mysql: ", syncErr.Error())
//		os.Exit(1)
//	}
//}
