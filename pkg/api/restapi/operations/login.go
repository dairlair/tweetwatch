// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/dairlair/tweetwatch/pkg/api/models"
)

// LoginHandlerFunc turns a function with the right signature into a login handler
type LoginHandlerFunc func(LoginParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginHandlerFunc) Handle(params LoginParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// LoginHandler interface for that can handle valid login params
type LoginHandler interface {
	Handle(LoginParams, *models.User) middleware.Responder
}

// NewLogin creates a new http.Handler for the login operation
func NewLogin(ctx *middleware.Context, handler LoginHandler) *Login {
	return &Login{Context: ctx, Handler: handler}
}

/*Login swagger:route POST /login login

Login login API

*/
type Login struct {
	Context *middleware.Context
	Handler LoginHandler
}

func (o *Login) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoginParams()

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
