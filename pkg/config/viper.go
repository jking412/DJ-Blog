package config

import (
	"DJ-Blog/internal/conf"
	"github.com/spf13/viper"
)

var internalViper *viper.Viper

func Init(configPath ...string) {
	internalViper = viper.New()
	internalViper.SetConfigName("config")
	if len(configPath) > 0 {
		internalViper.AddConfigPath(configPath[0])
	} else {
		internalViper.AddConfigPath(".")
	}
	internalViper.SetConfigType("yaml")
	if err := internalViper.ReadInConfig(); err != nil {
		panic(err)
	}
	internalViper.WatchConfig()
	internalViper.AutomaticEnv()
	conf.Init()
}
