package order

import (
	"testing"
	test "wumingtianqi/testing"
	"wumingtianqi/model/common"
	"wumingtianqi/model/order"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test model/order/remind_test.go
// go test -v model/order/remind_test.go

func TestRemindCondition(t *testing.T)  {
	test.Setup()
	session := common.Engine.NewSession()
	defer session.Close()

	// 1. 新建
	id := 1
	rc := &order.RemindCondition{
		Id:            id,
		WeatherId:     5,
		Variety:       "降低",
		Value:         "10",
		FormatText:    "%s降低%d度",
		Tips:          "注意保暖",
		Attribution:   "admin",
		Priority:      1,
		ConfigGroupId: 1,
	}
	t.Log("*** begin create session******")

	if err := rc.Create(); err != nil {
		panic(err)
	}

	// 2. 查询
	t.Log("*** begin query session******")

	rc2, has, err := order.QueryById(id)
	if err != nil || !has {
		t.Error("rc not found")
	} else {
		t.Log("rc: ", rc2)
	}
	t.Log("*** end query session****** ")

	// 3. 更改
	t.Log("*** begin update session****** ")
	// todo
	t.Log("*** end update session****** ")

	// 4. 删除
	t.Log("*** begin delete session****** ")
	_ = rc.Delete()
	t.Log("*** end delete session****** ")
}