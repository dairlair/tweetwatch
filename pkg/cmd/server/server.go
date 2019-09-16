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
	LogLevel string
	Postgres      storage.PostgresConfig
	GRPC          grpcServer.Config
	Twitterclient twitterclient.Config
}

type TwitterClientProvider func(config twitterclient.Config) twitterclient.Interface

type Providers struct {
	TwitterClientProvider TwitterClientProvider
}

// Instance stores the server state
type Instance struct {
	config        *Config
	providers     Providers
	grpcServer    *grpc.Server
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
	storageInstance := storage.NewStorage(connPool)

	// Create the twitterclient instance
	twitterClient := s.providers.TwitterClientProvider(s.config.Twitterclient)

	// Run gRPC server
	v1API := serviceV1.NewTweetwatchServiceServer(storageInstance, twitterClient)
	server, err := grpcServer.RunServer(v1API, s.config.GRPC)
	if err != nil {
		log.Fatalf("gRPC server error: %s\n", err)
		return err
	}
	s.grpcServer = server

	return nil
}
