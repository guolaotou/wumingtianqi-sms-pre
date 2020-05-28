package main

import "fmt"

// 1. 获取5分钟后提醒提醒用户列表
// model: user, order, Q_to_remind
// todo 最大的main改？ 思考是不是挪到这里
// 还可以参考正规cron写法，vanguard代码


// 2. 拼接提醒信息
// model: user, order, Q_to_remind, Q_to_send
// lib 短信模块


// 调用remind相关的后台（定时）任务
// go run cron/remind/main.go
func main() {
	fmt.Println("duandian1")
}

