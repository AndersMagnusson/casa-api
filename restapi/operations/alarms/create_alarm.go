// Code generated by go-swagger; DO NOT EDIT.

package alarms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateAlarmHandlerFunc turns a function with the right signature into a create alarm handler
type CreateAlarmHandlerFunc func(CreateAlarmParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAlarmHandlerFunc) Handle(params CreateAlarmParams) middleware.Responder {
	return fn(params)
}

// CreateAlarmHandler interface for that can handle valid create alarm params
type CreateAlarmHandler interface {
	Handle(CreateAlarmParams) middleware.Responder
}

// NewCreateAlarm creates a new http.Handler for the create alarm operation
func NewCreateAlarm(ctx *middleware.Context, handler CreateAlarmHandler) *CreateAlarm {
	return &CreateAlarm{Context: ctx, Handler: handler}
}

/*CreateAlarm swagger:route POST /alarms alarms createAlarm

Create/Update the alarm

*/
type CreateAlarm struct {
	Context *middleware.Context
	Handler CreateAlarmHandler
}

func (o *CreateAlarm) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAlarmParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
