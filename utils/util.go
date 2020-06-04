package utils

import (
	"errors"
	"reflect"
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

// 二维数组排序 参考
// https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field
func testtt() {
}