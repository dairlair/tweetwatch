package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/entity"
	log "github.com/sirupsen/logrus"
)

const (
	BroadcastingServiceName     = "TweetWatch"
	BroadcastingEventTweetSaved = "TweetSaved"
)

// BroadcasterInterface defines dependency which used for service configuration
// If service have a specified broadcaster then all processed tweets will be pushed to the broadcaster
type BroadcasterInterface interface {
	Broadcast(channel string, data []byte)
}

func (service *Service) broadcast(tweetId int64, tweetStreams entity.TweetStreamsInterface) {
	if service.broadcaster == nil {
		return
	}

	tweetModel := tweetModelFromEntity(tweetStreams.GetTweet())
	tweetModel.ID = &tweetId
	savedTweet := models.SavedTweet{
		Tweet:    tweetModel,
	}
	savedTweet.Streams = make([]*models.Stream, 0)
	for _, stream := range tweetStreams.GetStreams() {
		streamModel := streamModelFromEntity(stream)
		savedTweet.Streams = append(savedTweet.Streams, &streamModel)
	}

	json, err := savedTweet.MarshalJSON()
	if err != nil {
		log.Errorf("Saved tweet marshalling failed: %s", err)
	}

	log.Fatalf("Broadcast JSON: %s", json)
	service.broadcaster.Broadcast(BroadcastingServiceName+ "." +BroadcastingEventTweetSaved, json)
}