package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
		PreTele     string                       `json:"pre_tele"`
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
		userIdInt, postParams.PreTele, postParams.Telephone, postParams.City,
		postParams.RemindTime, postParams.OrderDetail)
	if err != nil {
		handler.SendResponse(context, err, nil)
		return
	}
	handler.SendResponse(context, errnum.OK, resultData)
	return
}


/**
 * @Author Evan
 * @Description 获取用户的所有手机号订单
 * @Date 16:39 2020-11-16
 * @Param
 * @return
 **/
func GetUserOrderTel(context *gin.Context) {
	// todo defer RecoverError
	userId := context.GetHeader("X-User-Id")
	userIdInt, _ := strconv.Atoi(userId)
	resultData, err := order.GetUserOrderTel(userIdInt)
	if err != nil {
		handler.SendResponse(context, err, nil)
	}
	handler.SendResponse(context, errnum.OK, resultData)
}


/**
 * @Author Evan
 * @Description 用户删除某订单（真删）
 * @Date 18:02 2020-11-16
 * @Param
 * @return
 **/
func DeleteUserOrderTel(context *gin.Context)  {
	// post参数解析
	// todo defer RecoverError
	type PostParams struct {
		OrderId int `json:"order_id"`
	}
	postParams := &PostParams{}
	if err := context.BindJSON(&postParams); err != nil {
		err = errnum.New(errnum.ErrParsingPostJson, err)
		fmt.Println("err: " + err.Error())
		handler.SendResponse(context, err, nil)
		return
	}
	userId := context.GetHeader("X-User-Id")
	userIdInt, _ := strconv.Atoi(userId)

	resultData, err := order.DeleteUserOrderTel(postParams.OrderId, userIdInt)
	if err != nil {
		handler.SendResponse(context, err, nil)
		return
	}
	handler.SendResponse(context, errnum.OK, resultData)
}

/**
 * @Author Evan
 * @Description
 * @Date 21:41 2020-11-25
 * @Param
 * @return
 **/
func UpdateUserOrderTel(context *gin.Context) {
	// post参数解析
	//type PostParams struct {
	//	OrderId
	//	PreTele     string                       `json:"pre_tele"`
	//	Telephone   string                       `json:"telephone"`
	//	City        string                       `json:"city"`
	//	RemindTime  string                       `json:"remind_time"`
	//	OrderDetail []orderModel.OrderDetailItem `json:"order_detail"`
	//}
	postParams := &orderModel.ResOrderAndDetail{}
	if err := context.BindJSON(&postParams); err != nil {
		err = errnum.New(errnum.ErrParsingPostJson, err)
		log.Print("err: " + err.Error())
		log.Println("err: " + err.Error())
		handler.SendResponse(context, err, nil)
		return
	}
	log.Println("postParmas", postParams)
	userId := context.GetHeader("X-User-Id")
	userIdInt, _ := strconv.Atoi(userId)
	resultData, err := order.UpdateUserOrderTel(*postParams, userIdInt)
	if err != nil {
		handler.SendResponse(context, err, nil)
		return
	}
	handler.SendResponse(context, errnum.OK, resultData)
}