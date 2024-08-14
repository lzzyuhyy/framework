package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

/*
	your nacos config should need filed as
	ipAddr,port,group,timeout,logDir,logLevel,cacheDir,dataID
*/

// get a nacos config
func getConfig() vo.NacosClientParam {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(viper.GetString("nacos.ipAddr"), viper.GetUint64("nacos.port"), constant.WithContextPath("/nacos")),
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

	return vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	}
}

func NewConfigClient() (config_client.IConfigClient, error) {
	return clients.NewConfigClient(getConfig())
}

func NewNamingClient() (naming_client.INamingClient, error) {
	return clients.NewNamingClient(getConfig())
}
