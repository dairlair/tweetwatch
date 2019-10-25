package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Ensures then when we update topic - the service must do follows:
// If topic is inactive or deleted - remove all his streams from twitterclient
func TestService_UpdateTopicHandler(t *testing.T) {
	// Create models and entities for test
	userModel := models.User{
		Email: swag.String(""),
		ID:    swag.Int64(1),
		Token: swag.String(""),
	}
	var topicID int64 = 1
	topicModel := models.CreateTopic{
		IsActive: swag.Bool(false),
		Name:     swag.String("Tesla, Inc."),
	}
	topicEntity := topicEntityFromModel(&topicModel, &userModel)
	topicEntity.ID = topicID

	var streams []entity.StreamInterface
	streams = append(streams, &entity.Stream{ID: 1, TopicID: topicID, Track: "Model X"})

	// Storage mock must check two methods called once: UpdateTopic, GetTopicStreams
	storageMock := createStorageMock()
	storageMock.On("UpdateTopic", &topicEntity).Return(&topicEntity, nil).Once()
	storageMock.On("GetTopicStreams", topicID).Return(streams, nil).Once()

	twitterMock := createTwitterclientMock()
	twitterMock.On("Unwatch").Return().Once()
	twitterMock.On("Watch", mock.AnythingOfType("chan entity.TweetStreamsInterface")).Return(nil)
	twitterMock.On("DeleteStreams", entity.GetStreamIDs(streams)).Return().Once()

	srv := NewService(storageMock, twitterMock)
	srv.UpdateTopicHandler(operations.UpdateTopicParams{
		Topic:       &topicModel,
		TopicID:     topicID,
	}, &userModel)

	storageMock.AssertExpectations(t)
	twitterMock.AssertExpectations(t)
}

// If topic is active and non deleted - add his streams to twitterclient