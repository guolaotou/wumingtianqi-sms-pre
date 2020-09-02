package wx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wumingtianqi-sms-pre/config"
	test "wumingtianqi-sms-pre/testing"
	"testing"
)

//  go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
//  go test handler/wx/wx_test.go
//  go test -v handler/wx/wx_test.go

func TestWxLogin(t *testing.T) {
	test.Setup()
	webConfig := config.GlobalConfig.Web

	// 参数拼接
	host := webConfig.Host
	port := webConfig.Port
	baseUrl := fmt.Sprintf("http://%s:%s/wx/login", host, port)

	wechatCode := "1"
	params := fmt.Sprintf("?wechatCode=%s", wechatCode)

	url := baseUrl + params

	// 请求
	resp, err := http.Get(url)
	if err != nil {
		println("err11", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("err2", err.Error())
	}
	println("body:", string(body))
}