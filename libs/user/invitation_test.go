package user

import (
	"fmt"
	"testing"
	test "wumingtianqi/testing"
	userLib "wumingtianqi/libs/user"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test libs/user/invitation_test.go
// go test -v libs/user/invitation_test.go
func TestInvitation(t *testing.T) {
	test.Setup()
	userId := 1
	invitationCode := "b2772ec266-3b7-ca9f18-e20f255039"
	resultData, err := userLib.GetInvitationReward(userId, invitationCode)
	if err != nil {
		println("err:", err.Error())
	}
	fmt.Println("resultData", resultData)
}
