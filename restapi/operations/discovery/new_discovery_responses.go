// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"casa/src/server/models"
)

// NewDiscoveryOKCode is the HTTP code returned for type NewDiscoveryOK
const NewDiscoveryOKCode int = 200

/*NewDiscoveryOK ok

swagger:response newDiscoveryOK
*/
type NewDiscoveryOK struct {
}

// NewNewDiscoveryOK creates NewDiscoveryOK with default headers values
func NewNewDiscoveryOK() *NewDiscoveryOK {
	return &NewDiscoveryOK{}
}

// WriteResponse to the client
func (o *NewDiscoveryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*NewDiscoveryDefault error

swagger:response newDiscoveryDefault
*/
type NewDiscoveryDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNewDiscoveryDefault creates NewDiscoveryDefault with default headers values
func NewNewDiscoveryDefault(code int) *NewDiscoveryDefault {
	if code <= 0 {
		code = 500
	}

	return &NewDiscoveryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the new discovery default response
func (o *NewDiscoveryDefault) WithStatusCode(code int) *NewDiscoveryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the new discovery default response
func (o *NewDiscoveryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the new discovery default response
func (o *NewDiscoveryDefault) WithPayload(payload *models.Error) *NewDiscoveryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new discovery default response
func (o *NewDiscoveryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewDiscoveryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
