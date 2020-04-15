package weather

import (
	"fmt"
	"github.com/intel-go/fastjson"
	"io/ioutil"
	"net/http"
	"wumingtianqi-sms-pre/config"
)

// 获取新知天气信息 （存入数据库是在另一个函数做）
func Get2XinZhiWhether(fakeData bool, cityPinYin string) (XinZhiWeatherDailyResults, error) {  // fakeData为true就不用真的去调用数据了，直接用假数据
	if fakeData {
		resBody := `{
    "results":[
        {
            "location":{
                "id":"WX4FBXXFKE4F",
                "name":"北京",
                "country":"CN",
                "path":"北京,北京,中国",
                "timezone":"Asia/Shanghai",
                "timezone_offset":"+08:00"
            },
            "daily":[
                {
                    "date":"2020-04-14",
                    "text_day":"多云",
                    "code_day":"4",
                    "text_night":"多云",
                    "code_night":"4",
                    "high":"26",
                    "low":"11",
                    "rainfall":"",
                    "precip":"",
                    "wind_direction":"南",
                    "wind_direction_degree":"180",
                    "wind_speed":"10",
                    "wind_scale":"2",
                    "humidity":"0"
                },
                {
                    "date":"2020-04-15",
                    "text_day":"多云",
                    "code_day":"4",
                    "text_night":"中雨",
                    "code_night":"14",
                    "high":"26",
                    "low":"11",
                    "rainfall":"0.0",
                    "precip":"",
                    "wind_direction":"南",
                    "wind_direction_degree":"180",
                    "wind_speed":"16.20",
                    "wind_scale":"3",
                    "humidity":"48"
                },
                {
                    "date":"2020-04-16",
                    "text_day":"多云",
                    "code_day":"4",
                    "text_night":"多云",
                    "code_night":"4",
                    "high":"22",
                    "low":"11",
                    "rainfall":"0.0",
                    "precip":"",
                    "wind_direction":"西北",
                    "wind_direction_degree":"315",
                    "wind_speed":"25.20",
                    "wind_scale":"4",
                    "humidity":"53"
                },
                {
                    "date":"2020-04-17",
                    "text_day":"多云",
                    "code_day":"4",
                    "text_night":"晴",
                    "code_night":"1",
                    "high":"22",
                    "low":"8",
                    "rainfall":"0.0",
                    "precip":"",
                    "wind_direction":"北",
                    "wind_direction_degree":"356",
                    "wind_speed":"34.20",
                    "wind_scale":"5",
                    "humidity":"43"
                },
                {
                    "date":"2020-04-18",
                    "text_day":"多云",
                    "code_day":"4",
                    "text_night":"小雨",
                    "code_night":"13",
                    "high":"22",
                    "low":"12",
                    "rainfall":"0.0",
                    "precip":"",
                    "wind_direction":"南",
                    "wind_direction_degree":"201",
                    "wind_speed":"16.20",
                    "wind_scale":"3",
                    "humidity":"48"
                }
            ],
            "last_update":"2020-04-15T11:17:52+08:00"
        }
    ]
}`
		var res2return XinZhiWeatherDailyResults
		err := fastjson.Unmarshal([]byte(resBody), &res2return)
		if err != nil {
			fmt.Println(err.Error())
			return XinZhiWeatherDailyResults{}, err
		}
		fmt.Println(res2return)
		fmt.Println(res2return.Results[0].Daily[0])
		return res2return, nil
	}
	url := config.GlobalConfig.Weather.XinZhiWeatherDailyUrl
	apiKey :=config.GlobalConfig.Weather.ApiKey

	// 拼接请求参数
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return XinZhiWeatherDailyResults{}, err
	}
	q := req.URL.Query()
	q.Add("key", apiKey)
	q.Add("location", cityPinYin)
	req.URL.RawQuery = q.Encode()
	//fmt.Println(req.URL.String())
	// https://api.seniverse.com/v3/weather/daily.json?location=beijing&language=zh-Hans&unit=c&start=-1&days=5&key=xxx

	// 请求 & 接收返回值
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			resBody, _ := ioutil.ReadAll(resp.Body)
			var res2return XinZhiWeatherDailyResults
			err = fastjson.Unmarshal(resBody, &res2return)
			if err != nil {
				fmt.Println(err.Error())
				return XinZhiWeatherDailyResults{}, err
			}
			//fmt.Println(res2return)
			//fmt.Println(res2return.Results[0].Daily[0])
			return res2return, nil
		}
	}
	return XinZhiWeatherDailyResults{}, nil
}

// todo 有路由的函数，做好路由handler函数和逻辑handler函数的解耦
func CityWeatherDailyGet(cityPinYin string) (XinZhiWeatherDailyResults, error) {
	res, err := Get2XinZhiWhether(true, cityPinYin)
	return res, err
}
