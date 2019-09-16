package gotwitter

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"

	log "github.com/sirupsen/logrus"
)

// Instance structure is used to store the server's state
type Instance struct {
	config twitterclient.Config

	// Currently watched streams
	streams map[int64]entity.StreamInterface
	// We will send to this channel all found tweets with associated streams
	output chan entity.TweetStreamsInterface
	// Internal resources
	client  *twitter.Client
	source  *twitter.Stream
}

// NewInstance creates new twitter instance scrapper
func NewInstance(config twitterclient.Config) twitterclient.Interface {
	log.Infof("Twitter: consumer key=%s, consumer_secret=%s, access token=%s, access secret=%s",
		config.TwitterConsumerKey,
		config.TwitterConsumerSecret,
		config.TwitterAccessToken,
		config.TwitterAccessSecret,
	)

	return &Instance{
		config:  config,
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

	return nil
}

// AddStream adds desired stream to the current instance of twitterclient
func (instance *Instance) AddStream(stream entity.StreamInterface) {
	if stream.GetID() < 1 {
		log.Errorf("stream without id can not be added")
		return
	}
	instance.streams[stream.GetID()] = stream
}

// GetStreams returns all the streams from the current instance of twitterclient
func (instance *Instance) GetStreams() map[int64]entity.StreamInterface {
	return instance.streams
}

// Watch starts watching
func (instance *Instance) Watch(output chan entity.TweetStreamsInterface) error {
	instance.output = output
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
	log.Infof("Tweet: %s\n", tweet.IDStr)
	log.Infof("%v\n\n", tweet)
	instance.processTweet(createTweetEntity(tweet))
}



func createTweetEntity(tweet *twitter.Tweet) entity.TweetInterface {
	var fullText string
	if tweet.ExtendedTweet != nil {
		fullText = tweet.ExtendedTweet.FullText
	} else {
		fullText = tweet.Text
	}
	return &entity.Tweet{
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