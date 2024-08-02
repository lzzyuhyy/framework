package grpc_core

import (
	"fmt"
	"github.com/lzzyuhyy/framework/nacos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
	"log"
	"net"
)

type Config struct {
	Service
}

type Service struct {
	Name  string `yaml:"name"`
	Port  uint64 `yaml:"port"`
	Group string `yaml:"group"`
}

// 注册grpc服务
func RegisterGrpcService(port int64, server func(s *grpc.Server)) error {
	config, err := nacos.GetConfig()
	if err != nil {
		return err
	}

	var conf Config
	err = yaml.Unmarshal([]byte(config), &conf)
	if err != nil {
		return err
	}

	//建立tpc链接
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Service.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	//初始化grpc服务
	s := grpc.NewServer()

	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(s, healthcheck)

	server(s)
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())

	// 注册服务
	err = nacos.RegisterServiceInstance(conf.Service.Port, conf.Service.Name, conf.Service.Group)
	if err != nil {
		return err
	}

	// 启动grpc
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}
