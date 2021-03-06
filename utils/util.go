package utils

import (
	"errors"
	//"github.com/robfig/cron"
	"reflect"
	"strconv"
	"time"
)

// 获取指定日期，并拼接成8位数字的str格式，例如20200606，默认今天
func GetSpecificDate8Str(offsetDay int) string {
	durationNum, _ := time.ParseDuration(strconv.Itoa(28800 + offsetDay * 3600 * 24) + "s") // 时区偏移量（北京时间）
	localDate := time.Now().UTC().Add(durationNum)
	localDateStr := localDate.Format("20060102")
	return localDateStr
}

// 获取指定日期，并拼接成8位数字的int格式，例如20200606，默认今天
func GetSpecificDate8Int(offsetDay int) int {
	localDateStr := GetSpecificDate8Str(offsetDay)
	localDateInt, _ := strconv.Atoi(localDateStr)
	return localDateInt
}

// 指定日期增加/减少n天，传入参数和返回参数都是8位int格式
func AddSpecificDays8Int(dateToOperateInt int, offsetDay int) int {
	// 数字转日期
	dateToOperate := Int8ToDate(dateToOperateInt)  // 20201005

	// 日期加指定天数，然后转成8位string
	durationNum, _ := time.ParseDuration(strconv.Itoa(28800 + offsetDay * 3600 * 24) + "s") // 时区偏移量（北京时间）
	targetDate := dateToOperate.Add(durationNum)
	targetDateStr8 := targetDate.Format("20060102")
	targetDateInt8, _ := strconv.Atoi(targetDateStr8)
	return targetDateInt8
}

// 数字转日期
func Int8ToDate(dateInt8 int) time.Time {
	dateStr8 := strconv.Itoa(dateInt8)
	parseTimeRes, _ := time.ParseInLocation("20060102" ,dateStr8, time.Local)
	return parseTimeRes
}

// 获取当前北京时间的时分(4位数字的str格式)
func GetLocalHourMin4Str() string {
	durationNum, _ := time.ParseDuration(strconv.Itoa(28800) + "s") // 时区偏移量（北京时间）
	localDate := time.Now().UTC().Add(durationNum)
	localDateStr := localDate.Format("1504")
	return localDateStr
}

// 获取指定时间的时分(4位数字的str格式)
func GetSpecifyDurationHourMin(duration time.Duration) string {
	localDate := time.Now().UTC().Add(duration)    // 加上指定偏移时间
	localDateStr := localDate.Format("1504")
	return localDateStr
}

func IsContain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target) // todo test
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("obj is not in target")
}

//func NewWithSeconds() *cron.Cron {
//	secondParser := cron.NewParser(cron.Second | cron.Minute |
//		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
//	return cron.New( cron.WithParser(secondParser), cron.WithChain())
//}
