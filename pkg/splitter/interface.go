// This package provides ability to clarify - from which stream/streams tweet was received.
// Actually, Twitter Enterprise API provides this functionality out of the box but for pretty penny only.
// Therefore, for academic purposes only, we have to do it on our own.
//
//
package splitter

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

type Interface func (tweetInterface entity.TweetInterface, stream []entity.StreamInterface)