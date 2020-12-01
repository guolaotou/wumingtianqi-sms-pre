package wx

import (
	"github.com/gin-gonic/gin"
	"log"

	//"nevermore/log"
	"wumingtianqi/handler"
	"wumingtianqi/libs/wx"
	"wumingtianqi/utils/errnum"
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
		return
	}
	res, err := wx.WxLogin(wechatCode)  // lib函数
	log.Println("res", res)
	if err != nil {
		log.Println(err.Error())
		handler.SendResponse(context, err, nil)
		return
	}
	handler.SendResponse(context, errnum.OK, res)
	return
}

