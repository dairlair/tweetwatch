package service

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func (service *Service) CreateTopicHandler(params operations.CreateTopicParams, user *models.User) middleware.Responder {
	topic := topicEntityFromModel(params.Topic, user)

	createdTopic, err := service.storage.AddTopic(&topic)
	if err != nil {
		payload := models.DefaultError{Message: swag.String(fmt.Sprintf("%s", err))}
		return operations.NewCreateTopicDefault(422).WithPayload(&payload)
	}

	if createdTopic == nil {
		payload := models.DefaultError{Message: swag.String("Topic not created with unknown reason")}
		return operations.NewCreateTopicDefault(422).WithPayload(&payload)
	}

	payload := topicModelFromEntity(createdTopic)
	return operations.NewCreateTopicOK().WithPayload(&payload)
}

func (service *Service) GetUserTopicsHandler(params operations.GetUserTopicsParams, user *models.User) middleware.Responder {
	topics, err := service.storage.GetUserTopics(*user.ID)

	if err != nil {
		payload := models.DefaultError{Message: swag.String("Topics not retrieved with unknown reason")}
		return operations.NewCreateTopicDefault(422).WithPayload(&payload)
	}

	var payload []*models.Topic
	for _, topic := range topics {
		model := topicModelFromEntity(topic)
		payload = append(payload, &model)
	}

	return operations.NewGetUserTopicsOK().WithPayload(payload)
}

func (service *Service) UpdateTopicHandler(params operations.UpdateTopicParams, user *models.User) middleware.Responder {
	topic := topicEntityFromModel(params.Topic, user)
	topic.ID = params.TopicID

	// Run update topic in storage
	updatedTopic, err := service.storage.UpdateTopic(&topic)
	if err != nil {
		payload := models.DefaultError{Message: swag.String(fmt.Sprintf("Topic not updated: %s", err))}
		return operations.NewUpdateTopicDefault(422).WithPayload(&payload)
	}

	payload := topicModelFromEntity(updatedTopic)
	return operations.NewUpdateTopicOK().WithPayload(&payload)
}

func topicEntityFromModel(model *models.CreateTopic, user *models.User) entity.Topic {
	topic := entity.Topic{
		UserID:  *user.ID,
		Name:    *model.Name,
		IsActive: *model.IsActive,
	}
	return topic
}

func topicModelFromEntity(entity entity.TopicInterface) models.Topic {
	model := models.Topic{
		ID:        swag.Int64(entity.GetID()),
		Name:      swag.String(entity.GetName()),
		CreatedAt: swag.String(entity.GetCreatedAt().Format("2006-01-02T15:04:05-0700")),
		IsActive:  swag.Bool(entity.GetIsActive()),
	}

	return model
}
