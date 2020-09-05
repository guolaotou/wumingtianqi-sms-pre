package wx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"wumingtianqi/config"
	"wumingtianqi/model/user"
	"wumingtianqi/utils/errnum"
)

func WxLogin(wechatCode string) (int, error) {
	/* 对应API：用户微信登录
		1.jscode -> open_id
		2.查数据库user_info表，查该open_id是否有对应model
		2.1 若有，直接返回id；
		2.2 若没有，则新建model，返回id
	*/
	// step 1： 获取用户open_id
	openId, sessionKey, err := GetUserOpenId(wechatCode)
	fmt.Println("openId",openId, "sessionKey: ", sessionKey)
	if err != nil {
		return -1, err
	}

	// step2: 查数据库user_info表，查该open_id是否有对应model
	u := new(user.UserInfo)
	u, has, err := u.QueryByOpenId(openId)

	if err != nil {
		err = errnum.New(errnum.DbError, nil)
		return -1, err
	} else if has == false {
		// 新建
		currentTime := time.Now()
		u.WxOpenId = openId
		u.WxUnionId = ""
		u.CreateTime = currentTime
		u.UpdateTime = currentTime
		if err := u.Create(); err != nil {
			err = errnum.New(errnum.DbError, errors.New("create sql error"))
			return -1, err
		}
	}

	userId := u.Id
	return userId, nil
}

func GetUserOpenId(wechatCode string) (string, string, error) {
	/*  获取用户open_id
	wechatCode: 小程序端调用wx.login获取的用户临时登录凭证code; todo golang标准注释学一下
	参考文档：// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html

	返回值：openId, sessionKey, error
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
	// 错误处理 todo 日后可以封装
	if err != nil {
		err = errnum.New(errnum.WxError, err)
		// todo 封装 log
		return "", "", err
	} else if resp == nil {
		err = errnum.New(errnum.WxError, errors.New("resp is None"))
		return "", "", err
	}
	code := resp.StatusCode
	if code != 200 {
		err = errnum.New(errnum.WxError, errors.New("Http error code: " + strconv.Itoa(code)))
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	resultData := map[string]interface{} {}
	err = json.Unmarshal(body, &resultData)
	if err != nil {
		err = errnum.New(errnum.WxError, errors.New("body unmarshal error"))
		return "", "", err
	}
	// 真正开始解析返回值
	// 1. 错误 {"errcode":40029,"errmsg":"invalid code, hints: [ req_id: zhBes.5ce-2V7a_ ]"}
	errCode := resultData["errcode"]
	if errCode != nil {
		errMsg := resultData["errmsg"]
		errMsgStr := ""
		if errMsg != nil {
			errMsgStr = errMsg.(string)
		}
		err = errnum.New(errnum.WxError, errors.New(fmt.Sprintf("Wx error code: %.f; Wx error desc: %s", errCode.(float64), errMsgStr)))
		return "", "", err
	}

	// 2. 正常 {"session_key":"xxx","openid":"xxx"}
	// 拿到open_id 和session_id
	openId := resultData["openid"]
	sessionKey := resultData["session_key"]
	if openId == nil || sessionKey == nil {
		err = errnum.New(errnum.WxError, errors.New("body unmarshal error open_id"))
		return "", "", err
	}
	openIdStr := openId.(string)
	sessionKeyStr := sessionKey.(string)

	return openIdStr, sessionKeyStr, nil
}
