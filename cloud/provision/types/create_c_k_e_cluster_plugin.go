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

// CreateCKEClusterPlugin create c k e cluster plugin
// swagger:model CreateCKEClusterPlugin
type CreateCKEClusterPlugin struct {

	// Plugin configuration
	Configuration interface{} `json:"configuration,omitempty"`

	// Specific implementation of this plugin type
	// Required: true
	Implementation *string `json:"implementation"`

	// Plugin type
	// Required: true
	// Enum: [cluster_management metrics logs cni csi cloud_controller_manager autoscaler]
	Type *string `json:"type"`

	// Plugin version
	Version string `json:"version,omitempty"`
}

// Validate validates this create c k e cluster plugin
func (m *CreateCKEClusterPlugin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImplementation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCKEClusterPlugin) validateImplementation(formats strfmt.Registry) error {

	if err := validate.Required("implementation", "body", m.Implementation); err != nil {
		return err
	}

	return nil
}

var createCKEClusterPluginTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cluster_management","metrics","logs","cni","csi","cloud_controller_manager","autoscaler"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createCKEClusterPluginTypeTypePropEnum = append(createCKEClusterPluginTypeTypePropEnum, v)
	}
}

const (

	// CreateCKEClusterPluginTypeClusterManagement captures enum value "cluster_management"
	CreateCKEClusterPluginTypeClusterManagement string = "cluster_management"

	// CreateCKEClusterPluginTypeMetrics captures enum value "metrics"
	CreateCKEClusterPluginTypeMetrics string = "metrics"

	// CreateCKEClusterPluginTypeLogs captures enum value "logs"
	CreateCKEClusterPluginTypeLogs string = "logs"

	// CreateCKEClusterPluginTypeCni captures enum value "cni"
	CreateCKEClusterPluginTypeCni string = "cni"

	// CreateCKEClusterPluginTypeCsi captures enum value "csi"
	CreateCKEClusterPluginTypeCsi string = "csi"

	// CreateCKEClusterPluginTypeCloudControllerManager captures enum value "cloud_controller_manager"
	CreateCKEClusterPluginTypeCloudControllerManager string = "cloud_controller_manager"

	// CreateCKEClusterPluginTypeAutoscaler captures enum value "autoscaler"
	CreateCKEClusterPluginTypeAutoscaler string = "autoscaler"
)

// prop value enum
func (m *CreateCKEClusterPlugin) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createCKEClusterPluginTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateCKEClusterPlugin) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateCKEClusterPlugin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCKEClusterPlugin) UnmarshalBinary(b []byte) error {
	var res CreateCKEClusterPlugin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
