package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	Address   string `mapstructure:"Address"`
	IndexName string `mapstructure:"IndexName"`
	Size      int64  `mapstructure:"Size"`
	Refresh   string `mapstructure:"Refresh"`
}

var Configuration Config

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return
}
