package wx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nevermore/log"
	"wumingtianqi-sms-pre/handler"
	"wumingtianqi-sms-pre/libs/wx"
	"wumingtianqi-sms-pre/utils/errnum"
)

// 下面是封装返回值的，包括error
type ResponseShanHuResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time int64    `json:"time"`
	Data string `json:"data"`
}
// API：用户微信登录
func WxLogin (context *gin.Context){
	/*
		1.jscode -> open_id
		2.查数据库user_info表，查该open_id是否有对应model
		2.1 若有，直接返回id；
		2.2 若没有，则新建model，返回id
	*/
	fmt.Println("context 内容", context)
	wx.WxLogin("1")
	var err error // 其实这个应该是lib层返回的
	err = nil
	if err != nil {
		err = errnum.New(errnum.S2SError, err)
		log.L().Error(err.Error())  // todo 以后再封装log模块
		handler.SendResponse(context, err, nil)
	}

	res := map[string]interface{} {
		"user_id": 1,
	}
	handler.SendResponse(context, errnum.OK, res)
	return
}

