package service

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)


func (service *Service) GetTopicTweetsHandler(params operations.GetTopicTweetsParams, user *models.User) middleware.Responder {
	tweets, err := service.storage.GetTopicTweets(params.TopicID)

	if err != nil {
		payload := models.DefaultError{Message: swag.String(fmt.Sprint(err))}
		return operations.NewGetTopicTweetsDefault(500).WithPayload(&payload)
	}

	var payload []*models.Tweet
	for _, tweet := range tweets {
		model := tweetModelFromEntity(tweet)
		payload = append(payload, &model)
	}

	return operations.NewGetTopicTweetsOK().WithPayload(payload)
}

func tweetModelFromEntity(entity entity.TweetInterface) models.Tweet {
	return models.Tweet{
		CreatedAt:       swag.String(entity.GetCreatedAt().Format("2006-01-02T15:04:05-0700")),
		FullText:        swag.String(entity.GetFullText()),
		ID:              swag.Int64(entity.GetID()),
		TwitterID:       swag.Int64(entity.GetTwitterID()),
		TwitterUserID:   swag.Int64(entity.GetTwitterUserID()),
		TwitterUsername: swag.String(entity.GetTwitterUsername()),
	}
}