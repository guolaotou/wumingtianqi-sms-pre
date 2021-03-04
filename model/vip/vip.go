package vip

import (
	"time"
	"wumingtianqi/model/common"
)

// vip权益映射表
type VipRightsMap struct {
	Id                  int       `json:"id" xorm:"pk autoincr INT(11)"`
	VipLevel            int       `json:"vip_level" xorm:"INT(3)"`
	WechatOrderMax      int       `json:"wechat_order_max" xorm:"INT(3) default 3 comment('微信订单最大配置数（-1代表无限）')"`
	TelOrderMax         int       `json:"tel_order_max" xorm:"INT(3) default 0 comment('手机号订单最大配置数（-1代表无限）')"`
	RemindPatternIdList string    `json:"remind_pattern_id_list" xorm:"TEXT comment('提醒模式id列表')"`
	CreateTime          time.Time `json:"create_time" xorm:"TIMESTAMP"`
	UpdateTime          time.Time `json:"update_time" xorm:"TIMESTAMP"`
}

func (m *VipRightsMap) Create() error {
	if _, err := common.Engine.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *VipRightsMap) Update() error {
	if _, err := common.Engine.Where("id=?", m.Id).Update(m); err != nil {
		return err
	}
	return nil
}

func (m * VipRightsMap) Delete() error {
	if _, err := common.Engine.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *VipRightsMap) QueryById(id int) (*VipRightsMap, bool, error) {
	has, err := common.Engine.Where("id=?", id).Get(m)
	return m, has, err
}

func (m * VipRightsMap) QueryByVipLevel(vipLevel int) (*VipRightsMap, bool, error){
	has, err := common.Engine.Where("vip_level=?", vipLevel).Get(m)
	return m, has, err
}