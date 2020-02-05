package void

import (
	"fmt"
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
func (instance *Instance) addStream(stream entity.StreamInterface) {
}

// AddStreams adds desired stream to the current instance of twitterclient
func (instance *Instance) AddStreams(streams []entity.StreamInterface) {
}

// DeleteStreams removes desired stream to the current instance of twitterclient
func (instance *Instance) DeleteStreams([]int64) {
}

// GetStreams returns all the streams from the current instance of twitterclient
func (instance *Instance) GetStreams() map[int64]entity.StreamInterface {
	return make(map[int64]entity.StreamInterface, 0)
}

// Watch starts watching
func (instance *Instance) Watch(output chan entity.TweetStreamsInterface) error {
	go func() {
		interval := 100 * time.Millisecond
		for i := 0; ; i++ {
			tweet := entity.Tweet{
				ID:              1,
				TwitterID:       time.Now().Unix() + int64(i),
				TwitterUserID:   int64(time.Now().Second() + i),
				TwitterUsername: "dairlair",
				FullText:        fmt.Sprintf("Just a fake tweet from void #%d", i),
				CreatedAt:       time.Now(),
			}
			tweetStreams := entity.NewTweetStreams(&tweet, entity.StreamsMapToSlice(instance.GetStreams()))
			<-time.After(interval)
			select {
			case output <- tweetStreams:
			default:
				log.Errorf("Can not send tweet and streams to output")
			}
		}
	}()
	return nil
}

// Unwatch stops watching
func (instance *Instance) Unwatch() {
}
