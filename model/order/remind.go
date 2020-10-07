package order

import "wumingtianqi/model/common"

// 提醒对象对照表
type RemindObject struct {
	Id            int    `json:"id" xorm:"pk autoincr INT(11)"`
	WeatherPinYin string `json:"weather_pin_yin" xorm:"VARCHAR(30)"`
	WeatherObject string `json:"weather_object" xorm:"VARCHAR(30)"`
	WeatherFlag   string `json:"weather_flag" xorm:"VARCHAR(30)"`
}

// 提醒条件表
type RemindCondition struct {
	Id            int    `json:"id" xorm:"pk autoincr INT(11)"`
	WeatherId     int    `json:"weather_id" xorm:"INT(11)"`
	Variety       string `json:"variety" xorm:"VARCHAR(30)"`
	Value         string `json:"value" xorm:"VARCHAR(30)"`
	FormatText    string `json:"format_text" xorm:"VARCHAR(50)"`
	Tips          string `json:"tips" xorm:"VARCHAR(40)"`
	Attribution   string `json:"attribution" xorm:"VARCHAR(30)"`
	Priority      int    `json:"priority" xorm:"INT(11)"`
	ConfigGroupId int    `json:"config_group_id" xorm:"INT(11)"`
}

func (rc *RemindCondition) Create() error {
	if _, err := common.Engine.InsertOne(rc); err != nil {
		return err
	}
	return nil
}

func (rc *RemindCondition) Update() error {
	if _, err := common.Engine.Where("id=?", rc.Id).Update(rc); err != nil {
		return err
	}
	return nil
}

func (rc *RemindCondition) Delete() error {
	if _, err := common.Engine.Delete(rc); err != nil {
		return err
	}
	return nil
}

func QueryById(id int) (*RemindCondition, bool, error) {
	rc := new(RemindCondition)
	has, err := common.Engine.Where("id=?", id).Get(rc)
	return rc, has, err
}



