package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (service *Service) GetStatusHandler(params operations.GetStatusParams, user *models.User) middleware.Responder {
	return operations.NewGetStatusOK().WithPayload(user)
}