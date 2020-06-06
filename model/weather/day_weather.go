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


func (s *DayWeather)ReplaceMysql(dayWeatherList []DayWeather) error {  // todo 尝试看xorm源码，实现replace功能 !important
	//values := "VALUES(%s)"
	// todo today 把拼接字符串写了；
	// todo 获取天气并存取的代码，至少北京完成，控制1小时抓取1次；存取 & 更新
	// todo Then 2个模式跑？拼接提醒信息
	// todo 邮件发送
	// ops 挂上服务器
	_ = `
			REPLACE INTO wumingtianqi.day_weather (city_pin_yin, date_id, text_day, code_day, text_night, code_night, high, low, wind_direction, wind_scale, wind_speed, humidity)
			VALUES ('tianjin', '20200508', '晴', '2', '雨'
				, '1', '40', '10', '南', '2'
				, '10', '20'), ('tianjin', '20200507', '晴', '1', '晴'
				, '1', '40', '10', '南', '2'
				, '10', '20');
			`
	return nil
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