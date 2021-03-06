// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerRemoveOptions options of remove container
// swagger:model ContainerRemoveOptions
type ContainerRemoveOptions struct {

	// force
	Force bool `json:"Force,omitempty"`

	// link
	Link bool `json:"Link,omitempty"`

	// volumes
	Volumes bool `json:"Volumes,omitempty"`
}

// Validate validates this container remove options
func (m *ContainerRemoveOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerRemoveOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerRemoveOptions) UnmarshalBinary(b []byte) error {
	var res ContainerRemoveOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
