package config

import (
	"github.com/intel-go/fastjson"
	"io/ioutil"
	"os"
)

func NewConfig() *Config {
	return &Config{
		Log:   Log {
			"app.log",
		},
		Debug: false,
	}
}

func LoadConfig() (*Config, error) {
	cfg := NewConfig()
	configPath := `conf/config.json`
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configPath = `conf/config.template.json`
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			panic("No `config.json` or `conf` folder found in current working directory: " + configPath)
		}
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = fastjson.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	GlobalConfig = cfg
	return cfg, nil
}
