package twitterclient

import (
	"fmt"

	"github.com/dairlair/twitwatch/pkg/twitterclient/twitterstream"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	log "github.com/sirupsen/logrus"
)

// InstanceInterface defines the main object interface which is created by this package.
type InstanceInterface interface {
	// Creates Twitter Streaming API client and validates credentials
	Start() error
	AddStream(twitterstream.Interface)
	GetStreams() map[int64]twitterstream.Interface
	// Runs watching for twits according to specified streams
	Watch() error
	// Stops watching for all specified streams
	Unwatch() error
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

// Watch starts watching
func (instance *Instance) Watch() error {
	tracks := []string{"Tesla", "Microsoft"}
	log.Infof("Starting Stream with tracks [%v]", tracks)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = instance.onTweet

	// Filter for stream
	filterParams := &twitter.StreamFilterParams{
		Track:         tracks,
		StallWarnings: twitter.Bool(true),
		Language:      []string{"en"},
	}
	stream, err := instance.client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal("Stream not connected... ", err)
		return err
	}
	instance.source = stream

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	return nil
}

// Unwatch stops watching
func (instance *Instance) Unwatch() error {
	return nil
}

// Unwatch stops watching
func (instance *Instance) onTwit() {
	log.Infof("Stopping stream...")
	instance.source.Stop()
}

func (instance *Instance) onTweet(tweet *twitter.Tweet) {
	fmt.Printf("Tweet: %s\n", tweet.IDStr)
	fmt.Printf("%v\n\n", tweet)
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
