package user

import (
	"errors"
	"fmt"
	"wumingtianqi/model/common"
	"wumingtianqi/model/user"
	"wumingtianqi/utils"
	"wumingtianqi/utils/errnum"
)

/**
 * @Author Evan
 * @Description 邀请码奖励获取接口
	step1: 邀请码有效性检查：该邀请码是否正在被锁定（并发控制），剩余次数检查；邀请码权益解析
	step2: 查询用户信息，若该用户当前已经是VIP，但非该邀请码对应VIP等级，则报错
	step3: 若该用户当前非VIP，或是该邀请码对应VIP等级，则可用: 给用户添加相应的权益，邀请码可用次数减一
	step4: 事务：锁该邀请码（暂时用update where的方法做），以上表执行事务（这里是重点，未来做好并发控制） // todo 参考事务开启方式？并二次验证事务成功与否
 * @Date 14:36 2020-10-04
 * @Param 
 * @return
 **/
func GetInvitationReward(invitationCode string) error {
	// step1
	invitationModel := &user.Invitation{}
	invitationModel, has, err := invitationModel.QueryByInvitationCode(invitationCode)
	if err != nil{
		println("QueryByInvitationCode error")
		return err
	} else if !has {
		println("QueryByInvitationCode not found")
		return errors.New("QueryByInvitationCode not found")
	}
	invitationVipLevel := invitationModel.Vip
	duration := invitationModel.Duration
	timesRemaining := invitationModel.TimesRemaining
	coin := invitationModel.Coin
	diamond := invitationModel.Diamond

	if timesRemaining <= 0 {
		err = errnum.New(errnum.RemainingNotEnough, nil)
		return err
	}

	// step2
	userId := 1 // 未来从参数中拿(中间件做了之后)

	// step2
	userInfoFlexibleModel := &user.UserInfoFlexible{}
	userInfoFlexibleModel, has, err = userInfoFlexibleModel.QueryByUserId(userId)
	if err != nil {
		println("get user_info_flexible model error: ", err.Error())
		return err
	} else if !has {
		println("user_info_flexible model not exist")
		return errors.New("user_info_flexible model not exist")
	} else {
		println("model: ", userInfoFlexibleModel)
	}
	userVipLevel := userInfoFlexibleModel.VipLevel
	if userVipLevel >= utils.VIP1 && userVipLevel != invitationVipLevel {
		err = errnum.New(errnum.UserAlreadyVip, nil)
		return err
	}

	// step3
	localDateInt := utils.GetSpecificDate8Int(0)
	if userInfoFlexibleModel.VipLevel == utils.VIP0 {  // 不是VIP
		userInfoFlexibleModel.VipLevel = invitationVipLevel
		userInfoFlexibleModel.ExpirationTime = localDateInt
	}
	userInfoFlexibleModel.Coin += coin
	userInfoFlexibleModel.Diamond += diamond
	// 过期时间顺延
	userInfoFlexibleModel.ExpirationTime = utils.AddSpecificDays8Int(userInfoFlexibleModel.ExpirationTime, duration)
	userInfoFlexibleModel.InvitationCode = invitationCode
	invitationModel.TimesRemaining -= 1

	// step4 事务：锁该邀请码（暂时用update where的方法做），以上表执行事务（这里是重点，未来做好并发控制）
	// todo 未来压测这里
	session := common.Engine.NewSession()
	defer session.Close()
	if session.Begin() != nil {  // 事务开启
		err = errnum.New(errnum.DbError, nil)
		return err
	}
	// 更新userInfoFlexible
	_, err = session.AllCols().Where("invitation_code=?", invitationCode).Update(*invitationModel)
	if err != nil {
		fmt.Println(err.Error())
		err = errnum.New(errnum.DbError, err)
		_ = session.Rollback()
		return err
	}
	_, err = session.Where("user_id=?", userId).Update(userInfoFlexibleModel)
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		_ = session.Rollback()
		return err
	}
	err = session.Commit()
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		_ = session.Rollback()
		return err
	}
	return nil
}