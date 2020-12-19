package router

import (
	"github.com/gin-gonic/gin"
	"wumingtianqi/handler/order"
	"wumingtianqi/handler/user"
	"wumingtianqi/handler/wx"
	"wumingtianqi/router/middleware"
)

func InitRouter(router *gin.Engine) {
	//router := gin.Default()
	router.GET("/test", )
	//router.POST("/wx/login", wx.WxLogin)
	router.GET("/wx/login", wx.WxLogin)

	// user
	router.GET("/v1/user/info", middleware.TokenParsing, user.GetUserInfo) // 获取用户信息
	router.POST("/v1/invitation/reward/get", middleware.TokenParsing, user.GetInvitationReward) // 邀请码获取奖励

	// 手机号订单
	router.POST("/v1/user/order/tel/add", middleware.TokenParsing, order.AddUserOrderTel)
	router.GET("/v1/user/order/tel/get", middleware.TokenParsing, order.GetUserOrderTel)
	router.POST("/v1/user/order/tel/delete", middleware.TokenParsing, order.DeleteUserOrderTel)
	router.POST("/v1/user/order/tel/update", middleware.TokenParsing, order.UpdateUserOrderTel)
}
