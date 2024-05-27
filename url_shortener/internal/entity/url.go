package entity

import (
	"time"
)

type IURL interface {
	// GetShort returns the shortened URL.
	GetShort() string

	// GetOrigin returns the original URL.
	GetOrigin() string

	// GetCreatedAt returns the time when the URL was created.
	GetCreatedAt() time.Time
}

// URL represents a shortened URL.
//
// Fields:
// - ID: the unique identifier of the URL.
// - Origin: the original URL.
// - Short: the shortened URL.
// - Clicks: the number of times the URL has been clicked.
// - CreatedAt: the time when the URL was created.
type URL struct {
	Short     string    `json:"short"`      // the shortened URL
	Origin    string    `json:"origin"`     // the original URL
	CreatedAt time.Time `json:"created_at"` // the time when the URL was created
}

func NewURL(short, origin string) IURL {
	return &URL{
		Short:     short,
		Origin:    origin,
		CreatedAt: time.Now(),
	}
}

// GetShort returns the shortened URL.
//
// Returns:
// - string: the shortened URL.
func (u *URL) GetShort() string {
	return u.Short
}

// GetOrigin returns the original URL.
//
// Returns:
// - string: the original URL.
func (u *URL) GetOrigin() string {
	return u.Origin
}

// GetCreatedAt returns the time when the URL was created.
//
// Returns:
// - time.Time: the time when the URL was created.
func (u *URL) GetCreatedAt() time.Time {
	return u.CreatedAt
}
