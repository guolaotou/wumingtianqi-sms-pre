package city

import (
	"wumingtianqi/model/common"
)

type City struct {
	Id       int    `json:"id" xorm:"pk autoincr INT(11)"`
	Province string `json:"province" xorm:"VARCHAR(20)"`
	City     string `json:"city" xorm:"VARCHAR(20)"`
	District string `json:"district" xorm:"VARCHAR(20)"`
	PinYin   string `json:"pin_yin" xorm:"index VARCHAR(30)"`
	Abbr     string `json:"abbr" xorm:"VARCHAR(60)"`
	Code     string `json:"code" xorm:"index VARCHAR(32) default 'A' comment('xinzhi code')"`
}

func (c * City) Create() error {
	if _, err := common.Engine.InsertOne(c); err != nil {
		return err
	}
	return nil
}

func (c *City) Update() error {
	if _, err := common.Engine.Where("id=?", c.Id).Update(c); err != nil {
		return err
	}
	return nil
}

func (c *City) Delete() error {
	if _, err := common.Engine.Delete(c); err != nil {
		return err
	}
	return nil
}

func QueryByCityCode(cityCode string) (*City, bool, error) {
	c := new(City)
	has, err := common.Engine.Where("code=?", cityCode).Get(c)
	return c, has, err
}

func GetAllCity() ([]City, error) {  // todo bug fixed
	cityList := make([]City, 0)
	err := common.Engine.Find(&cityList)
	return cityList, err
}
