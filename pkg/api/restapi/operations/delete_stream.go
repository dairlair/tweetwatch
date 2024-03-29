// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/dairlair/tweetwatch/pkg/api/models"
)

// DeleteStreamHandlerFunc turns a function with the right signature into a delete stream handler
type DeleteStreamHandlerFunc func(DeleteStreamParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStreamHandlerFunc) Handle(params DeleteStreamParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteStreamHandler interface for that can handle valid delete stream params
type DeleteStreamHandler interface {
	Handle(DeleteStreamParams, *models.User) middleware.Responder
}

// NewDeleteStream creates a new http.Handler for the delete stream operation
func NewDeleteStream(ctx *middleware.Context, handler DeleteStreamHandler) *DeleteStream {
	return &DeleteStream{Context: ctx, Handler: handler}
}

/*DeleteStream swagger:route DELETE /topics/{topicId}/streams/{streamId} deleteStream

Delete desired stream by Topic ID and Stream ID

*/
type DeleteStream struct {
	Context *middleware.Context
	Handler DeleteStreamHandler
}

func (o *DeleteStream) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStreamParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
