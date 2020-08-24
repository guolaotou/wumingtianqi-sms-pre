package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getUserWxAccessToken() string {
	baseUrl := "https://api.weixin.qq.com/sns/jscode2session"  // https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
	appId := "xxx"
	secret := "xxx"
	grantType := "authorization_code"
	jsCode := "xxx" // 小程序端调用wx.login获取的用户临时登录凭证code
	url := fmt.Sprintf("%s?appId=%s&secret=%s&grant_type=%s&js_code=%s", baseUrl, appId, secret, grantType, jsCode)

	resp, err :=   http.Get(url)
	if err != nil {
		println("err", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("err2", err.Error())
	}

	return string(body)
}

// go run scripts/wx_util/wx_util.go
func main() {
	/*
	获取用户accessToken
	参考文档：// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
	 */
	res := getUserWxAccessToken()
	fmt.Println("result:", res)
}


