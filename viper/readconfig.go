package viper

import (
	"github.com/spf13/viper"
)

func ReadConfig(path string) error {
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}
