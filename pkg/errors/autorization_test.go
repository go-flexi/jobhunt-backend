package errors

import (
	"fmt"
	"testing"
)

func TestToAuthorization(t *testing.T) {
	testTable := map[string]struct {
		err          error
		expectedErr  error
		expectedBool bool
	}{
		"authorization error": {
			err:          NewAuthorizationErr("permission denied"),
			expectedErr:  NewAuthorizationErr("permission denied"),
			expectedBool: true,
		},
		"authorization error with wraping": {
			err:          fmt.Errorf("wrap, %w", NewAuthorizationErr("permission denied")),
			expectedErr:  NewAuthorizationErr("permission denied"),
			expectedBool: true,
		},
		"not an authorization error": {
			err:          fmt.Errorf("not an authorization error"),
			expectedErr:  fmt.Errorf("not an authorization error"),
			expectedBool: false,
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			err, ok := ToAuthorization(tc.err)
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			if ok != tc.expectedBool {
				t.Errorf("expected bool %v, got %v", tc.expectedBool, ok)
			}
		})
	}
}
