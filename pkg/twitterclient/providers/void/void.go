package void

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	log "github.com/sirupsen/logrus"
	"time"
)

// Instance structure is used to store the server's state
type Instance struct {
	streams map[int64]entity.StreamInterface
}

// NewInstance creates new twitter scrapper (void)
func NewInstance(_ twitterclient.Config) twitterclient.Interface {
	return &Instance{
		streams: make(map[int64]entity.StreamInterface),
	}
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
func (instance *Instance) Watch(output chan entity.TweetStreamsInterface) error {
	go func() {
		tweet := entity.Tweet{
			ID:            1,
			TwitterID:     9381,
			TwitterUserID: 5234,
			FullText:      "Just a fake tweet from void",
			CreatedAt:     time.Now().Format("20060102150405"),
		}
		tweetStreams := entity.NewTweetStreams(&tweet, entity.StreamsMapToSlice(instance.GetStreams()))
		for i := 0; i < 5; i++ {
			<-time.After(time.Second)
			select {
			case output <- tweetStreams:
				log.Infof("Tweet with streams sent to output")
			default:
				log.Errorf("Can not send tweet and streams to output")
			}
		}
	} ()
	return nil
}

// Unwatch stops watching
func (instance *Instance) Unwatch() {
}