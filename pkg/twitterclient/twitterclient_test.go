package twitterclient

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInstance_Successfull(t *testing.T) {
	cfg := Config{}
	instance := NewInstance(cfg)
	assert.IsType(t, &Instance{}, instance, "Object must have type twitterclient.Instance")
}

func TestStart_Successfull(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for storage")
	}
	cfg := Config{
		TwitterConsumerKey:    os.Getenv("TWITWATCH_TEST_TWITTER_CONSUMER_KEY"),
		TwitterConsumerSecret: os.Getenv("TWITWATCH_TEST_TWITTER_CONSUMER_SECRET"),
		TwitterAccessToken:    os.Getenv("TWITWATCH_TEST_TWITTER_ACCESS_TOKEN"),
		TwitterAccessSecret:   os.Getenv("TWITWATCH_TEST_TWITTER_ACCESS_SECRET"),
	}
	instance := NewInstance(cfg)

	err := instance.Start()
	assert.Nil(t, err)
}
