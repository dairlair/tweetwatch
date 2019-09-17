package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

// Interface must be implemented by postgres based storage or something else.
type Interface interface {
	AddStream(entity.StreamInterface) (stream entity.StreamInterface, err error)
	GetStreams() (streams []entity.StreamInterface, err error)
	AddTweetStreams(tweetStreams entity.TweetStreamsInterface) (id int64, err error)
	SignUp(email string, password string) (token string, err error)
	Login(email string, password string) (token string, err error)
}
