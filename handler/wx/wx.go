package wx

import (
	"github.com/gin-gonic/gin"
	"nevermore/log"
	"wumingtianqi-sms-pre/handler"
	"wumingtianqi-sms-pre/libs/wx"
	"wumingtianqi-sms-pre/utils/errnum"
)

// API：用户微信登录
func WxLogin (context *gin.Context){
	/*
		1.jscode -> open_id
		2.查数据库user_info表，查该open_id是否有对应model
		2.1 若有，直接返回id；
		2.2 若没有，则新建model，返回id
	*/
	wechatCode := context.Query("wechatCode")
	if wechatCode == "" {
		handler.SendResponse(context, errnum.ParamsError, nil)
	}
	userId, err := wx.WxLogin(wechatCode)  // lib函数
	if err != nil {
		log.L().Error(err.Error())  // todo 以后再封装log模块
		handler.SendResponse(context, err, nil)
	}

	res := map[string]interface{} {
		"user_id": userId,
	}
	handler.SendResponse(context, errnum.OK, res)
	return
}

