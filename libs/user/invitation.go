package user

import (
	"errors"
	"fmt"
	"log"
	"time"
	"wumingtianqi/model/common"
	"wumingtianqi/model/user"
	"wumingtianqi/model/vip"
	"wumingtianqi/utils"
	"wumingtianqi/utils/errnum"
)

/**
 * @Author Evan
 * @Description 邀请码奖励获取接口
	step1: 邀请码有效性检查：该邀请码是否正在被锁定（todo 并发控制），剩余次数检查，邀请码权益解析；用户是否用过该邀请码；
	step2: 查询用户信息，若该用户当前已经是VIP，但非该邀请码对应VIP等级，则覆盖旧vip
	step3: 若该用户当前非VIP，或是该邀请码对应VIP等级，则可用: 给用户添加相应的权益，邀请码可用次数减一
	step4: 事务：锁该邀请码（暂时用update where的方法做），以上表执行事务（这里是重点，未来做好并发控制）
 * @Date 14:36 2020-10-04
 * @Param 
 * @return
 **/
func GetInvitationReward(userId int, invitationCode string) (map[string]interface{}, error) {
	// step1
	invitationModel := &user.Invitation{}
	invitationModel, has, err := invitationModel.QueryByInvitationCode(invitationCode)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("QueryByInvitationCode  err: " + err.Error())
		return nil, err
	} else if !has {
		log.Println("QueryByInvitationCode not found")
		return nil, errors.New("QueryByInvitationCode not found")
	}
	invitationVipLevel := invitationModel.Vip
	duration := invitationModel.Duration
	timesRemaining := invitationModel.TimesRemaining
	log.Println("timesRemaining", timesRemaining)
	coin := invitationModel.Coin
	diamond := invitationModel.Diamond

	if timesRemaining <= 0 {  // 邀请码的次数已被用完
		err = errnum.New(errnum.RemainingNotEnough, errors.New("RemainingNotEnough"))
		return nil, err
	}

	// 用户是否用过该邀请码
	userInvitationMap := user.UserInvitationMap{}
	_, has, err = userInvitationMap.QueryByUserInvitation(userId, invitationCode)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("QueryByUserInvitation  err: " + err.Error())
		return nil, err
	}
	if has {
		err = errnum.New(errnum.ErrUserUsedInvitationCode, err)
		return nil, err
	}

	// step2
	userInfoFlexibleModel := &user.UserInfoFlexible{}
	userInfoFlexibleModel, has, err = userInfoFlexibleModel.QueryByUserId(userId)
	if err != nil {
		println("get user_info_flexible model error: ", err.Error())
		return nil, err
	} else if !has {
		println("user_info_flexible model not exist")
		return nil, errors.New("user_info_flexible model not exist")
	} else {
		println("model: ", userInfoFlexibleModel)
	}
	localDateInt := utils.GetSpecificDate8Int(0)
	userVipLevel := userInfoFlexibleModel.VipLevel

	if userVipLevel >= utils.VIP1 && userVipLevel != invitationVipLevel {  // 用户当前已经是VIP，但非该邀请码对应VIP等级，则覆盖vip
		// 之前是报错；现在是做成与非VIP逻辑一样；todo 未来支持多vip模式，采用顺延旧vip的机制
		//err = errnum.New(errnum.UserAlreadyVip, nil)
		//return nil, err
		userInfoFlexibleModel.VipLevel = invitationVipLevel
		userInfoFlexibleModel.ExpirationTime = localDateInt
	}

	// step3
	if userInfoFlexibleModel.VipLevel == utils.VIP0 {  // 不是VIP
		userInfoFlexibleModel.VipLevel = invitationVipLevel
		userInfoFlexibleModel.ExpirationTime = localDateInt
	}
	userInfoFlexibleModel.Coin += coin
	userInfoFlexibleModel.Diamond += diamond
	// 过期时间顺延
	userInfoFlexibleModel.ExpirationTime = utils.AddSpecificDays8Int(userInfoFlexibleModel.ExpirationTime, duration)
	userInfoFlexibleModel.InvitationCode = invitationCode
	vipRightsMap := &vip.VipRightsMap{}
	vipRightsMap, has, err = vipRightsMap.QueryByVipLevel(invitationVipLevel)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("QueryByVipLevel err:", err.Error())
		return nil, err
	} else if !has {
		log.Println("QueryByVipLevel not found")
		return nil, errors.New("QueryByVipLevel not found")
	}
	// 解析出vip表的权益，赋予用户
	wechatOrderMax := vipRightsMap.WechatOrderMax
	telOrderMax := vipRightsMap.TelOrderMax
	_ = vipRightsMap.RemindPatternIdList  // 暂时前端写死
	// 赋予用户
	userInfoFlexibleModel.WechatOrderRemaining = wechatOrderMax
	userInfoFlexibleModel.TelOrderRemaining = telOrderMax
	currentTime := time.Now()
	userInfoFlexibleModel.UpdateTime = currentTime
	invitationModel.TimesRemaining -= 1
	invitationModel.UpdateTime = currentTime

	// 映射表记录
	userInvitationMap.UserId = userId
	userInvitationMap.InvitationCode = invitationCode
	userInvitationMap.CreateTime = currentTime

	// step4 事务：锁该邀请码（暂时用update where的方法做），以上表执行事务（这里是重点，未来做好并发控制）
	// todo 未来压测这里
	session := common.Engine.NewSession()
	defer session.Close()
	if session.Begin() != nil {  // 事务开启
		err = errnum.New(errnum.DbError, nil)
		return nil, err
	}
	// 4.1更新Invitation model
	rowsAffected, err := session.AllCols().Where(
		"invitation_code=?", invitationCode).And(
			"times_remaining>=1").Update(*invitationModel)
	log.Println("rowsAffected", rowsAffected)
	if rowsAffected <= 0 {
		err = errnum.New(errnum.RemainingNotEnough, errors.New("concurrent error"))
		log.Println("rowsAffected", err.Error())
		return nil, err
	}
	if err != nil {
		fmt.Println(err.Error())
		err = errnum.New(errnum.DbError, err)
		return nil, err
	}
	// 4.2更新userInfoFlexible
	_, err = session.Where("user_id=?", userId).Update(userInfoFlexibleModel)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		return nil, err
	}

	// 4.3新增UserInvitationMap
	if _, err = session.InsertOne(userInvitationMap); err != nil {
		err = errnum.New(errnum.DbError, err)
		return nil, err
	}

	// 4.4 提交事务
	err = session.Commit()
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		_ = session.Rollback()
		return nil, err
	}
	resultData := map[string]interface{}{
		"result": "success",
	}
	return resultData, nil
}