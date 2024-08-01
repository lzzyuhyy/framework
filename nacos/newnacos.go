package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

func getConfig() vo.NacosClientParam {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(viper.GetString("nacos.address"), viper.GetUint64("nacos.port"), constant.WithContextPath("/nacos")),
	}

	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(viper.GetUint64("nacos.timeout")),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(viper.GetString("nacos.logDir")),
		constant.WithCacheDir(viper.GetString("nacos.cacheDir")),
		constant.WithLogLevel(viper.GetString("nacos.logLevel")),
	)

	return vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	}
}

func GetConfig() (string, error) {
	configClient, err := clients.NewConfigClient(getConfig())
	if err != nil {
		return "", err
	}

	return configClient.GetConfig(vo.ConfigParam{
		DataId: viper.GetString("nacos.dataID"),
		Group:  viper.GetString("nacos.group"),
	})
}

func NewNamingClient() (naming_client.INamingClient, error) {
	return clients.NewNamingClient(getConfig())
}
