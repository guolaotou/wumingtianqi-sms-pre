package weather

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

