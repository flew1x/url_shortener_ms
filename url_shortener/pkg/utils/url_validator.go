package utils

import (
	"strings"
)

func ValidateOrigin(originURL string) error {
	// Check if the URL starts with "http://" or "https://"
	if !strings.HasPrefix(originURL, "http://") && !strings.HasPrefix(originURL, "https://") {
		return ErrNotValidURL
	}

	// Check if the URL contains a host
	if !strings.Contains(originURL, "//") {
		return ErrNotValidURL
	}

	// If all checks pass, return nil
	return nil
}
