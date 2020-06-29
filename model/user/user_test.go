package user

import (
	"testing"
	"wumingtianqi-sms-pre/model/common"
	test "wumingtianqi-sms-pre/testing"
	"wumingtianqi-sms-pre/model/user"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test model/user/user_test.go
// go test -v model/user/user_test.go

func TestUserToRemind(t *testing.T) {
	test.Setup()
	session := common.Engine.NewSession()
	defer session.Close()

	// 1. 新建
	subscriberId := 1
	utr := &user.UserToRemind{
		SubscriberId:  subscriberId,
		SubscriberName: "路飞",
		TelephoneNum:  "13800380038",
	}
	t.Log("*** begin create session******")
	if err := utr.Create(); err != nil {
		panic(err)
	}

	// 2. 查询
	t.Log("*** begin query session******")

	utr2, has, err := user.QueryById(subscriberId)
	if err != nil || !has {
		t.Error("rc not found")
	} else {
		t.Log("rc: ", utr2)
	}
	t.Log("*** end query session****** ")

	// 3. 更改
	t.Log("*** begin update session****** ")
	// todo
	t.Log("*** end update session****** ")

	// 4. 删除
	t.Log("*** begin delete session****** ")
	//_ = utr.Delete()
	t.Log("*** begin delete session****** ")
}