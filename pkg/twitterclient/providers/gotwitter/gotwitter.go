package gotwitter

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"

	log "github.com/sirupsen/logrus"
)

// Instance structure is used to store the server's state
type Instance struct {
	config twitterclient.Config
	// Internal resources
	storage twitterclient.StorageInterface
	client  *twitter.Client
	source  *twitter.Stream
	streams map[int64]entity.StreamInterface
}

// NewInstance creates new twitter instance scrapper
func NewInstance(config twitterclient.Config) twitterclient.InstanceInterface {
	log.Infof("Twitter: consumer key=%s, consumer_secret=%s, access token=%s, access secret=%s",
		config.TwitterConsumerKey,
		config.TwitterConsumerSecret,
		config.TwitterAccessToken,
		config.TwitterAccessSecret,
	)

	if config.Storage == nil {
		log.Warn("Storage not attached, tweets won't be saved")
	}

	return &Instance{
		config:  config,
		storage: config.Storage,
		streams: make(map[int64]entity.StreamInterface),
	}
}

// Start function runs twitter client and validates credentials for Twitter Streaming API
func (instance *Instance) Start() error {
	// Init Twitter Streaming API client
	client, err := createTwitterClient(instance.config)
	if err != nil {
		return err
	}
	instance.client = client

	// Restore state from the storage
	if err = instance.restoreStreams(); err != nil {
		return err
	}

	return nil
}

// AddStream adds desired stream to the current instance of twitterclient
func (instance *Instance) AddStream(stream entity.StreamInterface) {
	instance.streams[stream.GetID()] = stream
}

// GetStreams returns all the streams from the current instance of twitterclient
func (instance *Instance) GetStreams() map[int64]entity.StreamInterface {
	return instance.streams
}

// Watch starts watching
func (instance *Instance) Watch() error {
	tracks := instance.getTracks()
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
func (instance *Instance) Unwatch() {
	log.Infof("Stopping stream...")
	instance.source.Stop()
}

func (instance *Instance) getTracks() []string {
	var tracks []string
	for _, stream := range instance.GetStreams() {
		tracks = append(tracks, stream.GetTrack())
	}
	return tracks
}

func (instance *Instance) onTweet(tweet *twitter.Tweet) {
	fmt.Printf("Tweet: %s\n", tweet.IDStr)
	fmt.Printf("%v\n\n", tweet)
	instance.processTweet(createTweetEntity(tweet))
}

func (instance *Instance) processTweet(tweet entity.TwitInterface) {
	fmt.Printf("Process tweet: %v\n", tweet)
	if instance.storage == nil {
		return
	}
	id, err := instance.storage.AddTwit(tweet)
	if err != nil {
		log.Errorf("Tweet processing error. %s", err)
	} else {
		fmt.Printf("Tweet has been processed sucessfully and saved with ID: %d\n", id)
	}
}

func createTweetEntity(tweet *twitter.Tweet) entity.TwitInterface {
	var fullText string
	if tweet.ExtendedTweet != nil {
		fullText = tweet.ExtendedTweet.FullText
	} else {
		fullText = tweet.Text
	}
	return &entity.Twit{
		ID:            tweet.ID,
		TwitterID:     tweet.ID,
		TwitterUserID: tweet.User.ID,
		FullText:      fullText,
		CreatedAt:     tweet.CreatedAt,
	}
}

func createTwitterClient(config twitterclient.Config) (*twitter.Client, error) {
	oauthConfig := oauth1.NewConfig(config.TwitterConsumerKey, config.TwitterConsumerSecret)
	token := oauth1.NewToken(config.TwitterAccessToken, config.TwitterAccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	err := validateTwitterClientCredentials(client)

	if err != nil {
		log.Error("Authentication is failed. ", err)
		return nil, err
	}

	return client, nil
}

func validateTwitterClientCredentials(client *twitter.Client) error {
	// We use this hack to validate Twitter OAuth credentials
	_, _, err := client.Trends.Available()

	return err
}

func (instance *Instance) restoreStreams() error {
	// Init streams from database
	streams, err := instance.storage.GetActiveStreams()
	if err != nil {
		log.Error("Streams retrieving is failed. ", err)

		return err
	}
	for _, stream := range streams {
		instance.AddStream(stream)
	}

	return nil
}
