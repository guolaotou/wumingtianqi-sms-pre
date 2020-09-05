package user

import (
	"time"
	"wumingtianqi/model/common"
)

// 被提醒用户表
type UserToRemind struct {
	SubscriberId   int    `json:"subscriber_id" xorm:"pk autoincr INT(11)"`
	SubscriberName string `json:"subscriber_name" xorm:"VARCHAR(30)"`
	TelephoneNum   string `json:"telephone_num" xorm:"VARCHAR(33)"`
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

/*
用户信息表

Field	type	Desc	Unique	Index	Null
id	int(11)	用户id;自增主键	yes	yes	no
wx_open_id	varchar(100)	微信open_id	yes	yes	no
wx_union_id	varchar(100)	微信union_id	yes	yes	no
created	datetime	创建时间	no	yes	no
updated	datetime	更新时间	no	yes	no
 */

// 用户信息表（微信信息存储）
type UserInfo struct {
	Id         int       `json:"id" xorm:"pk autoincr INT(11)"`
	WxOpenId   string    `json:"wx_open_id" xorm:"VARCHAR(100) default '' comment('微信open_id') index"`
	WxUnionId  string    `json:"wx_union_id" xorm:"VARCHAR(100) default '' comment('微信open_id') index"`
	CreateTime time.Time `json:"create_time" xorm:"TIMESTAMP"`
	UpdateTime time.Time `json:"update_time" xorm:"TIMESTAMP"`
}

func (m *UserInfo) Create() error {
	if _, err := common.Engine.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInfo) Update() error {
	if _, err := common.Engine.Where("id=?", m.Id).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInfo) Delete() error {
	if _, err := common.Engine.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInfo) QueryById(id int) (*UserInfo, bool, error) {
	has, err := common.Engine.Where("id=?", id).Get(m)
	return m, has, err
}

func (m *UserInfo) QueryByOpenId(openId string) (*UserInfo, bool, error) {
	has, err := common.Engine.Where("wx_open_id=?", openId).Get(m)
	return m, has, err
}