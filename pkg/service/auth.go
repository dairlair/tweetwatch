// See https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
package service

import (
	"errors"
	"github.com/dairlair/tweetwatch/pkg/api/models"
	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"
	"time"
)

type Claims struct {
	UserID int64 `json:"userId"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (service *Service) JWTAuth(token string) (*models.User, error) {
	log.Infof("JWTAuth with token: %s", token)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return service.jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	//if err != nil {
	//	if err == jwt.ErrSignatureInvalid {
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	if !tkn.Valid {
		// @TODO Add custom error handling
		// w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("invalid token")
	}
	return &models.User{
		Email: swag.String(claims.Email),
		ID:    swag.Int64(claims.UserID),
		Token: swag.String(token),
	}, nil
}

func (service *Service) LoginHandler(params operations.LoginParams) middleware.Responder {
	id, err := service.storage.Login(*params.User.Email, params.User.Password.String())
	if err != nil {
		payload := models.LoginUserInputError{Code: swag.String("INVALID_CREDENTIALS")}
		return operations.NewLoginForbidden().WithPayload(&payload)
	}

	token, err := service.createJwtToken(*id, *params.User.Email)
	if err != nil {
		payload := models.DefaultError{Message: swag.String("JWT Token not created")}
		return operations.NewLoginDefault(500).WithPayload(&payload)
	}

	payload := models.User{
		ID: id,
		Email: params.User.Email,
		Token: token,
	}
	return operations.NewLoginOK().WithPayload(&payload)
}

func (service *Service) SignUpHandler(params operations.SignupParams) middleware.Responder {

	id, err := service.storage.SignUp(*params.User.Email, params.User.Password.String())
	if err != nil {
		payload := models.SignupUserInputError{Code: swag.String("EMAIL_ALREADY_TAKEN")}
		return operations.NewSignupForbidden().WithPayload(&payload)
	}

	token, err := service.createJwtToken(*id, *params.User.Email)
	if err != nil {
		payload := models.DefaultError{Message: swag.String("JWT Token not created")}
		return operations.NewSignupDefault(500).WithPayload(&payload)
	}

	payload := models.User{
		ID: id,
		Email: params.User.Email,
		Token: token,
	}
	return operations.NewSignupOK().WithPayload(&payload)
}

func (service *Service) createJwtToken(userID int64, email string) (token *string, err error) {
	expirationTime := time.Now().AddDate(1, 0, 0)
	claims := &Claims{
		UserID: userID,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := jwtToken.SignedString(service.jwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}