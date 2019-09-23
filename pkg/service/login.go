package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

func (service *Service) isRegisteredAuth(user string, pass string) (*models.UserResponse, error) {
	log.Warnf("Login with credentials: %s:%s\n", user, pass)
	id, err := service.storage.Login(user, pass)
	if err != nil {
		return nil, err
	}
	log.Warnf("Logged in with id: %d\n", id)
	email := "email"
	return &models.UserResponse{
		ID: id,
		Email:    &email,
	}, nil
}

func (service *Service) LoginHandler(params operations.LoginParams, user *models.UserResponse) middleware.Responder {
	log.Infof("LoginHandler with data %v, %v", params, user)
	token := "Some custom JWT token"
	id := "qwerty"
	payload := models.Token{
		Token: &token,
		User:  &id,
	}
	return operations.NewLoginOK().WithPayload(&payload)
}
