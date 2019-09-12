package gotwitter

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/splitter"
	"github.com/dairlair/tweetwatch/pkg/splitter/providers/substring"
	log "github.com/sirupsen/logrus"
)

// @TODO This function should be replaced to the injected dependency using.
func getSplitter() splitter.Interface {
	return substring.NewInstance()
}

// processTweet should find which streams are source of this tweet
func (instance *Instance) processTweet(tweet entity.TweetInterface) {
	log.Infof("Process tweet: %v\n", tweet)

	tweetStreams := getSplitter().Split(tweet, entity.StreamsMapToSlice(instance.GetStreams()))

	select {
	case instance.output <- tweetStreams:
	default:
		log.Errorf("Can not send tweet and streams to output. Tweet will be ignored.")
	}
}
