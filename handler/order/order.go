package order

import (
	"github.com/gin-gonic/gin"
	"wumingtianqi/handler"
	"wumingtianqi/libs/order"
	"wumingtianqi/utils/errnum"
)

// API：建立订单
/**
 * @Author Evan
 * @Description 新增手机号提醒订单，接口函数
 * @Date 18:42 2020-10-07
 * @Param context *gin.Context
 * @return 
 **/
func AddUserOrderTel(context *gin.Context) {
	// post参数解析，然后调用lib函数；

	resultData, err := order.AddUserOrderTel()
	if err != nil {
		handler.SendResponse(context, err, nil)
	}
	handler.SendResponse(context, errnum.OK, resultData)
}