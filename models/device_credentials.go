// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeviceCredentials device credentials
// swagger:model deviceCredentials

type DeviceCredentials struct {

	// password
	Password string `json:"password,omitempty"`

	// serial number
	SerialNumber string `json:"serialNumber,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

/* polymorph deviceCredentials password false */

/* polymorph deviceCredentials serialNumber false */

/* polymorph deviceCredentials username false */

// Validate validates this device credentials
func (m *DeviceCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceCredentials) UnmarshalBinary(b []byte) error {
	var res DeviceCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}