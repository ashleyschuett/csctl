// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AWSTrafficRule AWS traffic rule
// swagger:model AWSTrafficRule
type AWSTrafficRule struct {

	// List of CIDR blocks
	CidrBlocks []string `json:"cidr_blocks"`

	// From port
	// Required: true
	FromPort *int64 `json:"from_port"`

	// Protocol
	Protocol string `json:"protocol,omitempty"`

	// To port
	// Required: true
	ToPort *int64 `json:"to_port"`
}

// Validate validates this a w s traffic rule
func (m *AWSTrafficRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AWSTrafficRule) validateFromPort(formats strfmt.Registry) error {

	if err := validate.Required("from_port", "body", m.FromPort); err != nil {
		return err
	}

	return nil
}

func (m *AWSTrafficRule) validateToPort(formats strfmt.Registry) error {

	if err := validate.Required("to_port", "body", m.ToPort); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AWSTrafficRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSTrafficRule) UnmarshalBinary(b []byte) error {
	var res AWSTrafficRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
