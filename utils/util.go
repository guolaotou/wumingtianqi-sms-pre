package utils

import (
	"strconv"
	"time"
)

// 获取指定时间的时分，默认北京时间(4位数字的str格式)
func GetLocalHourMin4Str() string {
	durationNum, _ := time.ParseDuration(strconv.Itoa(28800) + "s") // 时区偏移量（北京时间）
	localDate := time.Now().UTC().Add(durationNum)
	localDateStr := localDate.Format("1504")
	return localDateStr
}

func GetSpecifyDurationHourMin(duration time.Duration) string {
	durationNum, _ := time.ParseDuration(strconv.Itoa(28800) + "s") // 时区偏移量
	localDate := time.Now().UTC().Add(durationNum).Add(duration)    // 北京时间加上指定偏移时间
	localDateStr := localDate.Format("1504")
	return localDateStr
}
