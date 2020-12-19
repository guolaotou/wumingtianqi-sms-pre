package user

import (
	"errors"
	"log"
	"wumingtianqi/model/user"
	"wumingtianqi/utils/errnum"
)

/**
 * @Author Evan
 * @Description 获取用户信息
 * @Date 20:01 2020-12-19
 * @Param 
 * @return 
 **/
func GetUserInfo(userId int) (map[string]interface{}, error) {
	// 获取用户哪些表的信息？然后打印出来
	userInfoFlexibleModel := &user.UserInfoFlexible{}
	userInfoFlexibleModel, has, err := userInfoFlexibleModel.QueryByUserId(userId)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("err: ", err)
		return nil, err
	} else if !has {
		log.Println("userInfoFlexibleModel not found")
		return nil, errors.New("userInfoFlexibleModel not found")
	}
	// 解析user信息，然后打印出来
	log.Println("userInfoFlexibleModel.InvitationCode", userInfoFlexibleModel.InvitationCode)
	log.Println("userInfoFlexibleModel.VipLevel", userInfoFlexibleModel.VipLevel)
	log.Println("userInfoFlexibleModel.WechatOrderRemaining", userInfoFlexibleModel.WechatOrderRemaining)
	log.Println("userInfoFlexibleModel.TelOrderRemaining", userInfoFlexibleModel.TelOrderRemaining)
	log.Println("userInfoFlexibleModel.TodayEditChanceRemaining", userInfoFlexibleModel.TodayEditChanceRemaining)
	log.Println("userInfoFlexibleModel.Coin", userInfoFlexibleModel.Coin)
	log.Println("userInfoFlexibleModel.Diamond", userInfoFlexibleModel.Diamond)
	log.Println("userInfoFlexibleModel.ExpirationTime", userInfoFlexibleModel.ExpirationTime)
	// todo 格式化日期，返回，前端展示？调用接口，展示页面，再放邀请码链接，输入邀请码获取更多权益？？
	//前端初始化页面，vip判断，如果非vip，那就让输入邀请码？（调研论坛的页面）
	// todo 删除以上注释
	resultData := map[string]interface{}{
		"vip_level":              userInfoFlexibleModel.VipLevel,
		"wechat_order_remaining": userInfoFlexibleModel.WechatOrderRemaining,
		"tel_order_remaining":    userInfoFlexibleModel.TelOrderRemaining,
		"expiration_time":        userInfoFlexibleModel.ExpirationTime, // todo 处理
	}
	return resultData, nil
}