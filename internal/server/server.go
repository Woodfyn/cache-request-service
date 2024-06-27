package server

import (
	"net"

	cache_service "github.com/Woodfyn/chat-api-cache-service/pkg/proto"
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

func (s *Server) Run(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	cache_service.RegisterCacheServiceServer(s.grpcServer, s.cacheServiceServer)

	if err := s.grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
