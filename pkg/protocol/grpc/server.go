package grpc

import (
	"context"
	v1 "github.com/dairlair/twitwatch/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Config contains options for gRPC server
type Config struct {
	ListenAddress string // ":8080", "127.0.0.1:84", etc...
}

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, v1API v1.TwitwatchServiceServer, gRPCConfig Config) error {
	listen, err := net.Listen("tcp", gRPCConfig.ListenAddress)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	v1.RegisterTwitwatchServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}

//// RunServer runs serve to expose TwitwatchService API over gRPC
//func RunServer(serviceServer apiV1.TwitwatchServiceServer, gRPCConfig Config) (*grpc.Server, error) {
//	listen, err := net.Listen("tcp", gRPCConfig.ListenAddress)
//	if err != nil {
//		return nil, err
//	}
//
//	// register service
//	server := grpc.NewServer()
//	apiV1.RegisterTwitwatchServiceServer(server, serviceServer)
//
//	// start gRPC server
//	log.Printf("starting gRPC server listen: %s\n", gRPCConfig.ListenAddress)
//	return server, server.Serve(listen)
//}
