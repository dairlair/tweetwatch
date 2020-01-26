package service

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

// BroadcasterInterface defines dependency which used for service configuration
// If service have a specified broadcaster then all processed tweets will be pushed to the broadcaster
type BroadcasterInterface interface {
	Broadcast(channel string, streamsInterface entity.TweetStreamsInterface)
}