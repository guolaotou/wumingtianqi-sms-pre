package weather

import (
	"testing"
	"wumingtianqi/libs/weather"
	test "wumingtianqi/testing"
)

// go clean -testcache  // 关闭go test的缓存，否则，create sql不会真的运行。cache说明：如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存
// go test -v libs/weather/weather_test.go

func TestWeather(t *testing.T) {
	test.Setup()
	//session := common.Engine.NewSession()
	//defer session.Close()

	weather.UpdateWeatherDaily()
}

// go clean -testcache & go test -v libs/weather/weather_test.go -test.run TestCity
func TestCity(t *testing.T) {
	test.Setup()
	_, err := weather.GetCityList()
	if err != nil {
		println("err", err.Error())
	}
}