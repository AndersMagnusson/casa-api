// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListAlertsHandlerFunc turns a function with the right signature into a list alerts handler
type ListAlertsHandlerFunc func(ListAlertsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAlertsHandlerFunc) Handle(params ListAlertsParams) middleware.Responder {
	return fn(params)
}

// ListAlertsHandler interface for that can handle valid list alerts params
type ListAlertsHandler interface {
	Handle(ListAlertsParams) middleware.Responder
}

// NewListAlerts creates a new http.Handler for the list alerts operation
func NewListAlerts(ctx *middleware.Context, handler ListAlertsHandler) *ListAlerts {
	return &ListAlerts{Context: ctx, Handler: handler}
}

/*ListAlerts swagger:route GET /alerts alerts listAlerts

Gets all alerts

*/
type ListAlerts struct {
	Context *middleware.Context
	Handler ListAlertsHandler
}

func (o *ListAlerts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAlertsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}