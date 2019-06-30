package server

import (
	grpcServer "github.com/dairlair/twitwatch/pkg/protocol/grpc"
	serviceV1 "github.com/dairlair/twitwatch/pkg/service/v1"
	"github.com/dairlair/twitwatch/pkg/storage"
	"google.golang.org/grpc"

	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
)

// Config is configuration for the Server
type Config struct {
	Postgres storage.PostgresConfig
	GRPC     grpcServer.Config
}

// Instance stores the server state
type Instance struct {
	config     *Config
	connPool   *pgx.ConnPool
	storage    *storage.Storage
	grpcServer *grpc.Server
}

// NewInstance creates new server instance and copy config into that.
func NewInstance(config *Config) *Instance {
	s := &Instance{
		config: config,
	}
	return s
}

// Start does startup all dependencies (postgres connections pool, gRPC server, etc..)
func (s *Instance) Start() error {

	// Create postgres connections pool
	connPool := storage.CreatePostgresConnection(s.config.Postgres)
	defer connPool.Close()

	// Create storage instance
	s.storage = storage.NewStorage(connPool)

	// Run gRPC server
	v1API := serviceV1.NewTwitwatchServiceServer(s.storage)
	server, err := grpcServer.RunServer(v1API, s.config.GRPC)
	if err != nil {
		log.Fatalf("gRPC server error: %s\n", err)
		return err
	}
	s.grpcServer = server

	return nil
}
