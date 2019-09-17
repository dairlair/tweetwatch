package service

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

func (service *Service) login(user string, pass string) (interface{}, error) {
	log.Warnf("Login with credentials: %s:%s\n", user, pass)
	token, err := service.storage.Login(user, pass)
	if err != nil {
		return nil, err
	}
	log.Warnf("Logged in with token: %s\n", token)
	return nil, nil
	//return User
	//token, err := service.storage.SignUpHandler(*params., params.User.Password.String())
	//
	//if err != nil {
	//	payload := models.ErrorResponse{Message: swag.String("Email already taken")}
	//	return operations.NewSignupDefault(422).WithPayload(&payload)
	//}
	//
	//message := fmt.Sprintf("User [%s] registered with token [%s]", *params.User.Email, token)
	//payload := models.GeneralResponse{
	//	Message: &message,
	//}

}

func (service *Service) LoginHandler(params operations.LoginParams, data interface{}) middleware.Responder {
	message := fmt.Sprintf("LoginHandler with data %v", data)
	payload := models.GeneralResponse{
		Message: &message,
	}
	return operations.NewSignupOK().WithPayload(&payload)
}
