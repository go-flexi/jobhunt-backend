package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-flexi/jobhunt-backend/pkg/errors"
)

func TestParseStrQueryParam(t *testing.T) {
	tests := map[string]struct {
		queryParams url.Values
		key         string
		wantValue   string
		wantFound   bool
	}{
		"KeyExists": {
			queryParams: url.Values{"foo": []string{"bar"}},
			key:         "foo",
			wantValue:   "bar",
			wantFound:   true,
		},
		"KeyDoesNotExist": {
			queryParams: url.Values{"foo": []string{"bar"}},
			key:         "baz",
			wantValue:   "",
			wantFound:   false,
		},
		"EmptyValue": {
			queryParams: url.Values{"empty": []string{""}},
			key:         "empty",
			wantValue:   "",
			wantFound:   false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req := &http.Request{URL: &url.URL{RawQuery: tt.queryParams.Encode()}}
			gotValue, gotFound := ParseStrQueryParam(tt.key, req)

			if gotValue != tt.wantValue || gotFound != tt.wantFound {
				t.Errorf("ParseStrQueryParam() = (%v, %v), want (%v, %v)", gotValue, gotFound, tt.wantValue, tt.wantFound)
			}
		})
	}
}

func TestParseChiURLParams(t *testing.T) {
	tests := map[string]struct {
		urlPath   string
		key       string
		wantValue string
		wantFound bool
	}{
		"ParamExists": {
			urlPath:   "/users/123",
			key:       "userID",
			wantValue: "123",
			wantFound: true,
		},
		"ParamDoesNotExist": {
			urlPath:   "/users/123",
			key:       "orderID",
			wantValue: "",
			wantFound: false,
		},
		"EmptyParam": {
			urlPath:   "/users/",
			key:       "userID",
			wantValue: "",
			wantFound: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// Create a new chi router
			r := chi.NewRouter()

			// Define a simple route for testing
			r.Get("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
				value, found := ParseChiURLParams(tt.key, r)
				if found {
					w.Write([]byte(value))
				} else {
					http.Error(w, "not found", http.StatusNotFound)
				}
			})

			// Create a test server
			ts := httptest.NewServer(r)
			defer ts.Close()

			// Perform a GET request to the test server
			resp, err := http.Get(ts.URL + tt.urlPath)
			if err != nil {
				t.Fatalf("Failed to perform GET request: %v", err)
			}
			defer resp.Body.Close()

			// Check if we expect the parameter to be found
			if tt.wantFound {
				if resp.StatusCode != http.StatusOK {
					t.Errorf("Expected status 200, got %d", resp.StatusCode)
				}
				// Verify the response body matches the expected value
				body := make([]byte, len(tt.wantValue))
				resp.Body.Read(body)
				if string(body) != tt.wantValue {
					t.Errorf("Expected value %s, got %s", tt.wantValue, string(body))
				}
			} else {
				if resp.StatusCode != http.StatusNotFound {
					t.Errorf("Expected status 404, got %d", resp.StatusCode)
				}
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	testTable := map[string]struct {
		key          string
		parseStrFunc func(key string, r *http.Request) (string, bool)
		expectedInt  int
		expectedErr  error
	}{
		"ok": {
			key: "foo",
			parseStrFunc: func(key string, r *http.Request) (string, bool) {
				return "123", true
			},
			expectedInt: 123,
			expectedErr: nil,
		},
		"empty": {
			key: "foo",
			parseStrFunc: func(key string, r *http.Request) (string, bool) {
				return "", false
			},
			expectedInt: 0,
			expectedErr: ErrEmpty,
		},
		"not an int": {
			key: "foo",
			parseStrFunc: func(key string, r *http.Request) (string, bool) {
				return "abc", true
			},
			expectedInt: 0,
			expectedErr: errors.NewValidationsErr(errors.NewValidation("foo", fmt.Errorf("value must be an integer"))),
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			req := &http.Request{}
			i, err := ParseInt(tc.key, req, tc.parseStrFunc)
			if i != tc.expectedInt {
				t.Errorf("expected int %v, got %v", tc.expectedInt, i)
			}

			if (err == nil && tc.expectedErr != nil) || err != nil && tc.expectedErr == nil {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
