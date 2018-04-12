// Code generated by go-swagger; DO NOT EDIT.

package alarms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"casa/src/server/models"
)

// DeleteAlarmNoContentCode is the HTTP code returned for type DeleteAlarmNoContent
const DeleteAlarmNoContentCode int = 204

/*DeleteAlarmNoContent Ok

swagger:response deleteAlarmNoContent
*/
type DeleteAlarmNoContent struct {
}

// NewDeleteAlarmNoContent creates DeleteAlarmNoContent with default headers values
func NewDeleteAlarmNoContent() *DeleteAlarmNoContent {
	return &DeleteAlarmNoContent{}
}

// WriteResponse to the client
func (o *DeleteAlarmNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteAlarmDefault error

swagger:response deleteAlarmDefault
*/
type DeleteAlarmDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAlarmDefault creates DeleteAlarmDefault with default headers values
func NewDeleteAlarmDefault(code int) *DeleteAlarmDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteAlarmDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete alarm default response
func (o *DeleteAlarmDefault) WithStatusCode(code int) *DeleteAlarmDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete alarm default response
func (o *DeleteAlarmDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete alarm default response
func (o *DeleteAlarmDefault) WithPayload(payload *models.Error) *DeleteAlarmDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete alarm default response
func (o *DeleteAlarmDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAlarmDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}