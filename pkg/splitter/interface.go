// This package provides ability to clarify - from which stream/streams tweet was received.
// Actually, Twitter Enterprise API provides this functionality out of the box but for pretty penny only.
// Therefore, for academic purposes only, we have to do it on our own.
package splitter

import (
	. "github.com/dairlair/tweetwatch/pkg/entity"
)

// Interface is used to inject some certain splitters as dependency.
type Interface interface {
	// Split accept tweet and all active streams. After that Split tries to recognize source stream for the tweet.
	// As instance:
	//   We have currently active streams: ["Tesla", "Trump", "Golang"] and with these streams we got a tweet:
	//   ""
	Split (tweetInterface TweetInterface, streams []StreamInterface) TweetStreamsInterface
}