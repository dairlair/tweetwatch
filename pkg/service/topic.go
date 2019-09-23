package service

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func (service *Service) CreateTopicHandler(params operations.CreateTopicParams, user *models.UserResponse) middleware.Responder {

	topic := entity.Topic{
		UserID:    *user.ID,
		Name:      *params.Topic.Name,
		Tracks:    params.Topic.Tracks,
	}

	createdTopic, err := service.storage.AddTopic(&topic)
	if err != nil {
		payload := models.ErrorResponse{Message: swag.String(fmt.Sprintf("%s", err))}
		return operations.NewCreateTopicDefault(422).WithPayload(&payload)
	}

	if createdTopic == nil {
		payload := models.ErrorResponse{Message: swag.String("Topic not created with unknown reason")}
		return operations.NewCreateTopicDefault(422).WithPayload(&payload)
	}

	payload := models.Topic{
		ID:    swag.Int64(createdTopic.GetID()),
		Name:  swag.String(createdTopic.GetName()),
		Tracks: createdTopic.GetTracks(),
		CreatedAt: swag.String(createdTopic.GetCreatedAt()),
	}
	return operations.NewCreateTopicOK().WithPayload(&payload)
}
