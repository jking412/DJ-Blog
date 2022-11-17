package config

import (
	"github.com/spf13/viper"
)

var internalViper *viper.Viper

func init() {
	internalViper = viper.New()
	internalViper.SetConfigName("config")
	internalViper.AddConfigPath(".")
	internalViper.AddConfigPath("./config")
	internalViper.SetConfigType("yaml")
	if err := internalViper.ReadInConfig(); err != nil {
		panic(err)
	}
	internalViper.WatchConfig()
	internalViper.AutomaticEnv()
}
