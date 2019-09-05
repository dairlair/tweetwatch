package server

import (
	grpcServer "github.com/dairlair/tweetwatch/pkg/protocol/grpc"
	serviceV1 "github.com/dairlair/tweetwatch/pkg/service/v1"
	"github.com/dairlair/tweetwatch/pkg/storage"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

// Config is configuration for the Server
type Config struct {
	Postgres      storage.PostgresConfig
	GRPC          grpcServer.Config
	Twitterclient twitterclient.Config
}

type TwitterClientProvider func (config twitterclient.Config) twitterclient.InstanceInterface

type Providers struct {
	CreateTwitterclient TwitterClientProvider
}

// Instance stores the server state
type Instance struct {
	config        *Config
	providers     Providers
	storage       *storage.Storage // @TODO Use storage interface instead of pointer
	grpcServer    *grpc.Server
	twitterClient twitterclient.InstanceInterface
}

// NewInstance creates new server instance and copy config into that.
func NewInstance(config *Config, providers Providers) *Instance {
	s := &Instance{
		config: config,
		providers: providers,
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
	s.config.Twitterclient.Storage = s.storage

	// Create the twitterclient instance
	s.twitterClient = s.providers.CreateTwitterclient(s.config.Twitterclient)
	err := s.twitterClient.Start()
	if err != nil {
		log.Fatalf("twitterclient error: %s\n", err)
		return err
	}

	err = s.twitterClient.Watch()
	if err != nil {
		log.Fatalf("twitterclient error: %s\n", err)
		return err
	}

	// Run gRPC server
	// @TODO Pass twitterClient as dependency to the service.
	v1API := serviceV1.NewTwitwatchServiceServer(s.storage)
	server, err := grpcServer.RunServer(v1API, s.config.GRPC)
	if err != nil {
		log.Fatalf("gRPC server error: %s\n", err)
		return err
	}
	s.grpcServer = server

	return nil
}
