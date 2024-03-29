// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/dairlair/tweetwatch/pkg/api/models"
)

// NewCreateTopicParams creates a new CreateTopicParams object
// no default values defined in spec.
func NewCreateTopicParams() CreateTopicParams {

	return CreateTopicParams{}
}

// CreateTopicParams contains all the bound params for the create topic operation
// typically these are obtained from a http.Request
//
// swagger:parameters createTopic
type CreateTopicParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*New Topic
	  Required: true
	  In: body
	*/
	Topic *models.CreateTopic
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateTopicParams() beforehand.
func (o *CreateTopicParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CreateTopic
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("topic", "body"))
			} else {
				res = append(res, errors.NewParseError("topic", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Topic = &body
			}
		}
	} else {
		res = append(res, errors.Required("topic", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
