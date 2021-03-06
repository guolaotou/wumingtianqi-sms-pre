package weather

import (
	"testing"
	"wumingtianqi/model/common"
	"wumingtianqi/model/weather"
	test "wumingtianqi/testing"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test model/weather/day_weather_test.go
// go test -v model/weather/day_weather_test.go

func TestDayWeather(t *testing.T) {
	test.Setup()
	session := common.Engine.NewSession()
	defer session.Close()

	// 1. 新建
	cityCode := "WX4FBXXFKE4F"
	dateId := 20200507
	d := &weather.DayWeather{
		CityCode:      cityCode,
		DateId:        dateId,
		TextDay:       "晴",
		CodeDay:       1,
		TextNight:     "晴",
		CodeNight:     1,
		High:          26,
		Low:           10,
		WindDirection: "南",
		WindScale:     2,
		WindSpeed:     10,
		Humidity:      20,
	}

	t.Log("*** begin create session****** ")

	//if _, err := session.Insert(d); err != nil {
	if err := d.Create(); err != nil {
		panic(err)
	}

	// 2. 查询
	t.Log("*** begin query session****** ")
	d2, has, err := weather.QueryByCityDate(cityCode, dateId)
	if  err != nil || !has {
		t.Error("city & date not found")
		panic(err)
	} else {
		t.Log("dayWeather: ", d2)
	}
	t.Log("*** end query session****** ")

	// 3. 更改
	t.Log("*** begin update session****** ")
	d2.High = 40
	d2.Update()
	d3, _, _ := weather.QueryByCityDate(cityCode, dateId)
	t.Log("dayWeather: ", d3)
	t.Log("*** end update session****** ")

	// 4.删除
	t.Log("*** begin delete session****** ")
	d.Delete()
	t.Log("*** end delete session****** ")
}
