package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func NewNacosClient() (string, error) {
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

type Config struct {
	Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host   string `yaml:"host"`
	Port   int64  `yaml:"port"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Dbname string `yaml:"dbname"`
}

func GetConfig() (string, error) {
	config, err := NewNacosClient()
	if err != nil {
		return "", nil
	}
	err = yaml.Unmarshal([]byte(config), &Config{})

	return config, err
}
