package utils

import "strings"

func ValidateOrigin(originURL string) error {
	// Check if the URL starts with "http://" or "https://"
	if !strings.HasPrefix(originURL, "http://") && !strings.HasPrefix(originURL, "https://") {
		return ErrNotValidURL
	}

	// Check if the URL contains a host
	if !strings.Contains(originURL, "//") {
		return ErrNotValidURL
	}

	// Check if the URL contains a top-level domain
	if !strings.HasSuffix(originURL, ".com") && !strings.HasSuffix(originURL, ".org") && !strings.HasSuffix(originURL, ".net") {
		// Add any other top-level domains you want to check
		return ErrNotValidURL
	}

	// If all checks pass, return nil
	return nil
}
