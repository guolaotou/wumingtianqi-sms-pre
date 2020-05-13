package weather

import "wumingtianqi-sms-pre/model/common"

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

func (s *DayWeather) Create() error {
	if _, err := common.Engine.InsertOne(s); err != nil {
		return err
	}
	return nil
}

func (s *DayWeather) Update() error {
	if _, err := common.Engine.Where(
		"city_pin_yin=?", s.CityPinYin).And(
			"date_id=?", s.DateId).Update(s); err != nil {
		return err
	}
	return nil
}

func (s *DayWeather) Delete() error {
	if _, err := common.Engine.Delete(s); err != nil {
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