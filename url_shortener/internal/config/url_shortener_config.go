package config

import "time"

const (
	LENGTH_SHORT_URL      = "length_short_url"
	LIVE_CACHE_EXPIRATION = "live_cache_expiration"
)

type IUrlConfig interface {
	// LengthShortURL returns the length of the short URL.
	LengthShortURL() int

	// LiveCaheExpiration returns the expiration time of the live cache.
	LiveCaheExpiration() time.Duration
}

type urlConfig struct{}

// NewUrlConfig returns a new instance of IUrlConfig.
//
// Returns:
// - IUrlConfig: a new instance of IUrlConfig.
func NewUrlConfig() IUrlConfig {
	return &urlConfig{}
}

// LengthShortURL returns the length of the short URL.
//
// Returns:
// - int: the length of the short URL.
func (u *urlConfig) LengthShortURL() int {
	return mustInt(LENGTH_SHORT_URL)
}

// LiveCacheExpiration returns the expiration time of the live cache.
//
// It reads the expiration time from the environment variable
// LIVE_CACHE_EXPIRATION and parses it as a duration.
//
// Returns:
// - time.Duration: the expiration time of the live cache.
func (u *urlConfig) LiveCaheExpiration() time.Duration {
	return mustDuration(LIVE_CACHE_EXPIRATION)
}
