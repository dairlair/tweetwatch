package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

// Interface must be implemented by postgres based storage or something else.
type Interface interface {
	AddTopic(entity.TopicInterface) (result entity.TopicInterface, err error)
	UpdateTopic(entity.TopicInterface) (result entity.TopicInterface, deletedStreamIDs []int64, insertedStreams []entity.StreamInterface, err error)
	GetUserTopics(userId int64) (result []entity.TopicInterface, err error)
	GetStreams() (streams []entity.StreamInterface, err error)
	AddTweetStreams(tweetStreams entity.TweetStreamsInterface) (id int64, err error)
	SignUp(email string, password string) (id *int64, err error)
	Login(email string, password string) (id *int64, err error)
	AddStream(streamInterface entity.StreamInterface) (result entity.StreamInterface, err error)
}
