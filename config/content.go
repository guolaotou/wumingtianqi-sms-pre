package config

type Config struct {
	Log        Log    `json:"log"`
	Debug      bool   `json:"-"`
	ConfigFile string `json:"-"`
}

type Log struct {
	AppLogFile string `json:"app-file"`
}

var GlobalConfig *Config
