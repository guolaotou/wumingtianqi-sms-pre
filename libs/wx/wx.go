package wx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wumingtianqi-sms-pre/config"
)

func WxLogin(wechatCode string) int {
	/* 对应API：用户微信登录
		1.jscode -> open_id
		2.查数据库user_info表，查该open_id是否有对应model
		2.1 若有，直接返回id；
		2.2 若没有，则新建model，返回id
	*/
	// step 1： 获取用户open_id
	res := GetUserOpenId(wechatCode)
	fmt.Println(res)
	return 1
}

func GetUserOpenId(wechatCode string) string {
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
