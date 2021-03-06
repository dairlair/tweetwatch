// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/dairlair/tweetwatch/pkg/api/models"
)

// DeleteStreamOKCode is the HTTP code returned for type DeleteStreamOK
const DeleteStreamOKCode int = 200

/*DeleteStreamOK Stream deleted

swagger:response deleteStreamOK
*/
type DeleteStreamOK struct {

	/*
	  In: Body
	*/
	Payload *models.DefaultSuccess `json:"body,omitempty"`
}

// NewDeleteStreamOK creates DeleteStreamOK with default headers values
func NewDeleteStreamOK() *DeleteStreamOK {

	return &DeleteStreamOK{}
}

// WithPayload adds the payload to the delete stream o k response
func (o *DeleteStreamOK) WithPayload(payload *models.DefaultSuccess) *DeleteStreamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete stream o k response
func (o *DeleteStreamOK) SetPayload(payload *models.DefaultSuccess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStreamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteStreamDefault Error

swagger:response deleteStreamDefault
*/
type DeleteStreamDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.DefaultError `json:"body,omitempty"`
}

// NewDeleteStreamDefault creates DeleteStreamDefault with default headers values
func NewDeleteStreamDefault(code int) *DeleteStreamDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteStreamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete stream default response
func (o *DeleteStreamDefault) WithStatusCode(code int) *DeleteStreamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete stream default response
func (o *DeleteStreamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete stream default response
func (o *DeleteStreamDefault) WithPayload(payload *models.DefaultError) *DeleteStreamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete stream default response
func (o *DeleteStreamDefault) SetPayload(payload *models.DefaultError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStreamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
