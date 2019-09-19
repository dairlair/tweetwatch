package service

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/storage"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
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
	api := operations.NewTweetwatchAPI(swaggerSpec)

	// set handlers
	api.Logger = log.Printf
	api.IsRegisteredAuth  =service.isRegisteredAuth
	api.SignupHandler = operations.SignupHandlerFunc(service.SignUpHandler)
	api.LoginHandler = operations.LoginHandlerFunc(service.LoginHandler)
	api.AccountHandler = operations.AccountHandlerFunc(func(params operations.AccountParams, user *models.User) middleware.Responder {
		message := fmt.Sprintf("AccountHandler [%v], [%v]", params, user.Email)
		payload := models.GeneralResponse{
			Message: &message,
		}
		return operations.NewAccountOK().WithPayload(&payload)
	})
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

	for _, stream := range streams {
		service.twitterclient.AddStream(stream)
	}

	err = service.twitterclient.Start()
	if err != nil {
		log.Fatalf("twitterclient error: %s\n", err)
	}
	_ = service.twitterclient.Watch(service.tweetStreamsChannel)
}