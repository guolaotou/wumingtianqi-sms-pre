package wx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wumingtianqi/config"
	test "wumingtianqi/testing"
	"testing"
)

//  go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
//  go test handler/wx/wx_test.go
//  go test -v handler/wx/wx_test.go

func TestWxLogin(t *testing.T) {
	// 这个测试用例需要保持服务运行。在项目根目录: go run main.go
	test.Setup()
	webConfig := config.GlobalConfig.Web

	// 参数拼接
	host := webConfig.Host
	port := webConfig.Port
	baseUrl := fmt.Sprintf("http://%s:%s/wx/login", host, port)

	wechatCode := "051RKGGa1TjcAz0ShxFa14CL7Z2RKGGP"  // 填写拿到的微信临时登录凭证code
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