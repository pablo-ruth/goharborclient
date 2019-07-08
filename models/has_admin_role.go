// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HasAdminRole has admin role
// swagger:model HasAdminRole
type HasAdminRole struct {

	// 1-has admin, 0-not.
	HasAdminRole bool `json:"has_admin_role,omitempty"`
}

// Validate validates this has admin role
func (m *HasAdminRole) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HasAdminRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HasAdminRole) UnmarshalBinary(b []byte) error {
	var res HasAdminRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
