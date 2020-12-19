package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wumingtianqi/handler"
	"wumingtianqi/libs/user"
	"wumingtianqi/utils/errnum"
)

/**
 * @Author Evan
 * @Description 获取用户信息接口
 * @Date 19:58 2020-12-19
 * @Param 
 * @return 
 **/
func GetUserInfo(context *gin.Context)  {
	//todo defer RecoverError
	userId := context.GetHeader("X-User-Id")
	userIdInt, _ := strconv.Atoi(userId)
	resultData, err := user.GetUserInfo(userIdInt)
	if err != nil {
		handler.SendResponse(context, err, nil)
		return
	}
	handler.SendResponse(context, errnum.OK, resultData)
	return
}