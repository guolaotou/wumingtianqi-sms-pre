package order

import (
	"wumingtianqi-sms-pre/model/common"
)

// 用户订阅表
type Order struct {
	OrderId         int    `json:"order_id" xorm:"pk autoincr INT(11)"`
	UserId          int    `json:"user_id" xorm:"INT(11)"`
	RemindCity      string `json:"remind_city" xorm:"json"`          // 城市的拼音
	RemindTime      string `json:"remind_time" xorm:"VARCHAR(4)"`    // 提醒时间， 0900
	//RemindPatternId int    `json:"remind_pattern_id" xorm:"INT(11)"` // 提醒模式表id
	//Value           int    `json:"value" xorm:"INT(11)"`             // 数值
}

type OrderDetail struct {
	Id              int `json:"id" xorm:"pk autoincr INT(11)"`
	OrderId         int `json:"order_id" xorm:"INT(11)"`
	RemindPatternId int `json:"remind_pattern_id" xorm:"INT(11)"`
	Value           int `json:"value" xorm:"INT(11)"` // 数值
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

func (m *Order) QueryListByUserId(userId int) ([]Order, error) {
	modelList := make([]Order, 0)
	err := common.Engine.Find(&modelList)
	return modelList, err
}

func (m *Order) QueryListByCity(city string) ([]Order, error) {
	modelList := make([]Order, 0)
	err := common.Engine.Find(&modelList)
	return modelList, err
}

func (m * Order) QueryListByTime(time string) ([]Order, error) {
	modelList := make([]Order, 0)
	err :=common.Engine.Where("remind_time=?", time).Find(&modelList)
	return modelList, err
}

func (m *OrderDetail) QueryListByOrderId(orderId int) ([]OrderDetail, error) {
	modelList := make([]OrderDetail, 0)
	err := common.Engine.Where("order_id=?", orderId).Find(&modelList)
	return modelList, err
}

// todo 以上写测试用例
