package router

import (
	"github.com/gin-gonic/gin"
	"wumingtianqi/handler/wx"
)


func InitRouter(router *gin.Engine) {
	//router := gin.Default()
	router.GET("/test", )
	//router.POST("/wx/login", wx.WxLogin)
	router.GET("/wx/login", wx.WxLogin)
}