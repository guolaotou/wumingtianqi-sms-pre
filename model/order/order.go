package order

import (
	"time"
	"wumingtianqi/model/common"
)

// 用户订阅表
// todo xorm 丰富default和comment
type Order struct {
	OrderId 	   int       `json:"order_id" xorm:"pk autoincr INT(11)"`
	//UserId     int       `json:"user_id" xorm:"INT(11)"`
	RemindCity     string    `json:"remind_city" xorm:"json"`            // 城市的拼音
	RemindTime     string    `json:"remind_time" xorm:"VARCHAR(4)"`      // 提醒时间， 0900
	SubscriberName string    `json:"subscriber_name" xorm:"VARCHAR(30)"` // 被提醒者的姓名
	TelephoneNum   string    `json:"telephone_num" xorm:"VARCHAR(33)"`
	Creator        int       `json:"creator" xorm:"INT(11)"`
	CreateTime     time.Time `json:"create_time" xorm:"TIMESTAMP"`
	UpdateTime     time.Time `json:"update_time" xorm:"TIMESTAMP"`
}

func (m *Order) Create() error {
	if _, err := common.Engine.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *Order) Update() error {
	if _, err := common.Engine.Where("order_id=?", m.OrderId).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *Order) Delete() error {
	if _, err := common.Engine.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Order) QueryOneByOrderId(orderId int) (*Order, bool, error) {
	has, err := common.Engine.Where("order_id=?", orderId).Get(m)
	return m, has, err
}

func (m *Order) QueryListByCreator(creator int) ([]Order, error) {
	modelList := make([]Order, 0)
	err := common.Engine.Where("creator=?", creator).Find(&modelList)
	return modelList, err
}

func (m *Order) QueryListByCity(city string) ([]Order, error) {
	modelList := make([]Order, 0)
	err := common.Engine.Where("city=?", city).Find(&modelList)
	return modelList, err
}

func (m * Order) QueryListByTime(time string) ([]Order, error) {
	modelList := make([]Order, 0)
	err :=common.Engine.Where("remind_time=?", time).Find(&modelList)
	return modelList, err
}

type OrderDetail struct {
	Id              int `json:"id" xorm:"pk autoincr INT(11)"`
	OrderId         int `json:"order_id" xorm:"INT(11)"`
	RemindPatternId int `json:"remind_pattern_id" xorm:"INT(11)"`
	Value           int `json:"value" xorm:"INT(11)"` // 数值
	CreateTime time.Time `json:"create_time" xorm:"TIMESTAMP"`
	UpdateTime time.Time `json:"update_time" xorm:"TIMESTAMP"`
}

func (m *OrderDetail) QueryListByOrderId(orderId int) ([]OrderDetail, error) {
	modelList := make([]OrderDetail, 0)
	err := common.Engine.Where("order_id=?", orderId).Find(&modelList)
	return modelList, err
}

// 新增订单时，params: order_detail的最小json结构
type OrderDetailItem struct {
	RemindPatternId int `json:"remind_pattern_id"`
	Value           int `json:"value"`
}

// todo 以上写测试用例
type ResOrderDetailItem struct {
	RemindPatternId int `json:"remind_pattern_id"`
	Value           int `json:"value"`
}
type ResOrderAndDetail struct {
	OrderId     int                  `json:"order_id"`
	PreTele     string               `json:"pre_tele"`
	Telephone   string               `json:"telephone"`
	City        string               `json:"city"`
	CityName    string               `json:"city_name"`
	RemindTime  string               `json:"remind_time"`
	OrderDetail []ResOrderDetailItem `json:"order_detail"`
}

