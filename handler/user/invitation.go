package user

import "github.com/gin-gonic/gin"

/**
 * @Author Evan
 * @Description 邀请码奖励获取接口
	step1: 邀请码有效性检查：该邀请码是否正在被锁定（并发控制），剩余次数检查
	step2: 若该用户当前已经是VIP，但非该邀请码对应VIP等级，则报错
	step3: 若该用户当前非VIP，或是该邀请码对应VIP等级，则可用
	step4: 事务：锁该邀请码，给用户重置，同时该邀请码可用次数减一（这里是重点，未来组好并发控制）
 * @Date 14:27 2020-10-04
 * @Param context *gin.Context
 * @return
 **/
// todo 中间件，验证用户身份
func GetInvitationReward(context *gin.Context){
	// mine 先写最后一步，给用户发放邀请码对应权益，在lib层写
	// 暂时用body传参，后面再弄中间件
	// todo 解析post参数
	//todo 从session_key 获取user_id
	// todo 写view层test函数
	// 放到header里，命名为token
	// 那么，就完成邀请机制，可以commit

	// 然后手机号配置接口，可以严重用户机制。只需要配置最简单的一种
}