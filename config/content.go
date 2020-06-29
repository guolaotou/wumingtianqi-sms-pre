package config

type Config struct {
	Main       Main    `json:"main"`
	Log        Log     `json:"log"`
	Weather    Weather `json:"weather"`
	Sms        Sms     `json:"sms"`
	Debug      bool    `json:"-"`
	ConfigFile string  `json:"-"`
}

type Main struct {
	Mysql string `json:"mysql"`
}

type Log struct {
	AppLogFile string `json:"app_file"`
}

type Weather struct {
	FakeData              bool   `json:"fake_data"`
	XinZhiWeatherDailyUrl string `json:"xin_zhi_weather_daily_url"`
	ApiKey                string `json:"api_key"` // 新知天气api配套的秘钥，可以在官网注册获取；可以试用一段时间，之后可能要收费
}

type Sms struct {
	SecretId    string `json:"secret_id"`
	SecretKey   string `json:"secret_key"`
	SmsSdkAppId string `json:"sms_sdk_app_id"`
	Sign        string `json:"sign"`
	TestPhone   string `json:"test_phone"`
}

var GlobalConfig *Config
