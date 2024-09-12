package errors

import "errors"

// Authorization is an error type that represents an authorization error.
type Authorization struct {
	reason string
}

// NewAuthorization creates a new Authorization error.
func NewAuthorizationErr(reason string) error {
	return Authorization{reason: reason}
}

// Error returns the error message.
func (a Authorization) Error() string {
	return a.reason
}

// ToAuthorization checks if the error is an Authorization error and returns it.
func ToAuthorization(err error) (error, bool) {
	var a Authorization
	if errors.As(err, &a) {
		return a, true
	}

	return err, false
}
