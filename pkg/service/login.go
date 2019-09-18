package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
)

func (service *Service) isRegisteredAuth(user string, pass string) (*models.User, error) {
	log.Warnf("Login with credentials: %s:%s\n", user, pass)
	token, err := service.storage.Login(user, pass)
	if err != nil {
		return nil, err
	}
	log.Warnf("Logged in with token: %s\n", token)
	email := "email"
	password := strfmt.Password("password")
	return &models.User{
		Email:    &email,
		Password: &password,
	}, nil
}

func (service *Service) LoginHandler(params operations.LoginParams, user *models.User) middleware.Responder {
	log.Infof("LoginHandler with data %v, %v", params, user)
	token := "Some custom JWT token"
	id := "qwqwe"
	payload := models.Token{
		Token: &token,
		User:  &id,
	}
	return operations.NewLoginOK().WithPayload(&payload)
}
