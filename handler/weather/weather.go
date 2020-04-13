package weather

import (
	"encoding/json"
	"fmt"
)

// todo 以后有路由的函数，做好路由handler函数和逻辑handler函数的解耦
func CityWeatherDailyGet(pinYin string) error {
	/*

	 */
	var jsonData interface{}

	data := `
{"results":[{"location":{"id":"WX4FBXXFKE4F","name":"北京","country":"CN","path":"北京,北京,中国","timezone":"Asia/Shanghai","timezone_offset":"+08:00"},"daily":[{"date":"2020-04-01","text_day":"多云","code_day":"4","text_night":"晴","code_night":"1","high":"15","low":"3","rainfall":"","precip":"","wind_direction":"西南","wind_direction_degree":"225","wind_speed":"10","wind_scale":"2","humidity":"0"},{"date":"2020-04-02","text_day":"晴","code_day":"0","text_night":"晴","code_night":"1","high":"20","low":"7","rainfall":"0.0","precip":"","wind_direction":"西南","wind_direction_degree":"225","wind_speed":"25.20","wind_scale":"4","humidity":"17"},{"date":"2020-04-03","text_day":"晴","code_day":"0","text_night":"晴","code_night":"1","high":"24","low":"7","rainfall":"0.0","precip":"","wind_direction":"西","wind_direction_degree":"270","wind_speed":"25.20","wind_scale":"4","humidity":"30"},{"date":"2020-04-04","text_day":"晴","code_day":"0","text_night":"晴","code_night":"1","high":"17","low":"4","rainfall":"0.0","precip":"","wind_direction":"南","wind_direction_degree":"194","wind_speed":"16.20","wind_scale":"3","humidity":"24"},{"date":"2020-04-05","text_day":"晴","code_day":"0","text_night":"晴","code_night":"1","high":"20","low":"6","rainfall":"0.0","precip":"","wind_direction":"西南","wind_direction_degree":"225","wind_speed":"16.20","wind_scale":"3","humidity":"25"}],"last_update":"2020-04-02T17:23:46+08:00"}]}
`
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		return err
	}

	fmt.Println(jsonData)

	// todo主要看这些数据
	/*
	{
		"date":"2020-04-01",
		"text_day":"多云",
		"code_day":"4",
		"text_night":"晴",
		"code_night":"1",
		"high":"15",
		"low":"3",
		"wind_direction":"西南",
		"wind_scale":"2",
	}
	 */
	// todo 封装调用api的函数
	// todo 解析心知天气
	// todo 存入天气预报表

	return err

}
