// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VerbCounter verb_counter
//
// # Holds verb counters
//
// swagger:model verb_counter
type VerbCounter struct {

	// count
	Count interface{} `json:"count,omitempty"`

	// verb
	Verb Verb `json:"verb,omitempty"`
}

// Validate validates this verb counter
func (m *VerbCounter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VerbCounter) validateVerb(formats strfmt.Registry) error {
	if swag.IsZero(m.Verb) { // not required
		return nil
	}

	if err := m.Verb.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("verb")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("verb")
		}
		return err
	}

	return nil
}

// ContextValidate validate this verb counter based on the context it is used
func (m *VerbCounter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVerb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VerbCounter) contextValidateVerb(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Verb) { // not required
		return nil
	}

	if err := m.Verb.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("verb")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("verb")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VerbCounter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VerbCounter) UnmarshalBinary(b []byte) error {
	var res VerbCounter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
