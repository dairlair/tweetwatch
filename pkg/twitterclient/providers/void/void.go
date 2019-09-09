package void

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
)

// Instance structure is used to store the server's state
type Instance struct {
}

// NewInstance creates new twitter scrapper (void)
func NewInstance(_ twitterclient.Config, _ chan <- entity.TweetStreamsInterface) twitterclient.Interface {
	return &Instance{}
}

// Start function runs twitter client and validates credentials for Twitter Streaming API
func (instance *Instance) Start() error {
	return nil
}

// AddStream adds desired stream to the current instance of twitterclient
func (instance *Instance) AddStream(stream entity.StreamInterface) {
}

// GetStreams returns all the streams from the current instance of twitterclient
func (instance *Instance) GetStreams() map[int64]entity.StreamInterface {
	return make(map[int64]entity.StreamInterface, 0)
}

// Watch starts watching
func (instance *Instance) Watch() error {
	return nil
}

// Unwatch stops watching
func (instance *Instance) Unwatch() {
}