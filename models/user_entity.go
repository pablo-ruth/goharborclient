// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserEntity user entity
// swagger:model UserEntity
type UserEntity struct {

	// The ID of the user.
	UserID int64 `json:"user_id,omitempty"`

	// The name of the user.
	Username string `json:"username,omitempty"`
}

// Validate validates this user entity
func (m *UserEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserEntity) UnmarshalBinary(b []byte) error {
	var res UserEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
