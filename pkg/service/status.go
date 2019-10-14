package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (service *Service) CheckStatusHandler(_ operations.CheckStatusParams, user *models.User) middleware.Responder {
	return operations.NewCheckStatusOK().WithPayload(user)
}