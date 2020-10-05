package user

import (
	"testing"
	"time"
	"wumingtianqi/model/common"
	"wumingtianqi/model/user"
	test "wumingtianqi/testing"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test model/user/user_test.go
// go test -v model/user/user_test.go
func TestUserToRemind(t *testing.T) {
	test.Setup()
	session := common.Engine.NewSession()
	defer session.Close()

	// 1. 新建
	subscriberId := 3
	utr := &user.UserToRemind{
		SubscriberId:   subscriberId,
		SubscriberName: "路飞",
		TelephoneNum:   "13800380038",
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
	_ = utr.Delete()
	t.Log("*** begin delete session****** ")
}

func TestUserInfo(t *testing.T) {
	test.Setup()
	session := common.Engine.NewSession()
	defer session.Close()

	// 1. 新建
	id := 4
	currentTime := time.Now()
	m := &user.UserInfo{
		Id:         id,
		WxOpenId:   "open_id111",
		WxUnionId:  "union_id111",
		CreateTime: currentTime,
		UpdateTime: currentTime,
	}
	t.Log("*** begin create session******")
	if err := m.Create(); err != nil {
		panic(err)
	}

	// 2. 查询
	t.Log("*** begin query session******")
	m2, has, err := m.QueryById(id)
	if err != nil || !has {
		t.Error("model not found")
	} else {
		t.Log("model: ", m2)
	}
	t.Log("*** end query session****** ")

	// 3. 更改
	// 4. 删除
	_ = m.Delete()
}

// go test -v model/user/user_test.go -test.run TestUserInfoFlexible
func TestUserInfoFlexible(t *testing.T) {
	test.Setup()
	session := common.Engine.NewSession()
	defer session.Close()

	// 1. 新建
	userId := 3
	currentTime := time.Now()
	m := &user.UserInfoFlexible{
		UserId:         userId,
		VipLevel:       1,
		Coin:           10,
		Diamond:        10,
		ExpirationTime: 20201003,
		Creator:        -1,
		CreateTime:     currentTime,
		UpdateTime:     currentTime,
	}
	t.Log("*** begin create session******")
	if err := m.Create(); err != nil {
		panic(err)
	}

	// 2. 查询
	t.Log("*** begin query session******")
	m2, has, err := m.QueryByUserId(userId)
	if err != nil || !has {
		t.Error("model not found")
	} else {
		t.Log("model: ", m2)
	}
	t.Log("*** end query session****** ")

	// 4. 删除
	_ = m.Delete()
}

func TestInvitation(t *testing.T) {
	test.Setup()
	session := common.Engine.NewSession()
	defer session.Close()

	// 1.新建
	id := 100
	currentTime := time.Now()
	m := &user.Invitation{
		Id:             id,
		InvitationCode: "xx11tyuk",
		TimesMax:       10,
		TimesRemaining: 10,
		Vip:            1,
		Duration:       100,
		Coin:           10000,
		Diamond:        10000,
		Creator:        -1,
		CreateTime:     currentTime,
		UpdateTime:     currentTime,
	}

	t.Log("*** begin create session******")
	if err := m.Create(); err != nil {
		panic(err)
	}
	// 2. 查询
	t.Log("*** begin query session******")
	m2, has, err := m.QueryById(id)
	if err != nil || !has {
		t.Error("model not found")
	} else {
		t.Log("model", m2)
	}
	t.Log("*** end query session****** ")

	// 3. 更改
	// 4. 删除
	_ = m.Delete()
}
