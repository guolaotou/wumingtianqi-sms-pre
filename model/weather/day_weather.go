package weather

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"wumingtianqi-sms-pre/model/common"
	"xorm.io/core"
)

// 历史天气表
type DayWeather struct {
	CityPinYin    string `json:"city_pin_yin" xorm:"pk VARCHAR(30)"`
	DateId        int    `json:"date_id" xorm:"pk INT(11)"`
	TextDay       string `json:"text_day" xorm:"VARCHAR(30)"`
	CodeDay       int    `json:"code_day" xorm:"INT(11)"`
	TextNight     string `json:"text_night" xorm:"VARCHAR(30)"`
	CodeNight     int    `json:"code_night" xorm:"INT(11)"`
	High          int    `json:"high" xorm:"INT(11)"`
	Low           int    `json:"low" xorm:"INT(11)"`
	WindDirection string `json:"wind_direction" xorm:"VARCHAR(30)"`
	WindScale     int    `json:"wind_scale" xorm:"INT(11)"`
	WindSpeed     int    `json:"wind_speed" xorm:"INT(11)"`
	Humidity      int    `json:"humidity" xorm:"INT(11)"`
}

func (m *DayWeather)TableName() string{
	return "day_weather"
}

func (m *DayWeather)PreReplaceMysql(dayWeatherModel DayWeather) string {  // todo 尝试看xorm源码，实现replace功能 !important
	//values := "VALUES(%s)"
	// todo today 把拼接字符串写了；
	// todo 获取天气并存取的代码，至少北京完成，控制1小时抓取1次；存取 & 更新
	// todo Then 2个模式跑？拼接提醒信息
	// todo 邮件发送
	// ops 挂上服务器

	bb := strings.ToLower("Name")
	fmt.Println("bb", bb)

	t := reflect.TypeOf(dayWeatherModel)
	val := reflect.ValueOf(dayWeatherModel)

	var columnStr string
	var valueStr string
	count := val.NumField()
	mapper := core.GonicMapper{}
	for i := 0; i < count; i++ { // todo 以后可以封装一个函数，Values后面一次性拼接多个model
		// 拼接字段
		columnStr += mapper.Obj2Table(t.Field(i).Name)
		if i != count - 1 {
			columnStr += ", "
		}

		// 拼接value
		field := val.Field(i)
		switch field.Kind() {
		case reflect.Int:
			valueStr += strconv.Itoa(int(field.Int()))
		case reflect.String:
			valueStr += "'" + field.String() + "'"
		}
		if i != count - 1 {
			valueStr += ", "
		}
	}
	toExecSql := fmt.Sprintf(`REPLACE INTO %s (%s) VALUES (%s);`, dayWeatherModel.TableName(), columnStr, valueStr)
	fmt.Println("toExecSql", toExecSql)

	//var aa  string
	//aa = "`"
	//	// 写一个函数，用于拼接sql值。里面识别非int,需要加引号？
	//	fmt.Println("lala", reflect.TypeOf(dayWeatherList[i].CityPinYin))
	//	aa += "'" + dayWeatherList[i].CityPinYin + "'"
	//	if i != len(dayWeatherList) - 1 {
	//		aa += ","
	//	}
	//aa += "`"


	_ = `
			REPLACE INTO wumingtianqi.day_weather (city_pin_yin, date_id, text_day, code_day, text_night, code_night, high, low, wind_direction, wind_scale, wind_speed, humidity)
			VALUES ('tianjin', '20200508', '晴', '2', '雨'
				, '1', '40', '10', '南', '2'
				, '10', '20'), ('tianjin', '20200507', '晴', '1', '晴'
				, '1', '40', '10', '南', '2'
				, '10', '20');
			`

	return toExecSql
}

func (m *DayWeather) Create() error {
	if _, err := common.Engine.InsertOne(m); err != nil {
		return err
	}
	return nil
}

func (m *DayWeather) Update() error {
	if _, err := common.Engine.Where(
		"city_pin_yin=?", m.CityPinYin).And(
			"date_id=?", m.DateId).Update(m); err != nil {
		return err
	}
	return nil
}

func (m *DayWeather) Delete() error {
	if _, err := common.Engine.Delete(m); err != nil {
		return err
	}
	return nil
}

func QueryByCityDate(city string, date int) (*DayWeather, bool, error) {
	d := new(DayWeather)
	has, err := common.Engine.Where(
		"city_pin_yin=?", city).And(
			"date_id=?", date).Get(d)
	return d, has, err
}