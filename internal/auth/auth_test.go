package auth

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	mockHeader := http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"ApiKey heresAnAPIKey"},
	}
	authHeader := mockHeader.Get("Authorization")

	mockMalformedHeader := http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer test-token"},
	}

	mockNoAuthHeader := http.Header{
		"Content-Type": {"application/json"},
	}

	mockEmptyHeader := http.Header{}

	tests := map[string]struct {
		input         http.Header
		want          string
		expectedError string
	}{
		"auth header": {
			input:         mockHeader,
			want:          strings.Split(authHeader, " ")[1],
			expectedError: "",
		},
		"no auth header": {
			input:         mockNoAuthHeader,
			want:          "",
			expectedError: "no authorization header included",
		},
		"malformed header": {
			input:         mockMalformedHeader,
			want:          "",
			expectedError: "malformed authorization header",
		},
		"empty header": {
			input:         mockEmptyHeader,
			want:          "",
			expectedError: "no authorization header included",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var errString string
			output, err := GetAPIKey(tc.input)
			if err != nil {
				errString = err.Error()
			}
			if !reflect.DeepEqual(tc.want, output) || errString != tc.expectedError {
				t.Fatalf("name: %v\n\texpected: %#v,%#v\n\tgot: %#v,%#v\n", name, tc.want, tc.expectedError, output, errString)
			}
		})
	}
}
