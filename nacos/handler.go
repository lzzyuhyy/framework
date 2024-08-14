package nacos

import (
	"framework/utils/getipv4address"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

const IPADDRESS = "127.0.0.1"

// get naocs config center's data
func GetNacosConfig() (string, error) {
	client, err := NewConfigClient()
	if err != nil {
		return "", err
	}

	return client.GetConfig(vo.ConfigParam{
		DataId: viper.GetString("nacos.dataID"),
		Group:  viper.GetString("nacos.group"),
	})
}

// register instance
func RegisterInstance(p uint64, sn, gn string) (bool, error) {
	client, err := NewNamingClient()
	if err != nil {
		return false, nil
	}

	return client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          getipv4address.GetLocalIPv4().String(),
		Port:        p,
		ServiceName: sn,
		Enable:      true,
		Healthy:     true, // please sure your grpc's health check is open
		GroupName:   gn,   // 默认值DEFAULT_GROUP
	})
}

// discover service
func DiscoverService(sn, gn string) (*model.Instance, error) {
	client, err := NewNamingClient()
	if err != nil {
		return nil, err
	}

	service, err := client.GetService(vo.GetServiceParam{
		ServiceName: sn,
		GroupName:   gn, // 默认值DEFAULT_GROUP
	})
	if err != nil {
		return nil, err
	}

	return &service.Hosts[0], nil
}

func DestroyInstance(p uint64, sn, gn string) (bool, error) {
	client, err := NewNamingClient()
	if err != nil {
		return false, nil
	}

	return client.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          getipv4address.GetLocalIPv4().String(),
		Port:        p,
		ServiceName: sn,
		GroupName:   gn, // 默认值DEFAULT_GROUP
	})
}
