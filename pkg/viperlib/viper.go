package viperlib

import "github.com/spf13/viper"

var Viper *viper.Viper

func InitConfig() {
	Viper = viper.New()
	Viper.SetConfigName("config")
	Viper.AddConfigPath("./config")
	Viper.AddConfigPath(".")
	Viper.SetConfigType("yaml")
	if err := Viper.ReadInConfig(); err != nil {
		panic(err)
	}
	Viper.WatchConfig()
	Viper.AutomaticEnv()
}

func GetString(key string) string {
	return Viper.GetString(key)
}

func GetStrings(key string) []string {
	return Viper.GetStringSlice(key)
}
