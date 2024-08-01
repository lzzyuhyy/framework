package nacos

import (
	"github.com/lzzyuhyy/framework/utils/getipv4address"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

func RegisterServiceInstance(port uint64, serviceName, group string) error {
	cli, err := NewNamingClient()
	if err != nil {
		panic(err)
	}

	ip, err := getipv4address.GetLocalIPv4()
	if err != nil {
		return err
	}
	success, err := cli.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		Enable:      true,
		ServiceName: serviceName,
		GroupName:   group,
	})

	if !success || err != nil {
		return err
	}

	return nil
}

// 获取服务信息
func GetService() ([]model.Instance, error) {
	cli, err := NewNamingClient()
	if err != nil {
		panic(err)
	}

	service, err := cli.GetService(vo.GetServiceParam{
		ServiceName: viper.GetString("nacos.servicename"),
		GroupName:   viper.GetString("nacos.group"),
	})
	if err != nil {
		return nil, err
	}

	return service.Hosts, nil
}

// 注销实例
func DeRegistryInstance(port uint64, serviceName, group string) (bool, error) {
	cli, err := NewNamingClient()
	if err != nil {
		panic(err)
	}

	ip, err := getipv4address.GetLocalIPv4()
	if err != nil {
		return false, err
	}

	return cli.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: serviceName,
		GroupName:   group, // 默认值DEFAULT_GROUP
	})
}
