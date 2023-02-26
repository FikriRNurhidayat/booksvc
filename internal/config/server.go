package config

import "github.com/spf13/viper"

func GetPort() int {
	return viper.GetInt("port")
}
