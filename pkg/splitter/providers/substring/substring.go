// See https://stackoverflow.com/questions/24836044/case-insensitive-string-search-in-golang
package substring

import (
	. "github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/splitter"
	log "github.com/sirupsen/logrus"
	"strings"
)

type instance struct {
}

func NewInstance() splitter.Interface {
	return &instance{}
}

func (instance *instance) Split(tweet TweetInterface, streams []StreamInterface) TweetStreamsInterface {
	log.Infof("Trying to split tweet...\n")
	var matchedStreams []StreamInterface
	for _, stream := range streams {

		if strings.Contains(strings.ToLower(tweet.GetFullText()), strings.ToLower(stream.GetTrack())) {
			matchedStreams = append(matchedStreams, stream)
			//log.Warnf("Matched [%s] with [%s]", stream.GetTrack(), tweet.GetFullText())
		} else {
			//log.Warnf("Not matched [%s] with [%s]", stream.GetTrack(), tweet.GetFullText())
		}
	}

	if len(matchedStreams) == 0 {
		log.Warnf("Not matched [%s] with any stream", tweet.GetFullText())
	}

	if len(matchedStreams) > 1 {
		log.Warnf("Matched [%s] with multiple streams", tweet.GetFullText())
	}

	return NewTweetStreams(tweet, matchedStreams)
}
