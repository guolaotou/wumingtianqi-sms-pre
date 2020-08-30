package wx

import (
	"github.com/gin-gonic/gin"
	"wumingtianqi-sms-pre/libs/wx"
)


// API：用户微信登录
func WxLogin (context *gin.Context){
	/*
		1.jscode -> open_id
		2.查数据库user_info表，查该open_id是否有对应model
		2.1 若有，直接返回id；
		2.2 若没有，则新建model，返回id
	*/
	wx.WxLogin("1")

}

