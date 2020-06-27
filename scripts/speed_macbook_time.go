package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

const ADay2Second = 10 // 希望把1天压缩到多少秒；暂未实现该功能；目前脚本模拟一天大概需要执行30秒
const OutputLog = 1 // "过"多少分钟，输出一条log

// 内部函数：获取指定时间的年月日时分，默认北京时间(10位数字的str格式)
func getLocalHourMin10Str() string {
	durationNum, _ := time.ParseDuration(strconv.Itoa(28800) + "s") // 时区偏移量（北京时间）
	localDate := time.Now().UTC().Add(durationNum)
	localDateStr := localDate.Format("0102150406")  // Mon Jan 2 15:04:05 -0700 MST 2006
	return localDateStr
}

func getSpecifyDurationHourMin(currentTime time.Time, duration time.Duration) string {
	durationNum, _ := time.ParseDuration(strconv.Itoa(28800) + "s") // 时区偏移量
	localDate := currentTime.Add(durationNum).Add(duration)  // 北京时间加上指定偏移时间
	//localDate = time.Now().UTC().Add(durationNum).Add(duration)  // 北京时间加上指定偏移时间
	localDateStr := localDate.Format("0102150406")
	return localDateStr
}

// 测试时，用于加快时间速度
// sudo go run scripts/speed_macbook_time.go
func main() {
	os := runtime.GOOS
	if os != "darwin" {
		fmt.Printf("该脚本只在mac系统下运行")
	}
	// todo 开发linux的测试脚本

	currentTime := getLocalHourMin10Str()
	command := `date ` + currentTime // sudo：先在执行该脚本的地方随便运行一个sudo命令，然后再执行该脚本
	cmd := exec.Command("/bin/bash", "-c", command)
	_, _ = cmd.Output()

	toChangedTime := time.Now().UTC()
	for i := 0; i < 1440; i++ {
		nowAfter1Min := getSpecifyDurationHourMin(toChangedTime, 1 * time.Minute)
		toChangedTime = toChangedTime.Add(1 * time.Minute)
		command = `date ` + nowAfter1Min
		cmd := exec.Command("/bin/bash", "-c", command)

		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("请用sudo执行该脚本")
			fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			return
		}
		if i % OutputLog == 0 {
			fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
			fmt.Printf("现在过了%d分钟\n", i + 1)
		}
		time.Sleep(1000 * time.Millisecond)  // 减速器：1s 充当1分钟
	}
}
