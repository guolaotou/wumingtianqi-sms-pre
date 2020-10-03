package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
	"wumingtianqi/config"
	"wumingtianqi/model"
	"wumingtianqi/model/user"
	"wumingtianqi/utils"
)

/**
 * @Author Evan
 * @Description 生成邀请码
	step1: 生成uuid
	step2：该uuid随机取32位（可重复取）进行md5加密
	step3: uuid随机取16位（不重复），md5加密后的字符串随机取16位（不重复），然后拼接
	注：(uuid参考 https://github.com/satori/go.uuid)
 * @Date 12:57 2020-10-03
 * @Param
 * @return string 最终生成的邀请码
 **/
func generateInvitationCode() string {
	// step1: 生成uuid
	u1 := uuid.NewV4()
	strUUID36 := u1.String()
	//fmt.Printf("UUIDv4: %s\n", strUUID36)

	// step2：该uuid随机取32位（可重复取）进行md5加密
	strUUID32 := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		index := r.Intn(36)  // uuid长度
		strUUID32 += string(strUUID36[index])
	}
	key := config.GlobalConfig.Util.Uuid2MD5Key  // 加盐
	strMd532, _ := MD5(strUUID32, key)

	// step3: uuid随机取16位（不重复），md5加密后的字符串随机取16位（不重复），然后拼接
	uuidIndexList := randomNM(35, 16, r)
	md5IndexList := randomNM(31, 16, r)
	strUUID16 := ""
	strMd516 := ""
	for i := 0; i < 16; i++ {
		strUUID16 += string(strUUID36[uuidIndexList[i]])
		strMd516 += string(strMd532[md5IndexList[i]])
	}
	// 预拼接，然后从0~31之间选一个位置做头
	strTemp64 := strUUID16 + strMd516 + strUUID16 + strMd516
	index := r.Intn(32)

	// 拼接
	strInvitation32 := strTemp64[index:index+32]
	//fmt.Println("strInvitation32", strInvitation32)

	return strInvitation32
}

/**
 * @Author Evan
 * @Description 给邀请码分配权限——一等邀请码
	（mine生成10个）VIP3，有效期99年，初始10000个钻石
 * @Date 13:17 2020-10-03
 * @Param
 * @return
 **/
func SetInvitationCodeAuthLevel1(invitationCode string) *user.Invitation {
	invitationModel := new(user.Invitation)
	currentTime := time.Now()
	invitationModel.InvitationCode = invitationCode
	invitationModel.TimesMax  = 1
	invitationModel.TimesRemaining  = 1
	invitationModel.Vip = utils.VIP3
	invitationModel.Duration = 366 * 99
	invitationModel.Coin = 0
	invitationModel.Diamond = 10000
	invitationModel.Creator = -1
	invitationModel.CreateTime = currentTime
	invitationModel.UpdateTime = currentTime
	return invitationModel
}

/**
 * @Author Evan
 * @Description 给邀请码分配权限——二等邀请码
	mine生成5个，每个可以给20个人用） 默认VIP2，有效期99年，默认100钻石
 * @Date 17:18 2020-10-03
 * @Param
 * @return
 **/
func SetInvitationCodeAuthLevel2(invitationCode string) *user.Invitation {
	invitationModel := new(user.Invitation)
	currentTime := time.Now()
	invitationModel.InvitationCode = invitationCode
	invitationModel.TimesMax = 20
	invitationModel.TimesRemaining = 20
	invitationModel.Vip = utils.VIP2
	invitationModel.Duration = 366 * 99
	invitationModel.Coin = 0
	invitationModel.Diamond = 100
	invitationModel.Creator = -1
	invitationModel.CreateTime = currentTime
	invitationModel.UpdateTime = currentTime
	return invitationModel
}
/**
 * @Author Evan
 * @Description 摇色子的工具
	给定最大值m，再给指定位数n，随机返回m个取值范围在0 ~ n以内的不重复的整数
	例如，给定m n 分别为 100, 5
		 返回 [3, 100, 2, 4, 0]
 * @Date 20:41 2020-10-02
 * @Param
	m int 最大值
	n int 位数
	rRand ...*rand.Rand 随机种子，如果没有就自己生成
 * @return
 **/
func randomNM(m int, n int, rRand ...*rand.Rand) []int {
	resMap := map[int]bool{}  // 摇色子出来的数就放到map里，遇到碰撞二次尝试。再碰撞，就不断直接加1找到不碰撞的值
	resList := make([]int, 0, len(resMap))
	if m < n - 1 {
		return []int{}
	} else if m == 0 {
		resList = append(resList, 0)
		return resList
	}

	// 获取随机种子
	var r *rand.Rand
	if rRand != nil {
		r = rRand[0]
	} else {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	// 算法开始
	for i := 0; i < n; i++ {
		temp := r.Intn(m)
		if _, ok := resMap[temp]; ok {
			// 存在，重试一次
			temp = r.Intn(m)
			if _, ok := resMap[temp]; ok {
				// 仍存在，则循环至多n-1次，每次在原有基础上加1，最后temp一定是不碰撞的值
				for j := 1; j < n; j++ {
					temp = (temp + j) % (m + 1)
					if _, ok := resMap[temp]; !ok {
						break
					}
				}
			}
		}
		resMap[temp] = true
	}

	// 遍历map的所有key，放到list里
	for k := range resMap {  // 参考：https://blog.csdn.net/yzf279533105/article/details/94010954
		resList = append(resList, k)
	}
	return resList
}

/**
 * @Author Evan
 * @Description 给字符拼接上秘钥串，然后MD5加密
 * @Date 20:36 2020-10-02
 * @Param
	originStr string 原始字符串
	key string 秘钥串
 * @return (string, error)
 **/
func MD5(originStr string, key string) (string, error) {
	// 判断key是否正确
	if len(key) <= 0 {
		return "", errors.New("encrypt need key")
	}
	text := originStr + key
	ctx := md5.New()
	ctx.Write([]byte(text))
	res := hex.EncodeToString(ctx.Sum(nil))
	return res, nil
}

/**
 * @Author Evan
 * @Description 全流程：生成10个一等邀请码，5个二等邀请码;为该邀请码配置权限，并把该邀请码存入数据库
 * @Date 17:20 2020-10-03
 * @Param
 * @return
 **/
func ProcessAll() {
	model.InitMysql()
	//session := common.Engine.NewSession()
	//defer session.Close()
	// 1. 生成10个一等邀请码，分别为邀请码配置权限，并把该邀请码存入数据库
	for i := 0; i < 10; i++ {
		invitationCode := generateInvitationCode()
		invitationModel := SetInvitationCodeAuthLevel1(invitationCode)
		if err := invitationModel.Create(); err != nil {
			println("Creating level 1")
			panic(err)
		}
	}

	// 2. 生成5个二等邀请码，分别为邀请码配置权限，并把该邀请码存入数据库
	for i := 0; i < 5; i++ {
		invitationCode := generateInvitationCode()
		invitationModel := SetInvitationCodeAuthLevel2(invitationCode)
		if err := invitationModel.Create(); err != nil {
			println("Creating level 2")
			panic(err)
		}
	}
	// alter table invitation AUTO_INCREMENT=1;
}




// go run scripts/invitation/invitation.go
func main() {
	/*
	邀请机制相关脚本
	 */
	if _, err := config.LoadConfig(); err != nil {
		fmt.Println(err.Error())
	}

	// 生成邀请码
	strInvitation32 := generateInvitationCode()
	println("strInvitation32", strInvitation32)

	// 全流程：生成10个一等邀请码，5个二等邀请码;为该邀请码配置权限，并把该邀请码存入数据库
	ProcessAll()
}