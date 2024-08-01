package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

func newNacosClient() (string, error) {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(viper.GetString("nacos.address"), viper.GetUint64("nacos.port"), constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(viper.GetUint64("nacos.timeout")),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(viper.GetString("nacos.logDir")),
		constant.WithCacheDir(viper.GetString("nacos.cacheDir")),
		constant.WithLogLevel(viper.GetString("nacos.logLevel")),
	)

	// 动态获取
	configClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		return "", err
	}

	return configClient.GetConfig(vo.ConfigParam{
		DataId: viper.GetString("nacos.dataID"),
		Group:  viper.GetString("naocs.group"),
	})
}
