// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"casa-api/models"
)

// AddAlertMethodNotAllowedCode is the HTTP code returned for type AddAlertMethodNotAllowed
const AddAlertMethodNotAllowedCode int = 405

/*AddAlertMethodNotAllowed Invalid input

swagger:response addAlertMethodNotAllowed
*/
type AddAlertMethodNotAllowed struct {
}

// NewAddAlertMethodNotAllowed creates AddAlertMethodNotAllowed with default headers values
func NewAddAlertMethodNotAllowed() *AddAlertMethodNotAllowed {
	return &AddAlertMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddAlertMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
}

/*AddAlertDefault error

swagger:response addAlertDefault
*/
type AddAlertDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAlertDefault creates AddAlertDefault with default headers values
func NewAddAlertDefault(code int) *AddAlertDefault {
	if code <= 0 {
		code = 500
	}

	return &AddAlertDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add alert default response
func (o *AddAlertDefault) WithStatusCode(code int) *AddAlertDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add alert default response
func (o *AddAlertDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add alert default response
func (o *AddAlertDefault) WithPayload(payload *models.Error) *AddAlertDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add alert default response
func (o *AddAlertDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAlertDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
