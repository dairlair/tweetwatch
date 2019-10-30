package storage

import (
	. "github.com/dairlair/tweetwatch/pkg/entity"
)

// Interface must be implemented by DBMS based storage.
type Interface interface {
	AddTopic(TopicInterface) (result TopicInterface, err error)
	UpdateTopic(TopicInterface) (result TopicInterface, err error)
	GetUserTopics(userID int64) (result []TopicInterface, err error)
	GetTopic(topicID int64) (topic TopicInterface, err error)
	GetTopicStreams(topicID int64) (streams []StreamInterface, err error)
	GetTopicTweets(topicID int64) (tweets []TweetInterface, err error)
	GetActiveStreams() (streams []StreamInterface, err error)
	DeleteTopic(streamID int64) error
	AddTweetStreams(tweetStreams TweetStreamsInterface) (id int64, err error)
	SignUp(email string, password string) (id *int64, err error)
	Login(email string, password string) (id *int64, err error)
	AddStream(streamInterface StreamInterface) (result StreamInterface, err error)
	UpdateStream(streamInterface StreamInterface) (result StreamInterface, err error)
	DeleteStream(streamID int64) error
}
