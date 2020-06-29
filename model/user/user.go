package user

import "wumingtianqi-sms-pre/model/common"

// 被提醒用户表
type UserToRemind struct {
	SubscriberId  int    `json:"subscriber_id" xorm:"pk autoincr INT(11)"`
	SubscriberName string `json:"subscriber_name" xorm:"VARCHAR(30)"`
	TelephoneNum  string `json:"telephone_num" xorm:"VARCHAR(33)"`
}

func (m *UserToRemind) Create() error {
	if _, err := common.Engine.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *UserToRemind) Update() error {
	if _, err := common.Engine.Where("order_id=?", m.SubscriberId).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *UserToRemind) Delete() error {
	if _, err := common.Engine.Delete(m); err != nil {
		return err
	}
	return nil
}

func QueryById(subscriberId int)(*UserToRemind, bool, error) {
	utr := new(UserToRemind)
	has, err := common.Engine.Where("subscriber_id=?", subscriberId).Get(utr)
	return utr, has, err
}