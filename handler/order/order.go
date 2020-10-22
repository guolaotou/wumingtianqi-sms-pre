package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"wumingtianqi/handler"
	"wumingtianqi/libs/order"
	orderModel "wumingtianqi/model/order"
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
	type PostParams struct {
		Telephone   string                       `json:"telephone"`
		City        string                       `json:"city"`
		RemindTime  string                       `json:"remind_time"`
		OrderDetail []orderModel.OrderDetailItem `json:"order_detail"`
	}
	postParams := &PostParams{}
	if err := context.BindJSON(&postParams); err != nil {
		err = errnum.New(errnum.ErrParsingPostJson, err)
		fmt.Println("err: " + err.Error())
		handler.SendResponse(context, err, nil)
		return
	}
	fmt.Println("postParams", postParams)
	userId := context.GetHeader("X-User-Id")
	userIdInt, _ := strconv.Atoi(userId)
	resultData, err := order.AddUserOrderTel(
		userIdInt, postParams.Telephone, postParams.City,
		postParams.RemindTime, postParams.OrderDetail)
	if err != nil {
		handler.SendResponse(context, err, nil)
		return
	}
	handler.SendResponse(context, errnum.OK, resultData)
	return
}
