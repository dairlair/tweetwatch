package service

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func (service *Service) CreateStreamHandler(params operations.CreateStreamParams, user *models.User) middleware.Responder {
	stream := streamEntityFromModel(params.TopicID, 0, params.Stream, user)

	createdStream, err := service.storage.AddStream(&stream)
	if err != nil {
		payload := models.DefaultError{Message: swag.String(fmt.Sprintf("%s", err))}
		return operations.NewCreateTopicDefault(422).WithPayload(&payload)
	}

	if createdStream == nil {
		payload := models.DefaultError{Message: swag.String("Stream not created with unknown reason")}
		return operations.NewCreateStreamDefault(422).WithPayload(&payload)
	}

	// @twitterclient: Add this stream to watching
	service.addStreamsToWatching([]entity.StreamInterface{createdStream})

	// @TODO Add stream watching if the topic is active
	//// Start watching created streams
	//if createdStream.GetIsActive() {
	//	service.addStreamsToWatching(createdTopic.GetStreams())
	//}

	payload := streamModelFromEntity(createdStream)
	return operations.NewCreateStreamOK().WithPayload(&payload)
}

func (service *Service) GetStreamsHandler(params operations.GetStreamsParams, user *models.User) middleware.Responder {
	streams, err := service.storage.GetTopicStreams(params.TopicID)

	if err != nil {
		payload := models.DefaultError{Message: swag.String(fmt.Sprint(err))}
		return operations.NewGetStreamsDefault(500).WithPayload(&payload)
	}

	var payload []*models.Stream
	for _, stream := range streams {
		model := streamModelFromEntity(stream)
		payload = append(payload, &model)
	}

	return operations.NewGetStreamsOK().WithPayload(payload)
}

func (service *Service) UpdateStreamHandler(params operations.UpdateStreamParams, user *models.User) middleware.Responder {
	stream := streamEntityFromModel(params.TopicID, params.StreamID, params.Stream, user)
	updatedStream, err := service.storage.UpdateStream(&stream)
	if err != nil {
		payload := models.DefaultError{Message: swag.String(fmt.Sprint(err))}
		return operations.NewUpdateStreamDefault(500).WithPayload(&payload)
	}
	if updatedStream == nil {
		payload := models.DefaultError{Message: swag.String("Stream not updated with unknown reason")}
		return operations.NewUpdateStreamDefault(500).WithPayload(&payload)
	}

	// @twitterclient: Update this stream in watching
	service.deleteStreamsFromWatching([]int64{params.StreamID})
	service.addStreamsToWatching([]entity.StreamInterface{updatedStream})

	payload := streamModelFromEntity(updatedStream)
	return operations.NewUpdateStreamOK().WithPayload(&payload)
}

func (service *Service) DeleteStreamHandler(params operations.DeleteStreamParams, user *models.User) middleware.Responder {
	err := service.storage.DeleteStream(params.StreamID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			payload := models.DefaultError{Message: swag.String("Not found")}
			return operations.NewDeleteStreamDefault(404).WithPayload(&payload)
		}
		payload := models.DefaultError{Message: swag.String(fmt.Sprint(err))}
		return operations.NewDeleteStreamDefault(500).WithPayload(&payload)
	}

	// @twitterclient: Remove this stream from watching
	service.deleteStreamsFromWatching([]int64{params.StreamID})

	payload := models.DefaultSuccess{Message:swag.String("Stream deleted successfully")}
	return operations.NewDeleteStreamOK().WithPayload(&payload)
}

func streamEntityFromModel(topicID int64, streamID int64, model *models.CreateStream, user *models.User) entity.Stream {
	stream := entity.Stream{
		ID: streamID,
		TopicID: topicID,
		Track:   *model.Track,
	}

	return stream
}

func streamModelFromEntity(entity entity.StreamInterface) models.Stream {
	stream := models.Stream{
		ID:        swag.Int64(entity.GetID()),
		Track:     swag.String(entity.GetTrack()),
		CreatedAt: swag.String(entity.GetCreatedAt().Format("2006-01-02T15:04:05-0700")),
	}

	return stream
}
