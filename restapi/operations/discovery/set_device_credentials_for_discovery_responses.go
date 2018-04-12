// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"casa/src/server/models"
)

// SetDeviceCredentialsForDiscoveryOKCode is the HTTP code returned for type SetDeviceCredentialsForDiscoveryOK
const SetDeviceCredentialsForDiscoveryOKCode int = 200

/*SetDeviceCredentialsForDiscoveryOK ok

swagger:response setDeviceCredentialsForDiscoveryOK
*/
type SetDeviceCredentialsForDiscoveryOK struct {
}

// NewSetDeviceCredentialsForDiscoveryOK creates SetDeviceCredentialsForDiscoveryOK with default headers values
func NewSetDeviceCredentialsForDiscoveryOK() *SetDeviceCredentialsForDiscoveryOK {
	return &SetDeviceCredentialsForDiscoveryOK{}
}

// WriteResponse to the client
func (o *SetDeviceCredentialsForDiscoveryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// SetDeviceCredentialsForDiscoveryUnauthorizedCode is the HTTP code returned for type SetDeviceCredentialsForDiscoveryUnauthorized
const SetDeviceCredentialsForDiscoveryUnauthorizedCode int = 401

/*SetDeviceCredentialsForDiscoveryUnauthorized Unauthorized

swagger:response setDeviceCredentialsForDiscoveryUnauthorized
*/
type SetDeviceCredentialsForDiscoveryUnauthorized struct {
}

// NewSetDeviceCredentialsForDiscoveryUnauthorized creates SetDeviceCredentialsForDiscoveryUnauthorized with default headers values
func NewSetDeviceCredentialsForDiscoveryUnauthorized() *SetDeviceCredentialsForDiscoveryUnauthorized {
	return &SetDeviceCredentialsForDiscoveryUnauthorized{}
}

// WriteResponse to the client
func (o *SetDeviceCredentialsForDiscoveryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// SetDeviceCredentialsForDiscoveryNotFoundCode is the HTTP code returned for type SetDeviceCredentialsForDiscoveryNotFound
const SetDeviceCredentialsForDiscoveryNotFoundCode int = 404

/*SetDeviceCredentialsForDiscoveryNotFound Not found

swagger:response setDeviceCredentialsForDiscoveryNotFound
*/
type SetDeviceCredentialsForDiscoveryNotFound struct {
}

// NewSetDeviceCredentialsForDiscoveryNotFound creates SetDeviceCredentialsForDiscoveryNotFound with default headers values
func NewSetDeviceCredentialsForDiscoveryNotFound() *SetDeviceCredentialsForDiscoveryNotFound {
	return &SetDeviceCredentialsForDiscoveryNotFound{}
}

// WriteResponse to the client
func (o *SetDeviceCredentialsForDiscoveryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

/*SetDeviceCredentialsForDiscoveryDefault error

swagger:response setDeviceCredentialsForDiscoveryDefault
*/
type SetDeviceCredentialsForDiscoveryDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSetDeviceCredentialsForDiscoveryDefault creates SetDeviceCredentialsForDiscoveryDefault with default headers values
func NewSetDeviceCredentialsForDiscoveryDefault(code int) *SetDeviceCredentialsForDiscoveryDefault {
	if code <= 0 {
		code = 500
	}

	return &SetDeviceCredentialsForDiscoveryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the set device credentials for discovery default response
func (o *SetDeviceCredentialsForDiscoveryDefault) WithStatusCode(code int) *SetDeviceCredentialsForDiscoveryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the set device credentials for discovery default response
func (o *SetDeviceCredentialsForDiscoveryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the set device credentials for discovery default response
func (o *SetDeviceCredentialsForDiscoveryDefault) WithPayload(payload *models.Error) *SetDeviceCredentialsForDiscoveryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the set device credentials for discovery default response
func (o *SetDeviceCredentialsForDiscoveryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SetDeviceCredentialsForDiscoveryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}