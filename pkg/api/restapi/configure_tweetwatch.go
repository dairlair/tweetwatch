// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/dairlair/tweetwatch/pkg/api/restapi/operations"

	models "github.com/dairlair/tweetwatch/pkg/api/models"
)

//go:generate swagger generate server --target ../../api --name Tweetwatch --spec ../../../api/swagger-spec/tweetwatch-server.yml --principal models.User --exclude-main

func configureFlags(api *operations.TweetwatchAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TweetwatchAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	api.JWTAuth = func(token string) (*models.User, error) {
		return nil, errors.NotImplemented("api key auth (JWT) Authorization from header param [Authorization] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.CreateTopicHandler == nil {
		api.CreateTopicHandler = operations.CreateTopicHandlerFunc(func(params operations.CreateTopicParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation .CreateTopic has not yet been implemented")
		})
	}
	if api.GetUserTopicsHandler == nil {
		api.GetUserTopicsHandler = operations.GetUserTopicsHandlerFunc(func(params operations.GetUserTopicsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation .GetUserTopics has not yet been implemented")
		})
	}
	if api.LoginHandler == nil {
		api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation .Login has not yet been implemented")
		})
	}
	if api.SignupHandler == nil {
		api.SignupHandler = operations.SignupHandlerFunc(func(params operations.SignupParams) middleware.Responder {
			return middleware.NotImplemented("operation .Signup has not yet been implemented")
		})
	}
	if api.UpdateTopicHandler == nil {
		api.UpdateTopicHandler = operations.UpdateTopicHandlerFunc(func(params operations.UpdateTopicParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation .UpdateTopic has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
