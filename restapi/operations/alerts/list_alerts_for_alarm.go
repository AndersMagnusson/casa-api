// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListAlertsForAlarmHandlerFunc turns a function with the right signature into a list alerts for alarm handler
type ListAlertsForAlarmHandlerFunc func(ListAlertsForAlarmParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAlertsForAlarmHandlerFunc) Handle(params ListAlertsForAlarmParams) middleware.Responder {
	return fn(params)
}

// ListAlertsForAlarmHandler interface for that can handle valid list alerts for alarm params
type ListAlertsForAlarmHandler interface {
	Handle(ListAlertsForAlarmParams) middleware.Responder
}

// NewListAlertsForAlarm creates a new http.Handler for the list alerts for alarm operation
func NewListAlertsForAlarm(ctx *middleware.Context, handler ListAlertsForAlarmHandler) *ListAlertsForAlarm {
	return &ListAlertsForAlarm{Context: ctx, Handler: handler}
}

/*ListAlertsForAlarm swagger:route GET /alarms/{alarmId}/alerts alerts listAlertsForAlarm

Gets all alerts for the alarm

*/
type ListAlertsForAlarm struct {
	Context *middleware.Context
	Handler ListAlertsForAlarmHandler
}

func (o *ListAlertsForAlarm) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAlertsForAlarmParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
