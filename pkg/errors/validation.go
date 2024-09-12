package errors

import (
	"bytes"
	"errors"
	"fmt"
)

// Validation is an error type that represents a validation error.
type Validation struct {
	Field string
	Err   error
}

// Error returns the error message.
func (v Validation) Error() string {
	return fmt.Sprintf("field: %s, error: %s", v.Field, v.Err.Error())
}

// NewValidation creates a new Validation.
func NewValidation(field string, err error) Validation {
	return Validation{Field: field, Err: err}
}

// ToValidation checks if the error is a Validation error and returns it.
// If the error is not a Validation, it returns the original error.
func ToValidation(err error) (error, bool) {
	var v Validation
	if ok := errors.As(err, &v); ok {
		return v, true
	}

	return err, false
}

// ValidationList is a list of Validation errors.
type ValidationList []Validation

// NewValidationsErr creates a new ValidationList error.
func NewValidationsErr(vs ...Validation) error {
	return ValidationList(vs)
}

// Error returns the error message.
func (vl ValidationList) Error() string {
	buf := bytes.Buffer{}
	for i, v := range vl {
		if i > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(v.Error())
	}

	return buf.String()
}

// ToValidationList checks if the error is a ValidationList error and returns it.
// If the error is not a ValidationList, it returns the original error.
func ToValidationList(err error) (error, bool) {
	var vl ValidationList
	if ok := errors.As(err, &vl); ok {
		return vl, true
	}

	return err, false
}
