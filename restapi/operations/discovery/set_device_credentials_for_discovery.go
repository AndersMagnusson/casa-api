// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SetDeviceCredentialsForDiscoveryHandlerFunc turns a function with the right signature into a set device credentials for discovery handler
type SetDeviceCredentialsForDiscoveryHandlerFunc func(SetDeviceCredentialsForDiscoveryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SetDeviceCredentialsForDiscoveryHandlerFunc) Handle(params SetDeviceCredentialsForDiscoveryParams) middleware.Responder {
	return fn(params)
}

// SetDeviceCredentialsForDiscoveryHandler interface for that can handle valid set device credentials for discovery params
type SetDeviceCredentialsForDiscoveryHandler interface {
	Handle(SetDeviceCredentialsForDiscoveryParams) middleware.Responder
}

// NewSetDeviceCredentialsForDiscovery creates a new http.Handler for the set device credentials for discovery operation
func NewSetDeviceCredentialsForDiscovery(ctx *middleware.Context, handler SetDeviceCredentialsForDiscoveryHandler) *SetDeviceCredentialsForDiscovery {
	return &SetDeviceCredentialsForDiscovery{Context: ctx, Handler: handler}
}

/*SetDeviceCredentialsForDiscovery swagger:route POST /discovery/devices/{serialNumber}/credentials discovery setDeviceCredentialsForDiscovery

Discover Axis devices

*/
type SetDeviceCredentialsForDiscovery struct {
	Context *middleware.Context
	Handler SetDeviceCredentialsForDiscoveryHandler
}

func (o *SetDeviceCredentialsForDiscovery) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSetDeviceCredentialsForDiscoveryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
