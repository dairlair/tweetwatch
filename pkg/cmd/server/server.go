package server

import (
	"github.com/dairlair/tweetwatch/pkg/api/restapi"
	"github.com/dairlair/tweetwatch/pkg/service"
	"github.com/dairlair/tweetwatch/pkg/storage"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

// Config is configuration for the Server
type Config struct {
	LogLevel      string
	Postgres      storage.PostgresConfig
	RESTListen    int
	Twitterclient twitterclient.Config
}

type TwitterClientProvider func(config twitterclient.Config) twitterclient.Interface

type Providers struct {
	TwitterClientProvider TwitterClientProvider
}

// Instance stores the server state
type Instance struct {
	config     *Config
	providers  Providers
	grpcServer *grpc.Server
}

// NewInstance creates new server instance and copy config into that.
func NewInstance(config *Config, providers Providers) *Instance {
	s := &Instance{
		config:    config,
		providers: providers,
	}
	return s
}

// Start does startup all dependencies (postgres connections pool, gRPC server, etc..)
func (instance *Instance) Start() {

	// create postgres connections pool
	connPool := storage.CreatePostgresConnection(instance.config.Postgres)
	defer connPool.Close()

	// create storage instance
	storageInstance := storage.NewStorage(connPool)

	// create the twitterclient instance
	twitterclientInstance := instance.providers.TwitterClientProvider(instance.config.Twitterclient)

	// run service
	serviceInstance := service.NewService(storageInstance, twitterclientInstance)

	// create server
	server := restapi.NewServer(serviceInstance.API)
	server.Port = instance.config.RESTListen
	log.Infof("Start listen on %d port", server.Port)
	defer func() {
		err := server.Shutdown()
		if err != nil {
			log.Warnf("server shutdown got error: %instance\n", err)
		}
	}()

	// run server
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
