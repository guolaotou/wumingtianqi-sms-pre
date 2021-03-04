package main

import (
	"fmt"
	"time"
	"wumingtianqi/config"
	"wumingtianqi/model"
	"wumingtianqi/model/vip"
)

/**
 * @Author Evan
 * @Description vip_rights_map表添加数据
 * @Date 16:18 2020-10-07
 * @Param
 * @return 
 **/
func addDataVipRightsMap() error {
	if _, err := config.LoadConfig(); err != nil {
		fmt.Println(err.Error())
	}
	model.InitMysql()
	// Vip0 新建
	currentTime := time.Now()
	m := &vip.VipRightsMap{
		Id:                  1,
		VipLevel:            0,
		WechatOrderMax:      3,
		TelOrderMax:         0,
		RemindPatternIdList: "[1,2,3,8]",
		CreateTime:          currentTime,
		UpdateTime:          currentTime,
	}
	if err := m.Create(); err != nil {
		panic(err)
	}

	// Vip1
	m = &vip.VipRightsMap{
		Id:                  2,
		VipLevel:            1,
		WechatOrderMax:      4,
		TelOrderMax:         0,
		RemindPatternIdList: "[1,2,3,4,6,7,8]",
		CreateTime:          currentTime,
		UpdateTime:          currentTime,
	}
	if err := m.Create(); err != nil {
		panic(err)
	}
	// Vip2
	m = &vip.VipRightsMap{
		Id:                  3,
		VipLevel:            2,
		WechatOrderMax:      5,
		TelOrderMax:         1,
		RemindPatternIdList: "[1,2,3,4,5,6,7,8]",
		CreateTime:          currentTime,
		UpdateTime:          currentTime,
	}
	if err := m.Create(); err != nil {
		panic(err)
	}
	// Vip3
	m = &vip.VipRightsMap{
		Id:                  4,
		VipLevel:            3,
		WechatOrderMax:      6,
		TelOrderMax:         6,
		RemindPatternIdList: "[1,2,3,4,5,6,7,8]",
		CreateTime:          currentTime,
		UpdateTime:          currentTime,
	}
	if err := m.Create(); err != nil {
		panic(err)
	}
	return nil
}

// go run scripts/vip/vip.go
func main() {
	err := addDataVipRightsMap()
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
}