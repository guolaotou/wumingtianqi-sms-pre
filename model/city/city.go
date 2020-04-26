package city

import (
	"fmt"
	"wumingtianqi-sms-pre/model/common"
)

type City struct {
	Id       int    `json:"id" xorm:"pk autoincr INT(11)"`
	Province string `json:"province" xorm:"VARCHAR(20)"`
	City     string `json:"city" xorm:"VARCHAR(20)"`
	District string `json:"district" xorm:"VARCHAR(20)"`
	PinYin   string `json:"pin_yin" xorm:"VARCHAR(30)"`
	Abbr     string `json:"abbr" xorm:"VARCHAR(60)"`
}

func (u *City) TableName() string {
	return "city"
}

func GetAllCity() ([]City, error) {
	cityList := make([]City, 0)
	err := common.Engine.Find(&cityList)
	return cityList, err
}

func GetOneCity()(*City, error) {
	city := new(City)
	has, err := common.Engine.Where("id=?", 1).Get(city)
	fmt.Println("has", has)
	return city, err
}
