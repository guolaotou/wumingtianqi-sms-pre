package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"reflect"
	"time"
	"wumingtianqi/config"
)

func generateInvitationCode() string {
	/* 生成邀请码
	step1: 生成uuid (uuid参考 https://github.com/satori/go.uuid)
	step2: 用
	 */

	// step1: 生成uuid
	u1 := uuid.NewV4()
	u1Str := u1.String()
	fmt.Printf("UUIDv4: %s\n", u1Str)

	// step2：该uuid随机取32位（可重复取）进行md5加密
	strUUID32 := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	println("duandian1", u1Str)
	for i := 0; i < 32; i++ {
		index := r.Intn(36)
		strUUID32 += string(u1Str[index])
	}
	key := config.GlobalConfig.Util.Uuid2MD5Key
	println("key", key)
	md5Res, _ := MD5(strUUID32, key)
	println("md5", md5Res)
	println("r", r)

	fmt.Println("r", reflect.TypeOf(r))
	resList := randomNM(0, 1, r)

	for i := 0; i < len(resList); i++ {
		println("i: ", resList[i])
	}

	// step3: uuid随机取16位，md5加密后的字符串随机取16位，然后拼接

	// step4: 为该邀请码配置权限，并把该邀请码存入数据库

	return u1Str
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
	println("resdfgh", res)
	return res, nil
}


// go run scripts/invitation/invitation.go
func main() {
	/*
	邀请机制相关脚本
	 */
	if _, err := config.LoadConfig(); err != nil {
		fmt.Println(err.Error())
	}
	generateInvitationCode()


}