package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (service *Service) isRegisteredAuth(user string, pass string) (*models.UserResponse, error) {
	id, err := service.storage.Login(user, pass)
	if err != nil {
		return nil, err
	}
	email := "email"
	return &models.UserResponse{
		ID:    id,
		Email: &email,
	}, nil
}

func (service *Service) LoginHandler(_ operations.LoginParams, user *models.UserResponse) middleware.Responder {
	payload := models.UserResponse{
		Email: user.Email,
		ID:    user.ID,
	}
	return operations.NewLoginOK().WithPayload(&payload)
}
