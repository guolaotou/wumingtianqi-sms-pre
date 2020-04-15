package config

type Config struct {
	Log        Log     `json:"log"`
	Weather    Weather `json:"weather"`
	Debug      bool    `json:"-"`
	ConfigFile string  `json:"-"`
}

type Log struct {
	AppLogFile string `json:"app_file"`
}

type Weather struct {
	XinZhiWeatherDailyUrl string `json:"xin_zhi_weather_daily_url"`
	ApiKey                string `json:"api_key"`  // 新知天气api配套的秘钥，可以在官网注册获取；可以试用一段时间，之后可能要收费
}

var GlobalConfig *Config
