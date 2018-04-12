// Code generated by go-swagger; DO NOT EDIT.

package alarms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"casa/src/server/models"
)

// ListAlarmLogsOKCode is the HTTP code returned for type ListAlarmLogsOK
const ListAlarmLogsOKCode int = 200

/*ListAlarmLogsOK Ok

swagger:response listAlarmLogsOK
*/
type ListAlarmLogsOK struct {

	/*
	  In: Body
	*/
	Payload models.ToggleAlarms `json:"body,omitempty"`
}

// NewListAlarmLogsOK creates ListAlarmLogsOK with default headers values
func NewListAlarmLogsOK() *ListAlarmLogsOK {
	return &ListAlarmLogsOK{}
}

// WithPayload adds the payload to the list alarm logs o k response
func (o *ListAlarmLogsOK) WithPayload(payload models.ToggleAlarms) *ListAlarmLogsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list alarm logs o k response
func (o *ListAlarmLogsOK) SetPayload(payload models.ToggleAlarms) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAlarmLogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.ToggleAlarms, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ListAlarmLogsDefault error

swagger:response listAlarmLogsDefault
*/
type ListAlarmLogsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListAlarmLogsDefault creates ListAlarmLogsDefault with default headers values
func NewListAlarmLogsDefault(code int) *ListAlarmLogsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListAlarmLogsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list alarm logs default response
func (o *ListAlarmLogsDefault) WithStatusCode(code int) *ListAlarmLogsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list alarm logs default response
func (o *ListAlarmLogsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list alarm logs default response
func (o *ListAlarmLogsDefault) WithPayload(payload *models.Error) *ListAlarmLogsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list alarm logs default response
func (o *ListAlarmLogsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAlarmLogsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
