package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts API Key from header
// Valid header: Authorization: ApiKey {:ApiKey}
func GetAPIKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("Authentication Not Found")
	}

	values := strings.Split(value, " ")
	if len(values) != 2 || values[0] != "ApiKey" {
		return "", errors.New("Invalid Authentication")
	}

	return values[1], nil
}
