package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func RegisterServiceInstance(port uint64, serviceName, group string) error {
	cli, err := NewNamingClient()
	if err != nil {
		panic(err)
	}

	success, err := cli.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        port,
		Enable:      true,
		ServiceName: serviceName,
		GroupName:   group,
		Healthy:     true,
	})

	if !success || err != nil {
		return err
	}

	return nil
}

// 获取服务信息
func GetService(serviceName, group string) ([]model.Instance, error) {
	cli, err := NewNamingClient()
	if err != nil {
		panic(err)
	}

	service, err := cli.GetService(vo.GetServiceParam{
		ServiceName: serviceName,
		GroupName:   group,
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

	return cli.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        port,
		ServiceName: serviceName,
		GroupName:   group, // 默认值DEFAULT_GROUP
	})
}
