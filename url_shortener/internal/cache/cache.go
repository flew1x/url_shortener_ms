package cache

import (
	"log/slog"

	"github.com/flew1x/url_shortener_auth_ms/internal/config"
	"github.com/redis/go-redis/v9"
)

// Cache represents the cache layer for the URL shortener service.
//
// It contains an IUrlCache instance that is responsible for caching
// URLs and their corresponding short URLs.
type Cache struct {
	// UrlCache is an IUrlCache implementation that is used to cache
	// URLs and their corresponding short URLs.
	UrlCache IUrlCache
}

// NewCache creates a new instance of the Cache struct.
//
// Parameters:
// - logger: a pointer to the logger.
// - config: a pointer to the Redis configuration.
// - redisClient: a pointer to the Redis client.
//
// Returns:
// - *Cache: a pointer to the Cache struct.
func NewCache(logger *slog.Logger, config config.IRedisConfig, urlConfig config.IUrlConfig, redisClient *redis.Client) *Cache {
	return &Cache{
		UrlCache: NewUrlCache(logger, config, urlConfig, redisClient),
	}
}
