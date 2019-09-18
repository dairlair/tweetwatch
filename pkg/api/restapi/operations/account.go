// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/dairlair/tweetwatch/pkg/api/models"
)

// AccountHandlerFunc turns a function with the right signature into a account handler
type AccountHandlerFunc func(AccountParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AccountHandlerFunc) Handle(params AccountParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AccountHandler interface for that can handle valid account params
type AccountHandler interface {
	Handle(AccountParams, *models.User) middleware.Responder
}

// NewAccount creates a new http.Handler for the account operation
func NewAccount(ctx *middleware.Context, handler AccountHandler) *Account {
	return &Account{Context: ctx, Handler: handler}
}

/*Account swagger:route GET /account account

Account account API

*/
type Account struct {
	Context *middleware.Context
	Handler AccountHandler
}

func (o *Account) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAccountParams()

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
