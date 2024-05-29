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

// GetCreatedAt implements IURL.
func (u *URL) GetCreatedAt() time.Time {
	return u.CreatedAt
}

// GetOrigin implements IURL.
func (u *URL) GetOrigin() string {
	return u.Origin
}

// GetShort implements IURL.
func (u *URL) GetShort() string {
	return u.Short
}

func NewURL(short, origin string) IURL {
	return &URL{
		Short:     short,
		Origin:    origin,
		CreatedAt: time.Now(),
	}
}
