package gotwitter

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/entity"
	log "github.com/sirupsen/logrus"
)

// processTweet should find which streams are source of this tweet
func (instance *Instance) processTweet(tweet entity.TweetInterface) {
	fmt.Printf("Process tweet: %v\n", tweet)

	// Now lets assume than every tweet belongs all streams, just for a test...
	tweetStreams := entity.NewTweetStreams(tweet, entity.StreamsMapToSlice(instance.GetStreams()))

	select {
	case instance.output <- tweetStreams:
		log.Infof("Tweet with streams sent to output")
	default:
		log.Errorf("Can not send tweet and streams to output. Tweet will be ignored.")
	}
}
