// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScalingPolicyConfiguration scaling policy configuration
// swagger:model ScalingPolicyConfiguration
type ScalingPolicyConfiguration struct {

	// Method by which to add capacity to the AutoscalingGroup
	//
	// Absolute represents an exact number of nodes, whereas percent represents a percentage (rounded up to the nearest whole number) of nodes in the pool.
	// Required: true
	// Enum: [absolute percent]
	AdjustmentType *string `json:"adjustment_type"`

	// Numerical representation of the number of nodes to scale the AutoscalingGroup down by determined by the adjustmentType
	// Required: true
	AdjustmentValue *int32 `json:"adjustment_value"`

	// The comparison operator to use when comparing the MetricsBackend metric value to the threshold value
	// Required: true
	ComparisonOperator *string `json:"comparison_operator"`

	// Numerical representation of the threshold at which when the comparison evaluates to true, the associated AutoscalingGroups should scale
	// Required: true
	Threshold *float32 `json:"threshold"`
}

// Validate validates this scaling policy configuration
func (m *ScalingPolicyConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdjustmentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdjustmentValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComparisonOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scalingPolicyConfigurationTypeAdjustmentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["absolute","percent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scalingPolicyConfigurationTypeAdjustmentTypePropEnum = append(scalingPolicyConfigurationTypeAdjustmentTypePropEnum, v)
	}
}

const (

	// ScalingPolicyConfigurationAdjustmentTypeAbsolute captures enum value "absolute"
	ScalingPolicyConfigurationAdjustmentTypeAbsolute string = "absolute"

	// ScalingPolicyConfigurationAdjustmentTypePercent captures enum value "percent"
	ScalingPolicyConfigurationAdjustmentTypePercent string = "percent"
)

// prop value enum
func (m *ScalingPolicyConfiguration) validateAdjustmentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, scalingPolicyConfigurationTypeAdjustmentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ScalingPolicyConfiguration) validateAdjustmentType(formats strfmt.Registry) error {

	if err := validate.Required("adjustment_type", "body", m.AdjustmentType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAdjustmentTypeEnum("adjustment_type", "body", *m.AdjustmentType); err != nil {
		return err
	}

	return nil
}

func (m *ScalingPolicyConfiguration) validateAdjustmentValue(formats strfmt.Registry) error {

	if err := validate.Required("adjustment_value", "body", m.AdjustmentValue); err != nil {
		return err
	}

	return nil
}

func (m *ScalingPolicyConfiguration) validateComparisonOperator(formats strfmt.Registry) error {

	if err := validate.Required("comparison_operator", "body", m.ComparisonOperator); err != nil {
		return err
	}

	return nil
}

func (m *ScalingPolicyConfiguration) validateThreshold(formats strfmt.Registry) error {

	if err := validate.Required("threshold", "body", m.Threshold); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScalingPolicyConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScalingPolicyConfiguration) UnmarshalBinary(b []byte) error {
	var res ScalingPolicyConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
