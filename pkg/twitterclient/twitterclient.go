package twitterclient

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	log "github.com/sirupsen/logrus"
)

// Instance structure is used to store the server's state
type Instance struct {
	config Config
	// Internal resources
	client *twitter.Client
	stream *twitter.Stream
}

// NewInstance creates new twitter instance scrapper
func NewInstance(config Config) *Instance {
	log.Infof("Twitter: consumer key=%s, consumer_secret=%s, access token=%s, access secret=%s",
		config.TwitterConsumerKey,
		config.TwitterConsumerSecret,
		config.TwitterAccessToken,
		config.TwitterAccessSecret,
	)

	return &Instance{
		config: config,
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
