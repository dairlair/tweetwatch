package grpc

import (
	apiV1 "github.com/dairlair/twitwatch/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Config contains options for gRPC server
type Config struct {
	ListenAddress string // ":8080", "127.0.0.1:84", etc...
}

// RunServer runs gRPC service to publish ToDo service
func RunServer(v1API apiV1.TwitwatchServiceServer, gRPCConfig Config) (*grpc.Server, error) {
	listen, err := net.Listen("tcp", gRPCConfig.ListenAddress)
	if err != nil {
		return nil, err
	}

	// register service
	server := grpc.NewServer()
	apiV1.RegisterTwitwatchServiceServer(server, v1API)

	// start gRPC server
	log.Printf("starting gRPC server on address [%s]\n", gRPCConfig.ListenAddress)
	return server, server.Serve(listen)
}