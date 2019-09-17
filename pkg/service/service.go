package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/restapi"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/storage"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	"github.com/go-openapi/loads"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	API *operations.TweetwatchAPI
	storage storage.Interface
	twitterclient twitterclient.Interface
	tweetStreamsChannel chan entity.TweetStreamsInterface
}

func NewService(s storage.Interface, t twitterclient.Interface) Service {
	service := Service{
		storage:s,
		twitterclient: t,
		tweetStreamsChannel: make(chan entity.TweetStreamsInterface, 100),
	}
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	// create new service API
	service.API = operations.NewTweetwatchAPI(swaggerSpec)

	// set handlers
	//service.API.IsRegisteredAuth = service.login
	service.API.IsRegisteredAuth = func(user string, pass string) (interface{}, error) {
		// The header: Authorization: Basic {base64 string} has already been decoded by the runtime as a
		// username:password pair
		log.Errorf("IsRegisteredAuth handler called\n")
		return service.login(user, pass)
	}
	service.API.Logger = log.Printf
	service.API.SignupHandler = operations.SignupHandlerFunc(service.SignUpHandler)
	service.API.LoginHandler = operations.LoginHandlerFunc(service.LoginHandler)

	// up...
	service.up()

	return service
}

func (service *Service) up() {
	log.Infof("Tweetwatch service up...")
	go func(input chan entity.TweetStreamsInterface, storage storage.Interface) {
		for tweetStreams := range input {
			log.Infof("Store tweet to the database: %d\n", tweetStreams.GetTweet().GetTwitterID())
			_, err := storage.AddTweetStreams(tweetStreams)
			if err != nil {
				log.Fatalf("storage error: %s\n", err)
			}
		}
	} (service.tweetStreamsChannel, service.storage)
	log.Infof("Tweetwatch service is ready to accept tweets")

	// Restore streams
	streams, err := service.storage.GetStreams()
	if err != nil {
		log.Fatalf("failed to restore streams: %s\n", err)
	}

	for _, stream := range streams {
		service.twitterclient.AddStream(stream)
	}

	err = service.twitterclient.Start()
	if err != nil {
		log.Fatalf("twitterclient error: %s\n", err)
	}
	_ = service.twitterclient.Watch(service.tweetStreamsChannel)
}