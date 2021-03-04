package order

import (
	"fmt"
	"testing"
	orderLib "wumingtianqi/libs/order"
	test "wumingtianqi/testing"
)

// go clean -testcache && go test -v libs/order/order_test.go -test.run TestOrder
// 注释libs/order.go 里的 `common.PubSub.Publish`
func TestOrder(t *testing.T) {
	test.Setup()
	//orderLib.FakeWeather()
	// 参数：时间，例如0900；查询所有order表中时间等于0900的
	orderLib.ProcessOrdersOfTime("1434")
}

// go clean -testcache && go test -v libs/order/order_test.go -test.run TestGetUserOrderTel
func TestGetUserOrderTel(t *testing.T) {
	test.Setup()
	userId := 1
	resultData, err := orderLib.GetUserOrderTel(userId)
	fmt.Println("resultData", resultData)
	fmt.Println("err", err)
}
