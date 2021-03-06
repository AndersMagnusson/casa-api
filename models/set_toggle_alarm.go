// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SetToggleAlarm set toggle alarm
// swagger:model setToggleAlarm

type SetToggleAlarm struct {

	// on
	On bool `json:"on,omitempty"`
}

/* polymorph setToggleAlarm on false */

// Validate validates this set toggle alarm
func (m *SetToggleAlarm) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SetToggleAlarm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetToggleAlarm) UnmarshalBinary(b []byte) error {
	var res SetToggleAlarm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
