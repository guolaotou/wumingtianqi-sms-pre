package user

import (
	"time"
	"wumingtianqi/model/common"
)

// 用户信息表（微信信息存储）
type UserInfo struct {
	Id         int       `json:"id" xorm:"pk autoincr INT(11)"`
	WxOpenId   string    `json:"wx_open_id" xorm:"VARCHAR(100) default '' comment('微信open_id') index"`
	WxUnionId  string    `json:"wx_union_id" xorm:"VARCHAR(100) default '' comment('微信union_id') index"`
	UserToken  string    `json:"user_token" xorm:"VARCHAR(100) default '' comment('用户token，前期用微信session key')"`
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

func (m *UserInfo) QueryByUserToken(userToken string) (*UserInfo, bool, error) {
	has, err := common.Engine.Where("user_token=?", userToken).Get(m)
	return m, has, err
}

// 用户灵活信息表
type UserInfoFlexible struct {
	UserId                   int       `json:"user_id" xorm:"pk INT(11)"`
	InvitationCode           string    `json:"invitation_code" xorm:"VARCHAR(100) default '' comment('邀请码')"`
	VipLevel                 int       `json:"vip_level" xorm:"INT(3) default 0"`
	WechatOrderRemaining     int       `json:"wechat_order_remaining" xorm:"INT(3) default 0 comment('微信订单剩余配置数')"`
	TelOrderRemaining        int       `json:"tel_order_remaining" xorm:"INT(3) default 0 comment('手机号订单剩余配置数')"`
	TodayTelRemindRemaining  int       `json:"today_tel_remind_remaining" xorm:"INT(3) default 0 comment('短信提醒当天剩余次数')"`
	LastRemindTime           int       `json:"last_remind_time" xorm:"INT(11) default 20000101 comment('上次提醒时间')"`
	Coin                     int       `json:"coin" xorm:"INT(20) default 0"`
	Diamond                  int       `json:"diamond" xorm:"INT(11) default 0"`
	ExpirationTime           int       `json:"expiration_time" xorm:"INT(11) default 20000101"`
	Creator                  int       `json:"creator" xorm:"INT(11) default -1"`
	CreateTime               time.Time `json:"create_time" xorm:"TIMESTAMP"`
	UpdateTime               time.Time `json:"update_time" xorm:"TIMESTAMP"`
}

func (m * UserInfoFlexible) Create() error {
	if _, err := common.Engine.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInfoFlexible) Update() error {
	if _, err := common.Engine.AllCols().Where("user_id=?", m.UserId).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInfoFlexible) Delete() error {
	if _, err := common.Engine.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInfoFlexible) QueryByUserId(userId int) (*UserInfoFlexible, bool, error) {
	has, err := common.Engine.Where("user_id=?", userId).Get(m)
	return m, has, err
}

// 存放邀请码的表 wiki: https://github.com/guolaotou/wumingtianqi-sms-pre/wiki/%E9%82%80%E8%AF%B7%E5%88%86%E7%BA%A7%E5%A5%96%E5%8A%B1%E6%95%B0%E6%8D%AE%E5%BA%93%E8%AE%BE%E8%AE%A1%E4%B8%8E%E6%8E%A5%E5%8F%A3%E8%AE%BE%E8%AE%A1
type Invitation struct {
	Id             int       `json:"id" xorm:"pk autoincr INT(11)"`
	InvitationCode string    `json:"invitation_code" xorm:"VARCHAR(100) unique index"`
	TimesMax       int       `json:"times_max" xorm:"INT(11)"`
	TimesRemaining int       `json:"times_remaining" xorm:"INT(11)"`
	Vip            int       `json:"vip" xorm:"INT(11)"`
	Duration       int       `json:"duration" xorm:"INT(11) default 0"`
	Coin           int       `json:"coin" xorm:"INT(20) default 0"`
	Diamond        int       `json:"diamond" xorm:"INT(11) default 0"`
	Creator        int       `json:"creator" xorm:"INT(11) default -1"`
	CreateTime     time.Time `json:"create_time" xorm:"TIMESTAMP"`
	UpdateTime     time.Time `json:"update_time" xorm:"TIMESTAMP"`
}

func (m *Invitation) Create() error {
	if _, err := common.Engine.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *Invitation) Update() error {
	if _, err := common.Engine.AllCols().Where("id=?", m.Id).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *Invitation) Delete() error {
	if _, err := common.Engine.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Invitation) QueryById(id int)(*Invitation, bool, error) {
	i := new(Invitation)
	has, err := common.Engine.Where("id=?", id).Get(i)
	return i, has, err
}

func (m *Invitation) QueryByInvitationCode(invitationCode string)(*Invitation, bool, error) {
	//i := new(Invitation)
	has, err := common.Engine.Where("invitation_code=?", invitationCode).Get(m)
	return m, has, err
}

// 用户-邀请码映射表
type UserInvitationMap struct {
	Id int `json:"id" xorm:"pk autoincr INT(11)"`
	UserId int `json:"user_id" xorm:"unique(user_invitation) INT(11)"`
	InvitationCode string `json:"invitation_code" xorm:"unique(user_invitation) VARCHAR(100)"`
	CreateTime     time.Time `json:"create_time" xorm:"TIMESTAMP"`
}

func (m *UserInvitationMap)Create() error {
	if _, err := common.Engine.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInvitationMap) Update() error {
	if _, err := common.Engine.AllCols().Where("id=?", m.Id).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInvitationMap) Delete() error {
	if _, err := common.Engine.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m * UserInvitationMap) QueryByUserId(userId int) ([]UserInvitationMap, error) {
	modelList := make([]UserInvitationMap, 0)
	err := common.Engine.Where("user_id=?", userId).Find(&modelList)
	return modelList, err
}

func (m *UserInvitationMap) QueryByInvitationCode(invitationCode string) ([]UserInvitationMap, error) {
	modelList := make([]UserInvitationMap, 0)
	err := common.Engine.Where("invitation_code=?", invitationCode).Find(&modelList)
	return modelList, err
}

func (m *UserInvitationMap) QueryByUserInvitation(userId int, invitationCode string) (*UserInvitationMap, bool, error) {
	has, err := common.Engine.Where("user_id=?", userId).And("invitation_code=?", invitationCode).Get(m)
	return m, has, err
}