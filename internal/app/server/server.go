package server

import (
	"net"

	cache_service "github.com/Woodfyn/cache-service/pkg/proto"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer         *grpc.Server
	cacheServiceServer cache_service.CacheServiceServer
}

func NewServer(cacheServiceServer cache_service.CacheServiceServer) *Server {
	return &Server{
		grpcServer:         grpc.NewServer(),
		cacheServiceServer: cacheServiceServer,
	}
}

func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	cache_service.RegisterCacheServiceServer(s.grpcServer, s.cacheServiceServer)

	return s.grpcServer.Serve(lis)
}
