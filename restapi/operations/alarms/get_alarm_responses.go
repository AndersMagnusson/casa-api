// Code generated by go-swagger; DO NOT EDIT.

package alarms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"casa/src/server/models"
)

// GetAlarmOKCode is the HTTP code returned for type GetAlarmOK
const GetAlarmOKCode int = 200

/*GetAlarmOK Ok

swagger:response getAlarmOK
*/
type GetAlarmOK struct {

	/*
	  In: Body
	*/
	Payload *models.Alarm `json:"body,omitempty"`
}

// NewGetAlarmOK creates GetAlarmOK with default headers values
func NewGetAlarmOK() *GetAlarmOK {
	return &GetAlarmOK{}
}

// WithPayload adds the payload to the get alarm o k response
func (o *GetAlarmOK) WithPayload(payload *models.Alarm) *GetAlarmOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get alarm o k response
func (o *GetAlarmOK) SetPayload(payload *models.Alarm) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAlarmOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAlarmDefault error

swagger:response getAlarmDefault
*/
type GetAlarmDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAlarmDefault creates GetAlarmDefault with default headers values
func NewGetAlarmDefault(code int) *GetAlarmDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAlarmDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get alarm default response
func (o *GetAlarmDefault) WithStatusCode(code int) *GetAlarmDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get alarm default response
func (o *GetAlarmDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get alarm default response
func (o *GetAlarmDefault) WithPayload(payload *models.Error) *GetAlarmDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get alarm default response
func (o *GetAlarmDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAlarmDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
