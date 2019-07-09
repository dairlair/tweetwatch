package twitterclient

import (
	"github.com/dairlair/twitwatch/pkg/twitterclient/twitterstream"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	log "github.com/sirupsen/logrus"
)

// InstanceInterface defines the main object interface which is created by this package.
type InstanceInterface interface {
	Start() error
	AddStream(twitterstream.Interface)
	GetStreams() map[int64]twitterstream.Interface
}

// Instance structure is used to store the server's state
type Instance struct {
	config Config
	// Internal resources
	client  *twitter.Client
	source  *twitter.Stream
	streams map[int64]twitterstream.Interface
}

// NewInstance creates new twitter instance scrapper
func NewInstance(config Config) InstanceInterface {
	log.Infof("Twitter: consumer key=%s, consumer_secret=%s, access token=%s, access secret=%s",
		config.TwitterConsumerKey,
		config.TwitterConsumerSecret,
		config.TwitterAccessToken,
		config.TwitterAccessSecret,
	)

	return &Instance{
		config:  config,
		streams: make(map[int64]twitterstream.Interface),
	}
}

// Start function runs twitter client and validates credentials for Twitter Streaming API
func (instance *Instance) Start() error {
	client, err := createTwitterClient(instance.config)
	if err != nil {
		log.Error("Authentication is failed. ", err)
		return err
	}

	instance.client = client

	return nil
}

// AddStream adds desired stream to the current instance of twitterclient
func (instance *Instance) AddStream(stream twitterstream.Interface) {
	instance.streams[stream.GetID()] = stream
}

// GetStreams returns all the streams from the current instance of twitterclient
func (instance *Instance) GetStreams() map[int64]twitterstream.Interface {
	return instance.streams
}

func createTwitterClient(config Config) (*twitter.Client, error) {
	oauthConfig := oauth1.NewConfig(config.TwitterConsumerKey, config.TwitterConsumerSecret)
	token := oauth1.NewToken(config.TwitterAccessToken, config.TwitterAccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	
	err := validateTwitterClientCredentials(client)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func validateTwitterClientCredentials(client *twitter.Client) error {
	// We use this hack to validate Twitter OAuth credentials
	_, _, err := client.Trends.Available()

	return err
}
