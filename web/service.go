package web

import (
	"github.com/gin-gonic/gin"
	"nevermore/log"
	"wumingtianqi-sms-pre/config"
	"wumingtianqi-sms-pre/router"
)

func ListenHttp() {
	/* 监听端口（默认8001）
		test访问0.0.0.0:8001/v1/wx/login
	 */
	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()
	router.InitRouter(app)  // 路由设置
	log.L().Info("WEb server started")

	//host := config.GlobalConfig.Web.Host
	port := config.GlobalConfig.Web.Port
	err := app.Run(":" + port)
	if err != nil {
		log.L().Error("Failed to start web server")
	}
}