package grpc

import (
	"encoding/json"
	"fmt"
	"framework/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func NewGrpcClientConsul(sn string, handler func(s *grpc.Server)) error {
	info, err := consul.GetKeyInfo(sn)
	if err != nil {
		return nil
	}
	var sc consul.ServiceConfig
	err = json.Unmarshal(info, &sc)
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", sc.Port))
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

	err = consul.RegisterService(sc)
	if err != nil {
		return err
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
	return nil
}
