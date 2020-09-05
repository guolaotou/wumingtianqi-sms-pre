package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wumingtianqi/config"
)

func getUserOpenId(wechatCode string) string {
	/*  获取用户open_id
	wechatCode: 小程序端调用wx.login获取的用户临时登录凭证code; todo golang标准注释学一下
	参考文档：// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
	 */
	if _, err := config.LoadConfig(); err != nil {
		fmt.Println(err.Error())
	}
	wxConfig := config.GlobalConfig.Wx
	baseUrl := "https://api.weixin.qq.com/sns/jscode2session"  // https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
	appId := wxConfig.AppId
	secret := wxConfig.Secret
	grantType := "authorization_code"
	url := fmt.Sprintf("%s?appId=%s&secret=%s&grant_type=%s&js_code=%s", baseUrl, appId, secret, grantType, wechatCode)

	resp, err := http.Get(url)
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

func getWxAccessToken() string {
	/*
	获取微信access_token
	参考文档：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
	 */
	// todo 最后改成一个定时任务
	if _, err := config.LoadConfig(); err != nil {
		fmt.Println(err.Error())
	}
	wxConfig := config.GlobalConfig.Wx
	baseUrl := "https://api.weixin.qq.com/cgi-bin/token"
	grantType := "client_credential"
	appId := wxConfig.AppId
	secret := wxConfig.Secret

	url := fmt.Sprintf("%s?grant_type=%s&appid=%s&secret=%s", baseUrl, grantType, appId, secret)

	resp, err := http.Get(url)
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
	放一些微信小程序后台用的脚本
	 */
	//res := getUserOpenId("xxx")
	res := getWxAccessToken()
	fmt.Println("result:", res)
}


