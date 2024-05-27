package entity

import (
	"time"
)

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

func NewURL(short, origin string) URL {
	return URL{
		Short:     short,
		Origin:    origin,
		CreatedAt: time.Now(),
	}
}
