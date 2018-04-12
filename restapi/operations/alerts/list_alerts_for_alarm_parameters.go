// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListAlertsForAlarmParams creates a new ListAlertsForAlarmParams object
// with the default values initialized.
func NewListAlertsForAlarmParams() ListAlertsForAlarmParams {
	var ()
	return ListAlertsForAlarmParams{}
}

// ListAlertsForAlarmParams contains all the bound params for the list alerts for alarm operation
// typically these are obtained from a http.Request
//
// swagger:parameters listAlertsForAlarm
type ListAlertsForAlarmParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Id of the alarm
	  Required: true
	  In: path
	*/
	AlarmID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *ListAlertsForAlarmParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rAlarmID, rhkAlarmID, _ := route.Params.GetOK("alarmId")
	if err := o.bindAlarmID(rAlarmID, rhkAlarmID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAlertsForAlarmParams) bindAlarmID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.AlarmID = raw

	return nil
}