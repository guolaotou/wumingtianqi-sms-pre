package user

import (
	"fmt"
	"testing"
	test "wumingtianqi/testing"
	userLib "wumingtianqi/libs/user"
)

// go clean -testcache && go test -v libs/user/user_info_test.go
func TestGetUserInfo(t *testing.T) {
	test.Setup()
	userId := 1
	resultData, err := userLib.GetUserInfo(userId)
	if err != nil {
		println("err: ", err.Error())
	}
	fmt.Println("resultData", resultData)
}
