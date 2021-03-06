// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ToggleAlarm toggle alarm
// swagger:model toggleAlarm

type ToggleAlarm struct {

	// alarm Id
	AlarmID string `json:"alarmId,omitempty"`

	// date
	Date strfmt.DateTime `json:"date,omitempty"`

	// on
	On bool `json:"on,omitempty"`
}

/* polymorph toggleAlarm alarmId false */

/* polymorph toggleAlarm date false */

/* polymorph toggleAlarm on false */

// Validate validates this toggle alarm
func (m *ToggleAlarm) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ToggleAlarm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ToggleAlarm) UnmarshalBinary(b []byte) error {
	var res ToggleAlarm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
