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
	stream := streamEntityFromModel(params.TopicID, params.Stream, user)

	createdStream, err := service.storage.AddStream(&stream)
	if err != nil {
		payload := models.DefaultError{Message: swag.String(fmt.Sprintf("%s", err))}
		return operations.NewCreateTopicDefault(422).WithPayload(&payload)
	}

	if createdStream == nil {
		payload := models.DefaultError{Message: swag.String("Stream not created with unknown reason")}
		return operations.NewCreateStreamDefault(422).WithPayload(&payload)
	}

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

func streamEntityFromModel(topicID int64, model *models.CreateStream, user *models.User) entity.Stream {
	stream := entity.Stream{
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
