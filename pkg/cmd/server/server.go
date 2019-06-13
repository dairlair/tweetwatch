package server

import (
	"errors"
	"fmt"
	grpcServer "github.com/dairlair/twitwatch/pkg/protocol/grpc"
	serviceV1 "github.com/dairlair/twitwatch/pkg/service/v1"
	"github.com/dairlair/twitwatch/pkg/storage"
	"google.golang.org/grpc"

	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
)

// PostgresConfig contains options for the Postgres connections pool
type PostgresConfig struct {
	DSN string
}

// Config is configuration for the Server
type Config struct {
	Postgres PostgresConfig
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

	// Create
	s.connPool = createPostgresConnection(s.config.Postgres)
	defer s.connPool.Close()

	// Create storage instance
	s.storage = storage.NewStorage(s.connPool)

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

func createPostgresConnection(config PostgresConfig) *pgx.ConnPool {
	pgConf, err := pgx.ParseURI(config.DSN)
	if err != nil {
		msg := fmt.Sprintf("Can not parse Postgres DSN: %s", err)
		panic(errors.New(msg))
	}
	log.Infof("PostgreSQL: host=%s, port=%d, username=%s, database=%s",
		pgConf.Host,
		pgConf.Port,
		pgConf.User,
		pgConf.Database,
	)
	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     pgConf.Host,
			Port:     pgConf.Port,
			User:     pgConf.User,
			Password: pgConf.Password,
			Database: pgConf.Database,
			RuntimeParams: map[string]string{
				"application_name": "twitwatch",
			},
		},
		MaxConnections: 1,
	})

	if err != nil {
		msg := fmt.Sprintf("Can not connect to Postgres: %s", err)
		panic(errors.New(msg))
	}

	return connPool
}