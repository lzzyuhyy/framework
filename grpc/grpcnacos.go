package grpc

import (
	"fmt"
	"framework/nacos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
	"log"
	"net"
)

const healthCheckService = "grpc.health.v1.Health"

type Config struct {
	Service
}

type Service struct {
	Port  uint64 `yaml:"port"`
	Group string `yaml:"group"`
}

func NewGrpcClientNacos(sn string, handler func(s *grpc.Server)) {
	var conf Config
	config, err := nacos.GetNacosConfig()
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(config), &conf)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Service.Port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	healthcheck := health.NewServer()
	healthcheck.SetServingStatus(healthCheckService, healthgrpc.HealthCheckResponse_NOT_SERVING)
	healthgrpc.RegisterHealthServer(s, healthcheck)

	// register server
	handler(s)

	reflection.Register(s)

	// register instance
	instance, err := nacos.RegisterInstance(conf.Service.Port, sn, conf.Service.Group)
	if err != nil || !instance {
		panic(err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
