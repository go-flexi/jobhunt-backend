package errors

import (
	"fmt"
	"testing"
)

func TestToValidation(t *testing.T) {
	testTable := map[string]struct {
		err          error
		expectedErr  error
		expectedBool bool
	}{
		"validation error": {
			err:          NewValidation("field", fmt.Errorf("error")),
			expectedErr:  NewValidation("field", fmt.Errorf("error")),
			expectedBool: true,
		},
		"validation error with wraping": {
			err:          fmt.Errorf("wrap, %w", NewValidation("field", fmt.Errorf("error"))),
			expectedErr:  NewValidation("field", fmt.Errorf("error")),
			expectedBool: true,
		},
		"not a validation error": {
			err:          fmt.Errorf("not a validation error"),
			expectedErr:  fmt.Errorf("not a validation error"),
			expectedBool: false,
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			err, ok := ToValidation(tc.err)
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			if ok != tc.expectedBool {
				t.Errorf("expected bool %v, got %v", tc.expectedBool, ok)
			}
		})
	}
}

func TestToValidationList(t *testing.T) {
	testTable := map[string]struct {
		err          error
		expectedErr  error
		expectedBool bool
	}{
		"validation list error": {
			err: NewValidationsErr(
				NewValidation("field", fmt.Errorf("error")),
				NewValidation("field2", fmt.Errorf("error2")),
			),
			expectedErr: NewValidationsErr(
				NewValidation("field", fmt.Errorf("error")),
				NewValidation("field2", fmt.Errorf("error2")),
			),
			expectedBool: true,
		},
		"validation list error with wraping": {
			err: fmt.Errorf("wrap, %w", NewValidationsErr(
				NewValidation("field", fmt.Errorf("error")),
				NewValidation("field2", fmt.Errorf("error2")),
			)),
			expectedErr: NewValidationsErr(
				NewValidation("field", fmt.Errorf("error")),
				NewValidation("field2", fmt.Errorf("error2")),
			),
			expectedBool: true,
		},
		"not a validation list error": {
			err:          fmt.Errorf("not a validation list error"),
			expectedErr:  fmt.Errorf("not a validation list error"),
			expectedBool: false,
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			err, ok := ToValidationList(tc.err)
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			if ok != tc.expectedBool {
				t.Errorf("expected bool %v, got %v", tc.expectedBool, ok)
			}
		})
	}
}
