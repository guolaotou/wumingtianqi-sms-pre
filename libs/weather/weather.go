package weather

import (
	"errors"
	"fmt"
	"github.com/intel-go/fastjson"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
	"wumingtianqi/config"
	"wumingtianqi/model/city"
	"wumingtianqi/model/common"
	"wumingtianqi/model/weather"
	"wumingtianqi/utils"
	"wumingtianqi/utils/errnum"
)

// 根据指定城市，获取新知天气信息
func Get2XinZhiWhether(fakeData bool, cityPinYin string) (XinZhiWeatherDailyResults, error) {  // fakeData为true就不用真的去调用数据了，直接用假数据
	if fakeData {
//		resBody := `{
//    "results":[
//        {
//            "location":{
//                "id":"WX4FBXXFKE4F",
//                "name":"北京",
//                "country":"CN",
//                "path":"北京,北京,中国",
//                "timezone":"Asia/Shanghai",
//                "timezone_offset":"+08:00"
//            },
//            "daily":[
//                {
//                    "date":"2020-04-14",
//                    "text_day":"多云",
//                    "code_day":"4",
//                    "text_night":"多云",
//                    "code_night":"4",
//                    "high":"26",
//                    "low":"11",
//                    "rainfall":"",
//                    "precip":"",
//                    "wind_direction":"南",
//                    "wind_direction_degree":"180",
//                    "wind_speed":"10",
//                    "wind_scale":"2",
//                    "humidity":"0"
//                },
//                {
//                    "date":"2020-04-15",
//                    "text_day":"多云",
//                    "code_day":"4",
//                    "text_night":"中雨",
//                    "code_night":"14",
//                    "high":"26",
//                    "low":"11",
//                    "rainfall":"0.0",
//                    "precip":"",
//                    "wind_direction":"南",
//                    "wind_direction_degree":"180",
//                    "wind_speed":"16.20",
//                    "wind_scale":"3",
//                    "humidity":"48"
//                },
//                {
//                    "date":"2020-04-16",
//                    "text_day":"多云",
//                    "code_day":"4",
//                    "text_night":"多云",
//                    "code_night":"4",
//                    "high":"22",
//                    "low":"11",
//                    "rainfall":"0.0",
//                    "precip":"",
//                    "wind_direction":"西北",
//                    "wind_direction_degree":"315",
//                    "wind_speed":"25.20",
//                    "wind_scale":"4",
//                    "humidity":"53"
//                },
//                {
//                    "date":"2020-04-17",
//                    "text_day":"多云",
//                    "code_day":"4",
//                    "text_night":"晴",
//                    "code_night":"1",
//                    "high":"22",
//                    "low":"8",
//                    "rainfall":"0.0",
//                    "precip":"",
//                    "wind_direction":"北",
//                    "wind_direction_degree":"356",
//                    "wind_speed":"34.20",
//                    "wind_scale":"5",
//                    "humidity":"43"
//                },
//                {
//                    "date":"2020-04-18",
//                    "text_day":"多云",
//                    "code_day":"4",
//                    "text_night":"小雨",
//                    "code_night":"13",
//                    "high":"22",
//                    "low":"12",
//                    "rainfall":"0.0",
//                    "precip":"",
//                    "wind_direction":"南",
//                    "wind_direction_degree":"201",
//                    "wind_speed":"16.20",
//                    "wind_scale":"3",
//                    "humidity":"48"
//                }
//            ],
//            "last_update":"2020-04-15T11:17:52+08:00"
//        }
//    ]
//}`
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
                    "date":"2020-06-26",
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
                    "date":"2020-06-27",
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
                    "date":"2020-06-28",
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
                    "date":"2020-06-29",
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
                    "date":"2020-06-30",
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

// todo 根据城市list，顺次异步调用天气信息（控制调用频率）；结果放在channel里
// todo 读取channel，在最后写入数据库（用sync.WaitGroup）

// todo 有路由的函数，做好路由handler函数和逻辑handler函数的解耦
func CityWeatherDailyGet(cityPinYin string) (XinZhiWeatherDailyResults, error) {
	res, err := Get2XinZhiWhether(config.GlobalConfig.Weather.FakeData, cityPinYin)
	return res, err
}

func dayWeather2Map(xinZhiWeather XinZhiWeatherDailyResults) map[int]XinZhiWeatherDailyItem{
	//fmt.Println("xinZhiWeather", xinZhiWeather.Results[0].Daily)
	//fmt.Println("xinZhiWeather", len(xinZhiWeather.Results[0].Daily))
	//fmt.Println("xinZhiWeather", xinZhiWeather)
	dayWeatherMap := make(map[int]XinZhiWeatherDailyItem, 0)

	for _, val := range xinZhiWeather.Results[0].Daily {
		date8Str := strings.Replace(val.Date, "-", "", -1)
		date8Int, _ := strconv.Atoi(date8Str)

		// 拿天气视频整理成map，每一天的信息都有
		for offset := -1; offset <= 1; offset++ {
			// 获取特定日期！
			localDateInt := utils.GetSpecificDate8Int(offset)  // 0 -> 今天
			if localDateInt == date8Int {
				dayWeatherMap[date8Int] = val
			}
		}
	}
	return dayWeatherMap
}

func UpdateWeatherDaily() {
	/*
	每日存储天气信息
	1. 获取城市列表
	2. 遍历城市列表，获取天气信息，放入channel（生产者）
	3. 不断从channel中获取天气信息，存入数据库（消费者）
	tips:
	控制生产者速率 20条 / min
	控制消费者，当存满100条，存一次
	以上为todo
	 */
	cityList, err := city.GetAllCity()
	if  err != nil {
		panic(err)
	}

	if cityList == nil {  // todo 验证这种写法
		return
	}

	for i := 0; i < len(cityList); i++ {
		// 获取城市天气信息
		pinYin := cityList[i].PinYin  // "beijing"
		println("pinYin", pinYin)
		xinZhiWeather, err := Get2XinZhiWhether(config.GlobalConfig.Weather.FakeData, pinYin)
		if err != nil {
			panic(err)
		}
		// todo 以日期做map
		fmt.Println(xinZhiWeather)
		if reflect.DeepEqual(xinZhiWeather, XinZhiWeatherDailyResults{}) {
			log.Printf(pinYin + " has no data")
			continue
		}
		dayWeatherMap := dayWeather2Map(xinZhiWeather)
		fmt.Println("dayWeatherMap", dayWeatherMap[20200414])

		// 存入数据库: 如果某天没有，那就存；否则，更新天气信息
		fmt.Println(xinZhiWeather)

		// 拿天气视频整理成map，每一天的信息都有
		for offset := -1; offset <= 1; offset++ {
			// 获取特定日期！
			localDateInt := utils.GetSpecificDate8Int(offset)  // 0 -> 今天
			fmt.Println("localDateInt", localDateInt)  // todo delete
			fmt.Println("dayWeatherMap[localDateInt]", dayWeatherMap[localDateInt])  // todo delete
			fmt.Println("dayWeatherMap[localDateInt]", XinZhiWeatherDailyItem{} == dayWeatherMap[localDateInt])  // todo delete
			if (XinZhiWeatherDailyItem{} == dayWeatherMap[localDateInt]) == true {
				fmt.Println("天气信息抓取有误" + pinYin + strconv.Itoa(localDateInt))
				// todo 报错，打tick
				continue
			}

			CodeDay, err := strconv.Atoi(dayWeatherMap[localDateInt].CodeDay)  // todo 下面代码考虑优化
			CodeNight, err := strconv.Atoi(dayWeatherMap[localDateInt].CodeNight)
			High, err := strconv.Atoi(dayWeatherMap[localDateInt].High)
			Low, err := strconv.Atoi(dayWeatherMap[localDateInt].Low)
			WindScale, err := strconv.Atoi(dayWeatherMap[localDateInt].WindScale)
			WindSpeed, err := strconv.Atoi(dayWeatherMap[localDateInt].WindSpeed)
			Humidity, err := strconv.Atoi(dayWeatherMap[localDateInt].Humidity)
			dayWeather := weather.DayWeather{
				CityPinYin:    pinYin,
				DateId:        localDateInt,
				TextDay:       dayWeatherMap[localDateInt].TextDay,
				CodeDay:       CodeDay,
				TextNight:     dayWeatherMap[localDateInt].TextNight,
				CodeNight:     CodeNight,
				High:          High,
				Low:           Low,
				WindDirection: dayWeatherMap[localDateInt].WindDirection,
				WindScale:     WindScale,
				WindSpeed:     WindSpeed,
				Humidity:      Humidity,
			}

			dayWeatherInstance := weather.DayWeather{}
			toExecSql := dayWeatherInstance.PreReplaceMysql(dayWeather)  // todo 可以打开测试

			log.Println("dayWeatherList", toExecSql)

			_, err = common.Engine.Exec(toExecSql)  // 参考 https://www.geek-share.com/detail/2717393840.html
			if err != nil {
				panic(err)
			}
		}

		time.Sleep(6 * time.Second)  // 先这样写，以后优化？
	}
}

/**
 * @Author Evan
 * @Description 获取城市列表
	tips （json格式化两个空格 https://tw.piliapp.com/json/formatter/）
	1.遍历数据库中按行存储的城市数据，首先转换成map嵌套格式，如下
	{
	  "北京市": {
		"北京市": {
		  "海淀区": "Haidian"
		}
	  },
	  "河北省": {
		"石家庄市": {
		  "长安区": "Changan"
		}
	  }
	}
	2.遍历上面map，得到list map的嵌套格式
	[
	  {
		"name": "北京市",
		"childs": [
		  {
			"name": "北京市",
			"childs": [
			  {
				"name": "海淀区",
				"pinyin": "haidian"
			  }
			]
		  }
		]
	  },
	  {
		"name": "河北省",
		"childs": [
		  {
			"name": "石家庄市",
			"childs": [
			  {
				"name": "长安区",
				"pinyin": "Changan"
			  }
			]
		  }
		]
	  }
	]

 * @Date 21:12 2021-02-19
 * @Param
 * @return
 **/
func GetCityList() ([]ProvinceItem, error) {
	// 读取city数据库
	cityListFromDb, err := city.GetAllCity()
	if err != nil {
		err = errnum.New(errnum.DbError, err)
		log.Println("GetAllCity err:", err.Error())
		return nil, err
	}
	if cityListFromDb == nil {
		return nil, errors.New("no cities")
	}

	// step1
	citiesMap := make(map[string]map[string]map[string]string, 0)
	var provinceSequenceList []string  // 记录`省`的顺序
	var citySequenceMapList = make(map[string][]string, 0)  // 记录`市`的顺序
	var districtSequenceMapList = make(map[string][]string, 0) // 记录`区`的顺序

	for i := 0; i < len(cityListFromDb); i++ {
		province := cityListFromDb[i].Province
		cityValue := cityListFromDb[i].City
		district := cityListFromDb[i].District
		pinYin := cityListFromDb[i].PinYin

		if _, ok := citiesMap[province]; !ok {
			citiesMap[province] = make(map[string]map[string]string, 0)
			provinceSequenceList = append(provinceSequenceList, province)
		}
		if _, ok := citiesMap[province][cityValue]; !ok {
			citiesMap[province][cityValue] = make(map[string]string, 0)
			citySequenceMapList[province] = append(citySequenceMapList[province], cityValue)
		}
		citiesMap[province][cityValue][district] = pinYin
		districtSequenceMapList[province+"::"+cityValue] = append(districtSequenceMapList[province+"::"+cityValue], district)

	}

	//// step2 返回的时候顺序随机
	//ProvinceList := make([]ProvinceItem, 0)  // 最后要的结果
	//for provinceName, provinceChildsValue := range citiesMap {
	//	// 最外面一层，这里是一个新的省份，初始化
	//	provinceChildsList := make([]ProvinceChildItem, 0)
	//
	//	for cityName, cityChildsValue := range provinceChildsValue {
	//		// 一个新的市，初始化
	//		cityChildsList := make([]CityChildItem, 0)
	//
	//		for districtName, pinYin := range cityChildsValue {
	//			cityChildsList = append(cityChildsList, CityChildItem{
	//				Name:   districtName,
	//				PinYin: pinYin,
	//			})
	//		}
	//		provinceChildsList = append(provinceChildsList, ProvinceChildItem{
	//			Name:  cityName,
	//			Childs: cityChildsList,
	//		})
	//	}
	//	ProvinceList = append(ProvinceList, ProvinceItem{
	//		Name:   provinceName,
	//		Childs: provinceChildsList,
	//	})
	//}
	// new step2 按照从数据库中顺次取回的顺序返回
	ProvinceList := make([]ProvinceItem, 0)  // 最后要的结果
	for _, provinceName := range provinceSequenceList {
		// 最外面一层，这里是一个新的省份，初始化
		provinceChildsList := make([]ProvinceChildItem, 0)

		for _, cityName := range citySequenceMapList[provinceName] {
			// 一个新的市，初始化
			cityChildsList := make([]CityChildItem, 0)

			for _, districtName := range districtSequenceMapList[provinceName + "::" + cityName] {
				cityChildsList = append(cityChildsList, CityChildItem{
					Name:   districtName,
					PinYin: citiesMap[provinceName][cityName][districtName],
				})
			}
			provinceChildsList = append(provinceChildsList, ProvinceChildItem{
				Name:   cityName,
				Childs: cityChildsList,
			})
		}
		ProvinceList = append(ProvinceList, ProvinceItem{
			Name:   provinceName,
			Childs: provinceChildsList,
		})
	}

	//jsons, errs := json.Marshal(ProvinceList[20:21]) //for test转换成JSON返回的是byte[]
	//if errs != nil {
	//	fmt.Println(errs.Error())
	//}
	//fmt.Println(string(jsons)) //byte[]转换成string 输出
	return ProvinceList, nil
}










