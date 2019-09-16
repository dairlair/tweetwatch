package gotwitter

import (
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	"os"
	"testing"

	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewInstance_Successful(t *testing.T) {
	cfg := twitterclient.Config{}
	instance := NewInstance(cfg)
	assert.IsType(t, &Instance{}, instance, "Object must have type twitterclient.Instance")
}

func TestStart_Successful(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for twitterclient successful start")
	}
	cfg := twitterclient.Config{
		TwitterConsumerKey:    os.Getenv("TWITWATCH_TEST_TWITTER_CONSUMER_KEY"),
		TwitterConsumerSecret: os.Getenv("TWITWATCH_TEST_TWITTER_CONSUMER_SECRET"),
		TwitterAccessToken:    os.Getenv("TWITWATCH_TEST_TWITTER_ACCESS_TOKEN"),
		TwitterAccessSecret:   os.Getenv("TWITWATCH_TEST_TWITTER_ACCESS_SECRET"),
	}
	instance := NewInstance(cfg)

	err := instance.Start()
	assert.Nil(t, err)
}

func TestStart_AuthFailed(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for twitterclient auth failed")
	}
	cfg := twitterclient.Config{
		TwitterConsumerKey:    "a",
		TwitterConsumerSecret: "b",
		TwitterAccessToken:    "c",
		TwitterAccessSecret:   "d",
	}
	instance := NewInstance(cfg)
	err := instance.Start()
	assert.NotNil(t, err, "Error must be not nil when try to start with wrong credentials")
}

func TestAddStream_Successful(t *testing.T) {
	cfg := twitterclient.Config{}
	instance := NewInstance(cfg)
	stream := entity.Stream{}
	instance.AddStream(&stream)
}

func TestGetStreams_Successful(t *testing.T) {
	cfg := twitterclient.Config{}
	instance := NewInstance(cfg)

	streams := map[int64]entity.Stream{
		1: {ID: 1, Track: "Tesla"},
		2: {ID: 2, Track: "Apple"},
		3: {ID: 3, Track: "Microsoft"},
	}

	for _, stream := range streams {
		instance.AddStream(&stream)
	}

	assert.Equal(t, len(streams), len(instance.GetStreams()))
	for _, stream := range instance.GetStreams() {
		assert.EqualValues(t, streams[stream.GetID()].ID, stream.GetID())
		assert.EqualValues(t, streams[stream.GetID()].Track, stream.GetTrack())
	}
}