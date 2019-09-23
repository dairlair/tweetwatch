package service

import (
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

func (service *Service) SignUpHandler(params operations.SignupParams) middleware.Responder {

	id, err := service.storage.SignUp(*params.User.Email, params.User.Password.String())

	if err != nil {
		payload := models.ErrorResponse{Message: swag.String("Email already taken")}
		return operations.NewSignupDefault(422).WithPayload(&payload)
	}

	// message := fmt.Sprintf("User [%s] registered with token [%s]", *params.User.Email, token)
	payload := models.UserResponse{
		ID: id,
		Email: params.User.Email,
	}
	return operations.NewSignupOK().WithPayload(&payload)
}