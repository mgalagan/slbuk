// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GrpcEntity grpc entity
//
// swagger:model grpcEntity
type GrpcEntity struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this grpc entity
func (m *GrpcEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this grpc entity based on context it is used
func (m *GrpcEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GrpcEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrpcEntity) UnmarshalBinary(b []byte) error {
	var res GrpcEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
