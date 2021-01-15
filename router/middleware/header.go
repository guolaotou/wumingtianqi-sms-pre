package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wumingtianqi/handler"
	"wumingtianqi/model/user"
	"wumingtianqi/utils/errnum"
)

// token有问题直接返回 context.Abort https://blog.csdn.net/minibrid/article/details/95354754
// google关键词: gin 终止中间件返回数据
/**
 * @Author Evan
 * @Description Token is a middleware function that convert token to user_id
 * @Date 20:59 2020-10-05
 * @Param context *gin.Context
 * @return
 **/
func TokenParsing(context *gin.Context) {
	token := context.GetHeader("X-WuMing-Token")
	userInfoModelInstance := &user.UserInfo{}
	if userInfoModel, has, err := userInfoModelInstance.QueryByUserToken(token); err != nil {
		err := errnum.New(errnum.DbError, err)
		handler.SendResponse(context, err, nil)
		context.Abort()
		return
	} else if !has {
		err = errnum.New(errnum.ErrTokenNotExist, nil)
		// todo log
		handler.SendResponse(context, err, nil)
		context.Abort()
		return
	} else {
		context.Request.Header.Add("X-User-Id", strconv.Itoa(userInfoModel.Id))
	}
	context.Next()
	return
}