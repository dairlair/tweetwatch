package gotwitter

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/entity"
)

// processTweet should find which streams are source of this tweet
func (instance *Instance) processTweet(tweet entity.TweetInterface) {
	fmt.Printf("Process tweet: %v\n", tweet)
}
