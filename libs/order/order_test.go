package order

import (
	"testing"
	test "wumingtianqi/testing"
	orderLib "wumingtianqi/libs/order"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test libs/order/order_test.go
// go test -v libs/order/order_test.go
func TestCity(t *testing.T) {
	test.Setup()
	orderLib.FakeWeather()
	// 参数：时间，例如0900；查询所有order表中时间等于0900的
	orderLib.SpliceOrders("2300")
}