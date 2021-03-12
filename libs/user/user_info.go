package user

import (
	"errors"
	"log"
	"strconv"
	"time"
	"wumingtianqi/model/user"
	"wumingtianqi/model/vip"
	"wumingtianqi/utils"
	"wumingtianqi/utils/errnum"
)

/**
 * @Author Evan
 * @Description 校验用户vip等级是否过期，若过期，则将vip置为0，用户权益全部置为vip0的权益
	tips: 定时任务提醒那边也用到这个
 * @Date 21:28 2021-03-04
 * @Param 
 * @return 
 **/
func CheckVipExpiration(userInfoFlexibleModel *user.UserInfoFlexible) (*user.UserInfoFlexible, error){
	// vip等级过期处理
	currentTime := time.Now()
	expirationTime := userInfoFlexibleModel.ExpirationTime
	currentTime8Int, _ := strconv.Atoi(currentTime.Format("20060102"))
	if currentTime8Int > expirationTime {  // 过期
		if userInfoFlexibleModel.VipLevel >= utils.VIP1 {
			// vip0等级映射表信息  todo redis缓存做
			vipRightsMap0 := &vip.VipRightsMap{}
			vipRightsMap0, has, err := vipRightsMap0.QueryByVipLevel(0)
			if err != nil {
				err = errnum.New(errnum.DbError, err)
				log.Println("QueryByVipLevel err:", err.Error())
				return nil, err
			} else if !has {
				log.Println("QueryByVipLevel not found")
				return nil, errors.New("QueryByVipLevel not found")
			}
			// 原vip等级映射表信息  todo redis缓存做
			vipRightsMapOld := &vip.VipRightsMap{}
			vipRightsMapOld, has, err = vipRightsMapOld.QueryByVipLevel(userInfoFlexibleModel.VipLevel)
			if err != nil {
				err = errnum.New(errnum.DbError, err)
				log.Println("QueryByVipLevel err:", err.Error())
				return nil, err
			} else if !has {
				log.Println("QueryByVipLevel not found")
				return nil, errors.New("QueryByVipLevel not found")
			}
			// 现在剩余提醒次数更新
			userInfoFlexibleModel.TelOrderRemaining = userInfoFlexibleModel.TelOrderRemaining - (vipRightsMapOld.TelOrderMax - vipRightsMap0.TelOrderMax)
			userInfoFlexibleModel.WechatOrderRemaining = userInfoFlexibleModel.WechatOrderRemaining - (vipRightsMapOld.WechatOrderMax - vipRightsMap0.WechatOrderMax)
			userInfoFlexibleModel.TodayTelRemindRemaining = userInfoFlexibleModel.TodayTelRemindRemaining - (vipRightsMapOld.TelOrderMax - vipRightsMap0.TelOrderMax)
			userInfoFlexibleModel.VipLevel = utils.VIP0

			// 提交更新
			if err := userInfoFlexibleModel.Update(); err != nil {
				err = errnum.New(errnum.DbError, err)
				return nil, err
			}
		}
	}
	return userInfoFlexibleModel, nil
}

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
		log.Println("err: ", err.Error())
		return nil, err
	} else if !has {
		log.Println("userInfoFlexibleModel not found")
		return nil, errors.New("userInfoFlexibleModel not found")
	}
	userInfoFlexibleModel, err = CheckVipExpiration(userInfoFlexibleModel)
	if err != nil {
		log.Println("err: ", err.Error())
		return nil, err
	}
	// todo 格式化日期，返回，前端展示？调用接口，展示页面，再放邀请码链接，输入邀请码获取更多权益？？
	resultData := map[string]interface{}{
		"vip_level":              userInfoFlexibleModel.VipLevel,
		"wechat_order_remaining": userInfoFlexibleModel.WechatOrderRemaining,
		"tel_order_remaining":    userInfoFlexibleModel.TelOrderRemaining,
		"expiration_time":        userInfoFlexibleModel.ExpirationTime, // 如果是vip0，那前端就隐掉该字段
	}
	return resultData, nil
}