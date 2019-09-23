package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (service *Service) CreateTopicHandler(params operations.CreateTopicParams, user *models.UserResponse) middleware.Responder {
	var id int64 = 1
	payload := models.Topic{
		ID:    &id,
		Name:  params.Topic.Name,
		Track: params.Topic.Track,
	}
	return operations.NewCreateTopicOK().WithPayload(&payload)
}
