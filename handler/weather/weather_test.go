package weather

import (
	"testing"
	"wumingtianqi-sms-pre/handler/weather"
	"wumingtianqi-sms-pre/model/common"
	test "wumingtianqi-sms-pre/testing"
)

// go test -v handler/weather/weather_test.go

func TestWeather(t *testing.T) {
	test.Setup()
	session := common.Engine.NewSession()
	defer session.Close()

	weather.StorageWeatherDaily()
}
