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
 * @Description 邀请码奖励获取接口
	step1: 邀请码有效性检查：该邀请码是否正在被锁定（并发控制），剩余次数检查
	step2: 若该用户当前已经是VIP，但非该邀请码对应VIP等级，则报错
	step3: 若该用户当前非VIP，或是该邀请码对应VIP等级，则可用
	step4: 事务：锁该邀请码，给用户重置，同时该邀请码可用次数减一（这里是重点，未来做好并发控制）
 * @Date 14:27 2020-10-04
 * @Param context *gin.Context
 * @return
 **/
func GetInvitationReward(context *gin.Context) {
	// todo defer RecoverError
	type InvitedInfo struct {
		InvitationCode string `json:"invitation_code"`
	}
	iInfo := &InvitedInfo{}
	if err := context.BindJSON(&iInfo); err != nil {
		err = errnum.New(errnum.ErrParsingPostJson, err)
		handler.SendResponse(context, err, nil)
		return
	}
	userId := context.GetHeader("X-User-Id")
	userIdInt, _ := strconv.Atoi(userId)
	resultData, err := user.GetInvitationReward(userIdInt, iInfo.InvitationCode)
	if err != nil {
		handler.SendResponse(context, err, nil)
		return
	}
	handler.SendResponse(context, errnum.OK, resultData)
	return
}