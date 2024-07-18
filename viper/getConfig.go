package viper

import (
	"github.com/spf13/viper"
)

// 读取用户配置文件中的数据用以做链接等操作，避免直接的配置信息传入
func GetConfig(path string) error {
	viper.SetConfigFile(path)   // 文件路径
	return viper.ReadInConfig() // 抛给外层使用---让用户按自己的业务逻辑处理错误
}
