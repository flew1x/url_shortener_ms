package config

import "time"

const (
	LENGTH_SHORT_URL      = "length_short_url"
	LIVE_CACHE_EXPIRATION = "live_cache_expiration"
)

type IURLConfig interface {
	// LengthShortURL returns the length of the short URL.
	LengthShortURL() int

	// LiveCaheExpiration returns the expiration time of the live cache.
	LiveCaheExpiration() time.Duration
}

type URLConfig struct{}

// NewURLConfig returns a new instance of IURLConfig.
//
// Returns:
// - IURLConfig: a new instance of IURLConfig.
func NewURLConfig() *URLConfig {
	return &URLConfig{}
}

// LengthShortURL returns the length of the short URL.
//
// Returns:
// - int: the length of the short URL.
func (u *URLConfig) LengthShortURL() int {
	return mustInt(LENGTH_SHORT_URL)
}

// LiveCacheExpiration returns the expiration time of the live cache.
//
// It reads the expiration time from the environment variable
// LIVE_CACHE_EXPIRATION and parses it as a duration.
//
// Returns:
// - time.Duration: the expiration time of the live cache.
func (u *URLConfig) LiveCaheExpiration() time.Duration {
	return mustDuration(LIVE_CACHE_EXPIRATION)
}
