package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wumingtianqi/handler"
	"wumingtianqi/model/user"
	"wumingtianqi/utils/errnum"
)

/**
 * @Author Evan
 * @Description Token is a middleware function that convert token to user_id
 * @Date 20:59 2020-10-05
 * @Param context *gin.Context
 * @return
 **/
func TokenParsing(context *gin.Context) {
	token := context.GetHeader("X-WuMing-Token")
	userInfoModel := &user.UserInfo{}
	if userInfoModel1, has, err := userInfoModel.QueryByUserToken(token); err != nil {
		err := errnum.New(errnum.DbError, err)
		handler.SendResponse(context, err, nil)
	} else if !has {
		err = errnum.New(errnum.ErrTokenNotExist, nil)
		// todo log
		handler.SendResponse(context, err, nil)
	} else {
		context.Request.Header.Add("X-User-Id", strconv.Itoa(userInfoModel1.Id))
	}
	context.Next()
	return
}