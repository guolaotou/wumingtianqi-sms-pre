package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"wumingtianqi/config"
	test "wumingtianqi/testing"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test handler/user/invitation_test.go
// go test -v handler/user/invitation_test.go
// 测试之前先运行项目: go run main
func TestGetInvitationReward(t *testing.T) {
	test.Setup()
	webConfig := config.GlobalConfig.Web

	// 参数拼接
	host := webConfig.Host
	port := webConfig.Port
	baseUrl := fmt.Sprintf("http://%s:%s/v1/invitation/reward/get", host, port)

	// post参数拼接
	header := map[string]string {
		"X-WuMing-Token": "token111",
	}
	invitationCode := "b2772ec266-3b7-ca9f18-e20f255039"
	body := fmt.Sprintf(`{"invitation_code": "%s"}`, invitationCode)
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
	if resp != nil{
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
