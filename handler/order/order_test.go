package order

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"wumingtianqi/config"
	orderModel "wumingtianqi/model/order"
	test "wumingtianqi/testing"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test handler/order/order_test.go
// go clean -testcache && go test -v handler/order/order_test.go
// 测试之前先运行项目: go run main
func TestAddUserOrderTel(t *testing.T) {
	test.Setup()
	webConfig := config.GlobalConfig.Web

	// 参数拼接
	host := webConfig.Host
	port := webConfig.Port
	baseUrl := fmt.Sprintf("http://%s:%s/v1/user/order/tel/add", host, port)

	// post参数拼接
	header := map[string]string {
		"X-WuMing-Token": "token111",
	}
	preTele := "+86"
	telephone := "18812341234"
	city := "haidian"
	remindTime := "2222"
	var orderDetailList []orderModel.OrderDetailItem
	orderDetailList = append(orderDetailList, orderModel.OrderDetailItem{
		RemindPatternId: 1,
		Value: -999,
	})
	orderDetailList = append(orderDetailList, orderModel.OrderDetailItem{
		RemindPatternId: 2,
		Value: 5,
	})
	bufOrderDetailList, _ := json.Marshal(orderDetailList)
	body := fmt.Sprintf(`{
		"pre_tele":"%s",
		"telephone":"%s",
		"city":"%s",
		"remind_time":"%s",
		"order_detail":%s
	}`,preTele, telephone, city, remindTime, string(bufOrderDetailList))
	fmt.Println("body: ", body)
	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(body))
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}
	for header, value := range header {  // 循环添加header，这里是加token
		req.Header.Add(header, value)
	}
	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}
	statusCode := resp.StatusCode
	fmt.Println("statusCode", statusCode)
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("err ioutil.ReadAll", err.Error())
	} else {
		fmt.Println("resBody", string(resBody))
	}
}


// go clean -testcache && go test -v handler/order/order_test.go -test.run TestGetUserOrderTel
// 测试之前先运行项目: go run main
func TestGetUserOrderTel(t *testing.T) {
	test.Setup()
	webConfig := config.GlobalConfig.Web

	// 参数拼接
	host := webConfig.Host
	port := webConfig.Port
	baseUrl := fmt.Sprintf("http://%s:%s/v1/user/order/tel/get", host, port)
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}

	// header添加
	header := map[string]string {
		"X-WuMing-Token": "token111",
	}
	for header, value := range header {
		req.Header.Add(header, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}
	statusCode := resp.StatusCode
	fmt.Println("statusCode", statusCode)
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("err ioutil.ReadAll", err.Error())
	} else {
		fmt.Println("resBody", string(resBody))
	}
}

// go clean -testcache && go test -v handler/order/order_test.go -test.run TestDeleteUserOrderTel
// 测试之前先运行项目: go run main
func TestDeleteUserOrderTel(t *testing.T) {
	test.Setup()
	webConfig := config.GlobalConfig.Web

	// 参数拼接
	host := webConfig.Host
	port := webConfig.Port
	baseUrl := fmt.Sprintf("http://%s:%s/v1/user/order/tel/delete", host, port)
	// post参数拼接
	header := map[string]string {
		"X-WuMing-Token": "token111",
	}
	orderId := 26
	body := fmt.Sprintf(`{"order_id":%d}`, orderId)
	fmt.Println("body", body)

	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(body))
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}
	for header, value := range header {  // 循环添加header，这里是加token
		req.Header.Add(header, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}
	statusCode := resp.StatusCode
	fmt.Println("statusCode", statusCode)
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("err ioutil.ReadAll", err.Error())
	} else {
		fmt.Println("resBody", string(resBody))
	}
}

// go clean -testcache && go test -v handler/order/order_test.go -test.run TestUpdateUserOrderTel
// 测试之前先运行项目：go run main
func TestUpdateUserOrderTel(t *testing.T) {
	test.Setup()
	webConfig := config.GlobalConfig.Web

	// 参数拼接
	host := webConfig.Host
	port := webConfig.Port
	baseUrl := fmt.Sprintf("http://%s:%s/v1/user/order/tel/update", host, port)

	// post参数拼接
	header := map[string]string {
		"X-Wuming-Token": "token111",
	}
	orderId := 25
	preTele := "+86"
	telephone := "18812341234"
	city := "chaoyang"
	remindTime := "2222"
	var orderDetailList []orderModel.OrderDetailItem
	orderDetailList = append(orderDetailList, orderModel.OrderDetailItem{
		RemindPatternId: 1,
		Value: -999,
	})
	orderDetailList = append(orderDetailList, orderModel.OrderDetailItem{
		RemindPatternId: 2,
		Value: 5,
	})
	bufOrderDetailList, _ := json.Marshal(orderDetailList)
	body := fmt.Sprintf(`{
		"order_id":%d,
		"pre_tele":"%s",
		"telephone":"%s",
		"city":"%s",
		"remind_time":"%s",
		"order_detail":%s
	}`,orderId, preTele, telephone, city, remindTime, string(bufOrderDetailList))
	log.Println("body: ", body)
	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(body))
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}
	for header, value := range header {
		req.Header.Add(header, value)
	}
	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}
	statusCode := resp.StatusCode
	fmt.Println("statusCode", statusCode)
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("err ioutil.ReadAll", err.Error())
	} else {
		fmt.Println("resBody", string(resBody))
	}
}