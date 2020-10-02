package user

import (
	"testing"
	"time"
	"wumingtianqi/model/common"
	test "wumingtianqi/testing"
	"wumingtianqi/model/user"
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
	//m2 := new(user.UserInfo)
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

/*
type Invitation struct {
	Id             int       `json:"id" xorm:"pk autoincr INT(11)"`
	InvitationCode string    `json:"invitation_code" xorm:"VARCHAR(100) unique index"`
	TimesMax       int       `json:"times_max" xorm:"INT(11)"`
	TimesRemaining int       `json:"times_remaining" xorm:"INT(11)"`
	Vip            int       `json:"vip" xorm:"INT(11)"`
	Duration       int       `json:"duration" xorm:"INT(11) default 0"`
	Coin           int       `json:"coin" xorm:"INT(20) default 0"`
	Diamond        int       `json:"diamond" xorm:"INT(11) default 0"`
	Creator        int       `json:"diamond" xorm:"INT(11) default -1"`
	CreateTime     time.Time `json:"create_time" xorm:"TIMESTAMP"`
	UpdateTime     time.Time `json:"update_time" xorm:"TIMESTAMP"`
}
 */
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