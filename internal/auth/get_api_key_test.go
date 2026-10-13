package auth

import (
	"net/http"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name  string
		input http.Header
		want  string
		error bool
	}{
		{name: "Good", input: http.Header{"Authorization": []string{"ApiKey 123456"}}, want: "123456", error: false},
		{name: "No Key", input: http.Header{"Authorization": []string{"ApiKey123456"}}, want: "123456", error: true},
		{name: "No Header", input: http.Header{}, want: "123456", error: true},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if err != nil {
			if !tc.error {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		} else if tc.want != got {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

/*
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
	} */
