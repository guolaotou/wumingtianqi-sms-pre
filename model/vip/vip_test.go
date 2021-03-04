package vip

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
	"wumingtianqi/model/vip"
	test "wumingtianqi/testing"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test model/vip/vip_test.go
// go test -v model/vip/vip_test.go
func TestVipRightsMap(t *testing.T) {
	// go test -v model/vip/vip_test.go -test.run TestVipRightsMap
	test.Setup()
	//session := common.Engine.NewSession()
	//defer session.Close()

	// 1. 新建
	id := 1
	vipLevel := 0
	currentTime := time.Now()
	m := &vip.VipRightsMap{
		Id:                  id,
		VipLevel:            vipLevel,
		WechatOrderMax:      3,
		TelOrderMax:         0,
		RemindPatternIdList: "[1]",
		CreateTime:          currentTime,
		UpdateTime:          currentTime,
	}
	t.Log("*** begin create session******")
	if err := m.Create(); err != nil {
		panic(err)
	}
	// 2. 查询
	t.Log("*** begin query session******")
	m, has, err := m.QueryById(id)
	if err != nil || !has {
		t.Error("model not found")
		panic("model not found")
	} else {
		t.Log("model: ", m)
	}
	t.Log("*** end query session****** ")

	// 3. 更改
	remindPatternIdListOldStr := m.RemindPatternIdList
	fmt.Println("remindPatternIdListOld", remindPatternIdListOldStr)
	var remindPatternIdList []int
	err = json.Unmarshal([]byte(remindPatternIdListOldStr), &remindPatternIdList)
	if err != nil {
		panic("json.Unmarshal" + err.Error())
	}
	// 如果解析出来的提醒模式id列表不为空，打印出第一个数字
	if remindPatternIdList != nil {
		println("first element: ", remindPatternIdList[0])
	}
	remindPatternIdList = append(remindPatternIdList, 999)
	remindPatternIdListNewBytes, err := json.Marshal(remindPatternIdList)
	if err != nil {
		panic("json.Marshal" + err.Error())
	}
	m.RemindPatternIdList = string(remindPatternIdListNewBytes)
	println("m.RemindPatternIdList", m.RemindPatternIdList)
	err = m.Update()
	if err != nil {
		panic("m.Update()" + err.Error())
	}

	// 4.删除
	t.Log("*** begin delete session****** ")
	err = m.Delete()
	if err != nil {
		panic("m.Delete()" + err.Error())
	}
	t.Log("*** end delete session****** ")
}
