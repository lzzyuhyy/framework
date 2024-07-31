package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

// 发现服务
func RegisterServiceInstance() error {
	address := viper.GetString("nacos.address")
	port := viper.GetUint64("nacos.port")
	serverName := viper.GetString("nacos.servername")
	groupName := viper.GetString("nacos.group")
	cli, err := newNacosClient(address, port)
	if err != nil {
		panic(err)
	}
	success, err := (*cli).RegisterInstance(vo.RegisterInstanceParam{
		Ip:          address,
		Port:        port,
		ServiceName: serverName,
		GroupName:   groupName,
	})

	if !success || err != nil {
		return err
	}

	return nil
}

// 获取服务信息
func GetService(address string, port uint64, serverName, groupName string) (model.Service, error) {
	cli, err := newNacosClient(address, port)
	if err != nil {
		panic(err)
	}

	return (*cli).GetService(vo.GetServiceParam{
		ServiceName: serverName,
		GroupName:   groupName, // 默认值DEFAULT_GROUP
	})
}
