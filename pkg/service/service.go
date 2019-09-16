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
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	API *operations.GreeterAPI
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
	service.API = operations.NewGreeterAPI(swaggerSpec)
	// set handlers
	service.API.SignupHandler = operations.SignupHandlerFunc(service.SignUp)
	service.API.GetGreetingHandler = operations.GetGreetingHandlerFunc(
		func(params operations.GetGreetingParams) middleware.Responder {
			name := swag.StringValue(params.Name)
			if name == "" {
				name = "World"
			}

			greeting := fmt.Sprintf("Hello, %s!", name)
			return operations.NewGetGreetingOK().WithPayload(greeting)
		})
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


func (service *Service) SignUp (params operations.SignupParams) middleware.Responder {

	token, err := service.storage.SignUp(*params.User.Username, params.User.Password.String())

	if err != nil {
		return middleware.NotImplemented("Error handling not implemented")
	}

	message := fmt.Sprintf("User [%s] registered with token [%s]", *params.User.Username, token)
	payload := models.GeneralResponse{
		Message: &message,
	}
	return operations.NewSignupOK().WithPayload(&payload)
}