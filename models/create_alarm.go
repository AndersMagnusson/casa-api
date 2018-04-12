// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateAlarm create alarm
// swagger:model createAlarm

type CreateAlarm struct {

	// continous
	Continous bool `json:"continous,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Array of serial numbers
	Devices []string `json:"devices"`

	// id
	ID string `json:"id,omitempty"`

	// motion detection
	MotionDetection bool `json:"motionDetection,omitempty"`

	// sms
	Sms *Sms `json:"sms,omitempty"`
}

/* polymorph createAlarm continous false */

/* polymorph createAlarm description false */

/* polymorph createAlarm devices false */

/* polymorph createAlarm id false */

/* polymorph createAlarm motionDetection false */

/* polymorph createAlarm sms false */

// Validate validates this create alarm
func (m *CreateAlarm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSms(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAlarm) validateDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	return nil
}

func (m *CreateAlarm) validateSms(formats strfmt.Registry) error {

	if swag.IsZero(m.Sms) { // not required
		return nil
	}

	if m.Sms != nil {

		if err := m.Sms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sms")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAlarm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAlarm) UnmarshalBinary(b []byte) error {
	var res CreateAlarm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
