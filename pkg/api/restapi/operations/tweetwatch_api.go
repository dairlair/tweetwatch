// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	models "github.com/dairlair/tweetwatch/pkg/api/models"
)

// NewTweetwatchAPI creates a new Tweetwatch instance
func NewTweetwatchAPI(spec *loads.Document) *TweetwatchAPI {
	return &TweetwatchAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		CreateTopicHandler: CreateTopicHandlerFunc(func(params CreateTopicParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation CreateTopic has not yet been implemented")
		}),
		GetUserTopicsHandler: GetUserTopicsHandlerFunc(func(params GetUserTopicsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation GetUserTopics has not yet been implemented")
		}),
		LoginHandler: LoginHandlerFunc(func(params LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation Login has not yet been implemented")
		}),
		SignupHandler: SignupHandlerFunc(func(params SignupParams) middleware.Responder {
			return middleware.NotImplemented("operation Signup has not yet been implemented")
		}),
		UpdateTopicHandler: UpdateTopicHandlerFunc(func(params UpdateTopicParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation UpdateTopic has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		JWTAuth: func(token string) (*models.User, error) {
			return nil, errors.NotImplemented("api key auth (JWT) Authorization from header param [Authorization] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*TweetwatchAPI the tweetwatch API */
type TweetwatchAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// JWTAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	JWTAuth func(string) (*models.User, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// CreateTopicHandler sets the operation handler for the create topic operation
	CreateTopicHandler CreateTopicHandler
	// GetUserTopicsHandler sets the operation handler for the get user topics operation
	GetUserTopicsHandler GetUserTopicsHandler
	// LoginHandler sets the operation handler for the login operation
	LoginHandler LoginHandler
	// SignupHandler sets the operation handler for the signup operation
	SignupHandler SignupHandler
	// UpdateTopicHandler sets the operation handler for the update topic operation
	UpdateTopicHandler UpdateTopicHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TweetwatchAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TweetwatchAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TweetwatchAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TweetwatchAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TweetwatchAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TweetwatchAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TweetwatchAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TweetwatchAPI
func (o *TweetwatchAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.JWTAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.CreateTopicHandler == nil {
		unregistered = append(unregistered, "CreateTopicHandler")
	}

	if o.GetUserTopicsHandler == nil {
		unregistered = append(unregistered, "GetUserTopicsHandler")
	}

	if o.LoginHandler == nil {
		unregistered = append(unregistered, "LoginHandler")
	}

	if o.SignupHandler == nil {
		unregistered = append(unregistered, "SignupHandler")
	}

	if o.UpdateTopicHandler == nil {
		unregistered = append(unregistered, "UpdateTopicHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TweetwatchAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TweetwatchAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {

		case "JWT":

			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.JWTAuth(token)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *TweetwatchAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *TweetwatchAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TweetwatchAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TweetwatchAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the tweetwatch API
func (o *TweetwatchAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TweetwatchAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/topics"] = NewCreateTopic(o.context, o.CreateTopicHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/topics"] = NewGetUserTopics(o.context, o.GetUserTopicsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = NewLogin(o.context, o.LoginHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/signup"] = NewSignup(o.context, o.SignupHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/topics/{topicId}"] = NewUpdateTopic(o.context, o.UpdateTopicHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TweetwatchAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TweetwatchAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TweetwatchAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TweetwatchAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
