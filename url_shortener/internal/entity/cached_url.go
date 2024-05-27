package entity

// CachedURL represents a shortened URL in the cache.
//
// Fields:
// - Origin: the original URL.
// - Short: the shortened URL.
type CachedURL struct {
	ID     string `json:"id"`
	Origin string `json:"origin"` // the original URL
	Short  string `json:"short"`  // the shortened URL
}

// NewCachedURL creates a new CachedURL.
//
// Parameters:
// - origin: the original URL.
// - short: the shortened URL.
//
// Returns:
// - CachedURL: the new CachedURL.
func NewCachedURL(origin, short, id string) CachedURL {
	return CachedURL{
		ID:     id,
		Origin: origin,
		Short:  short,
	}
}
