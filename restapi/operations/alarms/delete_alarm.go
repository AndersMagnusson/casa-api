// Code generated by go-swagger; DO NOT EDIT.

package alarms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteAlarmHandlerFunc turns a function with the right signature into a delete alarm handler
type DeleteAlarmHandlerFunc func(DeleteAlarmParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAlarmHandlerFunc) Handle(params DeleteAlarmParams) middleware.Responder {
	return fn(params)
}

// DeleteAlarmHandler interface for that can handle valid delete alarm params
type DeleteAlarmHandler interface {
	Handle(DeleteAlarmParams) middleware.Responder
}

// NewDeleteAlarm creates a new http.Handler for the delete alarm operation
func NewDeleteAlarm(ctx *middleware.Context, handler DeleteAlarmHandler) *DeleteAlarm {
	return &DeleteAlarm{Context: ctx, Handler: handler}
}

/*DeleteAlarm swagger:route DELETE /alarms/{id} alarms deleteAlarm

Delete the alarm

*/
type DeleteAlarm struct {
	Context *middleware.Context
	Handler DeleteAlarmHandler
}

func (o *DeleteAlarm) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAlarmParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
