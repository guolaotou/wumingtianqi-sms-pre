package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"wumingtianqi/config"
	test "wumingtianqi/testing"
)

// go clean -testcache && go test -v handler/user/user_info_test.go -test.run TestGetUserInfo
// 测试之前你先运行项目: go run main.go
func TestGetUserInfo(t *testing.T) {
	test.Setup()
	webConfig := config.GlobalConfig.Web

	// 参数拼接
	host := webConfig.Host
	port := webConfig.Port
	baseUrl := fmt.Sprintf("http://%s:%s/v1/user/info", host, port)
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}

	// header添加
	header := map[string]string {
		"X-Wuming-Token": "token111",
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
