package weather

import (
	"github.com/gin-gonic/gin"
	"wumingtianqi/handler"
	"wumingtianqi/libs/weather"
	"wumingtianqi/utils/errnum"
)

/**
 * @Author Evan
 * @Description 获取天气列表
 * @Date 21:18 2021-02-20
 * @Param 
 * @return 
 **/
// todo 接口文档
func GetCityList(context *gin.Context) {
	// todo defer RecoverError
	resultData, err := weather.GetCityList()
	if err != nil {
		handler.SendResponse(context, err, nil)
		return
	}
	handler.SendResponse(context, errnum.OK, resultData)
	return
}