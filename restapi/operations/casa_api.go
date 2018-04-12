// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"casa/src/server/restapi/operations/alarms"
	"casa/src/server/restapi/operations/alerts"
	"casa/src/server/restapi/operations/devices"
	"casa/src/server/restapi/operations/discovery"
)

// NewCasaAPI creates a new Casa instance
func NewCasaAPI(spec *loads.Document) *CasaAPI {
	return &CasaAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		AlertsAddAlertHandler: alerts.AddAlertHandlerFunc(func(params alerts.AddAlertParams) middleware.Responder {
			return middleware.NotImplemented("operation AlertsAddAlert has not yet been implemented")
		}),
		DevicesAddDeviceHandler: devices.AddDeviceHandlerFunc(func(params devices.AddDeviceParams) middleware.Responder {
			return middleware.NotImplemented("operation DevicesAddDevice has not yet been implemented")
		}),
		AlarmsCreateAlarmHandler: alarms.CreateAlarmHandlerFunc(func(params alarms.CreateAlarmParams) middleware.Responder {
			return middleware.NotImplemented("operation AlarmsCreateAlarm has not yet been implemented")
		}),
		AlarmsDeleteAlarmHandler: alarms.DeleteAlarmHandlerFunc(func(params alarms.DeleteAlarmParams) middleware.Responder {
			return middleware.NotImplemented("operation AlarmsDeleteAlarm has not yet been implemented")
		}),
		DevicesDeleteDeviceHandler: devices.DeleteDeviceHandlerFunc(func(params devices.DeleteDeviceParams) middleware.Responder {
			return middleware.NotImplemented("operation DevicesDeleteDevice has not yet been implemented")
		}),
		AlarmsGetAlarmHandler: alarms.GetAlarmHandlerFunc(func(params alarms.GetAlarmParams) middleware.Responder {
			return middleware.NotImplemented("operation AlarmsGetAlarm has not yet been implemented")
		}),
		AlarmsListAlarmLogsHandler: alarms.ListAlarmLogsHandlerFunc(func(params alarms.ListAlarmLogsParams) middleware.Responder {
			return middleware.NotImplemented("operation AlarmsListAlarmLogs has not yet been implemented")
		}),
		AlarmsListAlarmsHandler: alarms.ListAlarmsHandlerFunc(func(params alarms.ListAlarmsParams) middleware.Responder {
			return middleware.NotImplemented("operation AlarmsListAlarms has not yet been implemented")
		}),
		AlertsListAlertsHandler: alerts.ListAlertsHandlerFunc(func(params alerts.ListAlertsParams) middleware.Responder {
			return middleware.NotImplemented("operation AlertsListAlerts has not yet been implemented")
		}),
		AlertsListAlertsByIDHandler: alerts.ListAlertsByIDHandlerFunc(func(params alerts.ListAlertsByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation AlertsListAlertsByID has not yet been implemented")
		}),
		AlertsListAlertsForAlarmHandler: alerts.ListAlertsForAlarmHandlerFunc(func(params alerts.ListAlertsForAlarmParams) middleware.Responder {
			return middleware.NotImplemented("operation AlertsListAlertsForAlarm has not yet been implemented")
		}),
		DiscoveryListDiscoveredCamerasHandler: discovery.ListDiscoveredCamerasHandlerFunc(func(params discovery.ListDiscoveredCamerasParams) middleware.Responder {
			return middleware.NotImplemented("operation DiscoveryListDiscoveredCameras has not yet been implemented")
		}),
		DevicesListOfDevicesHandler: devices.ListOfDevicesHandlerFunc(func(params devices.ListOfDevicesParams) middleware.Responder {
			return middleware.NotImplemented("operation DevicesListOfDevices has not yet been implemented")
		}),
		DiscoveryNewDiscoveryHandler: discovery.NewDiscoveryHandlerFunc(func(params discovery.NewDiscoveryParams) middleware.Responder {
			return middleware.NotImplemented("operation DiscoveryNewDiscovery has not yet been implemented")
		}),
		DiscoverySetDeviceCredentialsForDiscoveryHandler: discovery.SetDeviceCredentialsForDiscoveryHandlerFunc(func(params discovery.SetDeviceCredentialsForDiscoveryParams) middleware.Responder {
			return middleware.NotImplemented("operation DiscoverySetDeviceCredentialsForDiscovery has not yet been implemented")
		}),
		AlarmsToggleAlarmHandler: alarms.ToggleAlarmHandlerFunc(func(params alarms.ToggleAlarmParams) middleware.Responder {
			return middleware.NotImplemented("operation AlarmsToggleAlarm has not yet been implemented")
		}),
	}
}

/*CasaAPI the casa API */
type CasaAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// AlertsAddAlertHandler sets the operation handler for the add alert operation
	AlertsAddAlertHandler alerts.AddAlertHandler
	// DevicesAddDeviceHandler sets the operation handler for the add device operation
	DevicesAddDeviceHandler devices.AddDeviceHandler
	// AlarmsCreateAlarmHandler sets the operation handler for the create alarm operation
	AlarmsCreateAlarmHandler alarms.CreateAlarmHandler
	// AlarmsDeleteAlarmHandler sets the operation handler for the delete alarm operation
	AlarmsDeleteAlarmHandler alarms.DeleteAlarmHandler
	// DevicesDeleteDeviceHandler sets the operation handler for the delete device operation
	DevicesDeleteDeviceHandler devices.DeleteDeviceHandler
	// AlarmsGetAlarmHandler sets the operation handler for the get alarm operation
	AlarmsGetAlarmHandler alarms.GetAlarmHandler
	// AlarmsListAlarmLogsHandler sets the operation handler for the list alarm logs operation
	AlarmsListAlarmLogsHandler alarms.ListAlarmLogsHandler
	// AlarmsListAlarmsHandler sets the operation handler for the list alarms operation
	AlarmsListAlarmsHandler alarms.ListAlarmsHandler
	// AlertsListAlertsHandler sets the operation handler for the list alerts operation
	AlertsListAlertsHandler alerts.ListAlertsHandler
	// AlertsListAlertsByIDHandler sets the operation handler for the list alerts by Id operation
	AlertsListAlertsByIDHandler alerts.ListAlertsByIDHandler
	// AlertsListAlertsForAlarmHandler sets the operation handler for the list alerts for alarm operation
	AlertsListAlertsForAlarmHandler alerts.ListAlertsForAlarmHandler
	// DiscoveryListDiscoveredCamerasHandler sets the operation handler for the list discovered cameras operation
	DiscoveryListDiscoveredCamerasHandler discovery.ListDiscoveredCamerasHandler
	// DevicesListOfDevicesHandler sets the operation handler for the list of devices operation
	DevicesListOfDevicesHandler devices.ListOfDevicesHandler
	// DiscoveryNewDiscoveryHandler sets the operation handler for the new discovery operation
	DiscoveryNewDiscoveryHandler discovery.NewDiscoveryHandler
	// DiscoverySetDeviceCredentialsForDiscoveryHandler sets the operation handler for the set device credentials for discovery operation
	DiscoverySetDeviceCredentialsForDiscoveryHandler discovery.SetDeviceCredentialsForDiscoveryHandler
	// AlarmsToggleAlarmHandler sets the operation handler for the toggle alarm operation
	AlarmsToggleAlarmHandler alarms.ToggleAlarmHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *CasaAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CasaAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CasaAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CasaAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CasaAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CasaAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CasaAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CasaAPI
func (o *CasaAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AlertsAddAlertHandler == nil {
		unregistered = append(unregistered, "alerts.AddAlertHandler")
	}

	if o.DevicesAddDeviceHandler == nil {
		unregistered = append(unregistered, "devices.AddDeviceHandler")
	}

	if o.AlarmsCreateAlarmHandler == nil {
		unregistered = append(unregistered, "alarms.CreateAlarmHandler")
	}

	if o.AlarmsDeleteAlarmHandler == nil {
		unregistered = append(unregistered, "alarms.DeleteAlarmHandler")
	}

	if o.DevicesDeleteDeviceHandler == nil {
		unregistered = append(unregistered, "devices.DeleteDeviceHandler")
	}

	if o.AlarmsGetAlarmHandler == nil {
		unregistered = append(unregistered, "alarms.GetAlarmHandler")
	}

	if o.AlarmsListAlarmLogsHandler == nil {
		unregistered = append(unregistered, "alarms.ListAlarmLogsHandler")
	}

	if o.AlarmsListAlarmsHandler == nil {
		unregistered = append(unregistered, "alarms.ListAlarmsHandler")
	}

	if o.AlertsListAlertsHandler == nil {
		unregistered = append(unregistered, "alerts.ListAlertsHandler")
	}

	if o.AlertsListAlertsByIDHandler == nil {
		unregistered = append(unregistered, "alerts.ListAlertsByIDHandler")
	}

	if o.AlertsListAlertsForAlarmHandler == nil {
		unregistered = append(unregistered, "alerts.ListAlertsForAlarmHandler")
	}

	if o.DiscoveryListDiscoveredCamerasHandler == nil {
		unregistered = append(unregistered, "discovery.ListDiscoveredCamerasHandler")
	}

	if o.DevicesListOfDevicesHandler == nil {
		unregistered = append(unregistered, "devices.ListOfDevicesHandler")
	}

	if o.DiscoveryNewDiscoveryHandler == nil {
		unregistered = append(unregistered, "discovery.NewDiscoveryHandler")
	}

	if o.DiscoverySetDeviceCredentialsForDiscoveryHandler == nil {
		unregistered = append(unregistered, "discovery.SetDeviceCredentialsForDiscoveryHandler")
	}

	if o.AlarmsToggleAlarmHandler == nil {
		unregistered = append(unregistered, "alarms.ToggleAlarmHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CasaAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CasaAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *CasaAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *CasaAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *CasaAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CasaAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the casa API
func (o *CasaAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CasaAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/alarms/{alarmId}/alerts/{id}"] = alerts.NewAddAlert(o.context, o.AlertsAddAlertHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/devices"] = devices.NewAddDevice(o.context, o.DevicesAddDeviceHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/alarms"] = alarms.NewCreateAlarm(o.context, o.AlarmsCreateAlarmHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/alarms/{id}"] = alarms.NewDeleteAlarm(o.context, o.AlarmsDeleteAlarmHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/devices/{id}"] = devices.NewDeleteDevice(o.context, o.DevicesDeleteDeviceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alarms/{id}"] = alarms.NewGetAlarm(o.context, o.AlarmsGetAlarmHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alarms/{id}/toggle"] = alarms.NewListAlarmLogs(o.context, o.AlarmsListAlarmLogsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alarms"] = alarms.NewListAlarms(o.context, o.AlarmsListAlarmsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alerts"] = alerts.NewListAlerts(o.context, o.AlertsListAlertsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alarms/{alarmId}/alerts/{id}"] = alerts.NewListAlertsByID(o.context, o.AlertsListAlertsByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alarms/{alarmId}/alerts"] = alerts.NewListAlertsForAlarm(o.context, o.AlertsListAlertsForAlarmHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/discovery"] = discovery.NewListDiscoveredCameras(o.context, o.DiscoveryListDiscoveredCamerasHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/devices"] = devices.NewListOfDevices(o.context, o.DevicesListOfDevicesHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/discovery"] = discovery.NewNewDiscovery(o.context, o.DiscoveryNewDiscoveryHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/discovery/devices/{serialNumber}/credentials"] = discovery.NewSetDeviceCredentialsForDiscovery(o.context, o.DiscoverySetDeviceCredentialsForDiscoveryHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/alarms/{id}/toggle"] = alarms.NewToggleAlarm(o.context, o.AlarmsToggleAlarmHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CasaAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *CasaAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}