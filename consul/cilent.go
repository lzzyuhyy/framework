package consul

import (
	"fmt"
	"framework/utils/getipv4address"
	capi "github.com/hashicorp/consul/api"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
)

func NewClient() (*capi.Client, error) {
	return capi.NewClient(capi.DefaultConfig())
}

type ServiceConfig struct {
	Port int    `json:"Port"`
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

func RegisterService(sc ServiceConfig) error {
	client, err := NewClient()
	if err != nil {
		return nil
	}
	ip, err := getipv4address.GetLocalIPv4()
	if err != nil {
		return err
	}

	return client.Agent().ServiceRegister(
		&capi.AgentServiceRegistration{
			Address: ip.String(),
			Port:    sc.Port,
			ID:      sc.ID,
			Name:    sc.Name,
			Tags:    []string{"grpc"},
			Check: &capi.AgentServiceCheck{
				CheckID:                        sc.ID,
				Name:                           "check" + sc.Name,
				Interval:                       "5s",                                       // 指定运行此检查的频率
				Timeout:                        "5s",                                       // 超时时间
				GRPC:                           fmt.Sprintf("%s:%d", ip.String(), sc.Port), // 健康检查HTTP请求
				DeregisterCriticalServiceAfter: "30s",                                      // 触发注销的时间
			},
		},
	)
}

func DiscoverService(serviceName string) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("consul://127.0.0.1:8500/%v?wait=14s", serviceName)
	return grpc.NewClient(target)
}

func PutKey(key string, val []byte) error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	kv := client.KV()
	_, err = kv.Put(&capi.KVPair{
		Key:   key,
		Value: val,
	}, nil)
	return err
}

func GetKeyInfo(key string) ([]byte, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	kv := client.KV()
	get, _, err := kv.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return get.Value, err
}
