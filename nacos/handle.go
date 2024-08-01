package nacos

// 发现服务
//func RegisterServiceInstance() error {
//	address := viper.GetString("nacos.address")
//	port := viper.GetUint64("nacos.port")
//	serverName := viper.GetString("nacos.servername")
//	groupName := viper.GetString("nacos.group")
//	cli, err := newNacosClient(address, port)
//	if err != nil {
//		panic(err)
//	}
//	success, err := (*cli).RegisterInstance(vo.RegisterInstanceParam{
//		Ip:          address,
//		Port:        port,
//		ServiceName: serverName,
//		GroupName:   groupName,
//	})
//
//	if !success || err != nil {
//		return err
//	}
//
//	return nil
//}

// 获取服务信息
//func GetService() (model.Service, error) {
//	address := viper.GetString("nacos.address")
//	port := viper.GetUint64("nacos.port")
//	serverName := viper.GetString("nacos.servername")
//	groupName := viper.GetString("nacos.group")
//	cli, err := newNacosClient(address, port)
//	if err != nil {
//		panic(err)
//	}
//
//	return (*cli).GetService(vo.GetServiceParam{
//		ServiceName: serverName,
//		GroupName:   groupName, // 默认值DEFAULT_GROUP
//	})
//}
