package config

import (
	"github.com/spf13/viper"
)

var Config LocalConfig

func InitConfig() {
	viper.SetConfigFile("./config/.env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	readConfig()
}

func readConfig() {
	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
