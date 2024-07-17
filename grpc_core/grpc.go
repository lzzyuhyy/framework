package grpc_core

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// 注册grpc服务
func RegisterGrpcService(port int64, server func(s *grpc.Server)) error {
	//建立tpc链接
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	//初始化grpc服务
	s := grpc.NewServer()
	server(s)
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	// 启动grpc
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}
