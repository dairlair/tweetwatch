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
	jwtKey []byte
}

func NewService(s storage.Interface, t twitterclient.Interface) Service {
	service := Service{
		storage:s,
		twitterclient: t,
		tweetStreamsChannel: make(chan entity.TweetStreamsInterface, 100),
		// @TODO take jwt key from config
		jwtKey: []byte("something"),
	}
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	// create new service API
	api := operations.NewTweetwatchAPI(swaggerSpec)

	// set handlers
	api.Logger = log.Printf
	api.JWTAuth = service.JWTAuth
	api.SignupHandler = operations.SignupHandlerFunc(service.SignUpHandler)
	api.LoginHandler = operations.LoginHandlerFunc(service.LoginHandler)
	api.CreateTopicHandler = operations.CreateTopicHandlerFunc(service.CreateTopicHandler)
	api.GetUserTopicsHandler = operations.GetUserTopicsHandlerFunc(service.GetUserTopicsHandler)
	api.UpdateTopicHandler = operations.UpdateTopicHandlerFunc(service.UpdateTopicHandler)
	api.CreateStreamHandler = operations.CreateStreamHandlerFunc(service.CreateStreamHandler)
	api.GetStreamsHandler = operations.GetStreamsHandlerFunc(service.GetStreamsHandler)
	api.UpdateStreamHandler = operations.UpdateStreamHandlerFunc(service.UpdateStreamHandler)
	api.DeleteStreamHandler = operations.DeleteStreamHandlerFunc(service.DeleteStreamHandler)
	api.GetStatusHandler = operations.GetStatusHandlerFunc(service.GetStatusHandler)
	service.API = api

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

	service.twitterclient.AddStreams(streams)

	err = service.twitterclient.Start()
	if err != nil {
		log.Fatalf("twitterclient error: %s\n", err)
	}
	_ = service.twitterclient.Watch(service.tweetStreamsChannel)
}

/**
 * @TODO Add updated methods to use single stream, not a slice.
 * @TODO Use method with slice to stop all streams of inactivated topic or to add streams of activated topic.
 */
func (service *Service) addStreamsToWatching(streams []entity.StreamInterface) {
	service.twitterclient.Unwatch()
	service.twitterclient.AddStreams(streams)
	if err := service.twitterclient.Watch(service.tweetStreamsChannel); err != nil {
		log.Errorf("twitterclient does not resume watching: %s", err)
	}
}

func (service *Service) deleteStreamsFromWatching(streamIDs []int64) {
	service.twitterclient.Unwatch()
	service.twitterclient.DeleteStreams(streamIDs)
	if err := service.twitterclient.Watch(service.tweetStreamsChannel); err != nil {
		log.Errorf("twitterclient does not resume watching: %s", err)
	}
}