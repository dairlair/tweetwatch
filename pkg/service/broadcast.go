package service

import (
	"encoding/json"
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/entity"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	broadcastingEventTweetSaved = "TweetSaved"
)

// BroadcasterInterface defines dependency which used for service configuration
// If service have a specified broadcaster then all processed tweets will be pushed to the broadcaster
type BroadcasterInterface interface {
	Broadcast(channel string, data []byte)
}

func (service *Service) broadcast(tweetID int64, tweetStreams entity.TweetStreamsInterface) {
	if service.broadcaster == nil {
		return
	}

	tweet := tweetStreams.GetTweet()
	bm := BroadcastMessage{
		OriginTime:     tweet.GetCreatedAt(),
		OriginID:       "twitter",
		OriginEntity:   "status",
		OriginEntityID: fmt.Sprintf("%d", tweet.GetTwitterID()),
		OriginText:     tweet.GetFullText(),
		OriginUserId:   fmt.Sprintf("%d", tweet.GetTwitterUserID()),
		OriginUsername: tweet.GetTwitterUsername(),
		Streams:        nil,
	}

	for _, stream := range tweetStreams.GetStreams() {
		bm.Streams = append(bm.Streams, BroadcastStream{
			StreamID: stream.GetID(),
			TopicID:  stream.GetTopicID(),
			Track:    stream.GetTrack(),
		})
	}

	js, err := json.Marshal(bm)
	if err != nil {
		log.Errorf("Saved tweet marshalling failed: %s", err)
	} else {
		service.broadcaster.Broadcast(broadcastingEventTweetSaved, js)
	}
}

// @TODO Move this structure definition into the Swagger Specification
type BroadcastStream struct {
	StreamID int64  `json:"streamId"`
	TopicID  int64  `json:"topicId"`
	Track    string `json:"track"`
}

// @TODO Move this structure definition into the Swagger Specification
type BroadcastMessage struct {
	OriginTime     time.Time         `json:"originTime,string"`
	OriginID       string            `json:"originId"`
	OriginEntity   string            `json:"originEntity"`
	OriginEntityID string            `json:"originEntityId"`
	OriginText     string            `json:"originText"`
	OriginUserId   string            `json:"originUserId"`
	OriginUsername string            `json:"originUsername"`
	Streams        []BroadcastStream `json:"streams"`
}
