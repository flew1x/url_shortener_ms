package cache

import (
	"context"
	"log/slog"

	"github.com/flew1x/url_shortener_auth_ms/internal/config"
	"github.com/flew1x/url_shortener_auth_ms/internal/entity"
	"github.com/redis/go-redis/v9"
)

type IUrlCache interface {
	// Get retrieves a URL from the cache using its long URL.
	GetByShortUrl(ctx context.Context, shortUrl string) (entity.URL, error)

	// Set saves a URL in the cache using its short URL.
	SetByShortUrl(ctx context.Context, url entity.URL) error

	// Get retrieves a URL from the cache using its short URL.
	GetByLongUrl(ctx context.Context, longUrl string) (entity.URL, error)

	// Set saves a URL in the cache using its long URL.
	SetByLongUrl(ctx context.Context, url entity.URL) error
}

// redisUserTokenCache is an implementation of IUrlCache interface
// that stores data in Redis cache.
type redisUserTokenCache struct {
	// logger is used for logging.
	logger *slog.Logger
	// config is a configuration for Redis client.
	config config.IRedisConfig
	// urlConfig is a configuration for URLs.
	urlConfig config.IUrlConfig
	// client is a Redis client.
	client *redis.Client
}

func NewUrlCache(logger *slog.Logger, config config.IRedisConfig, urlConfig config.IUrlConfig, client *redis.Client) IUrlCache {
	return &redisUserTokenCache{logger: logger, config: config, urlConfig: urlConfig, client: client}
}

// createUrlKey constructs the key for a URL in the cache.
//
// Parameters:
// - shortURL: the short URL for which to create the key.
//
// Returns:
// - string: the key for the URL in the cache.
func (c *redisUserTokenCache) createUrlKey(shortURL string) string {
	return "url:" + shortURL
}

// GetByShortUrl retrieves a URL from the cache using its long URL.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - longURL: the long URL to retrieve from the cache.
//
// Returns:
// - entity.URL: the URL retrieved from the cache.
// - error: an error if the operation failed.
func (c *redisUserTokenCache) GetByShortUrl(ctx context.Context, shortURL string) (entity.URL, error) {
	key := c.createUrlKey(shortURL)

	c.logger.Debug("Getting URL from cache ", slog.String("key", key))
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		c.logger.Debug("Failed to get URL from cache: %s", err)
		return entity.URL{}, err
	}

	c.logger.Debug("Retrieved URL from cache ", slog.String("key", key), slog.String("value", value))
	return entity.URL{Short: key, Origin: value}, nil
}

// SetByShortUrl saves a URL in the cache using its short URL.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - url: the URL to save in the cache.
//
// Returns:
// - error: an error if the operation failed.
func (c *redisUserTokenCache) SetByShortUrl(ctx context.Context, url entity.URL) error {
	key := c.createUrlKey(url.Short)

	c.logger.Debug("Saving URL to cache ", slog.String("key", key), slog.String("value", url.Origin))
	err := c.client.Set(ctx, key, url.Origin, c.urlConfig.LiveCaheExpiration()).Err()
	if err != nil {
		c.logger.Debug("Failed to save URL to cache: %s", err)
		return err
	}

	c.logger.Debug("Saved URL to cache ", slog.String("key", key), slog.String("value", url.Origin))
	return nil
}

// GetByLongUrl retrieves a URL from the cache using its long URL.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - longUrl: the long URL to retrieve from the cache.
//
// Returns:
// - entity.URL: the URL retrieved from the cache.
// - error: an error if the operation failed.
func (c *redisUserTokenCache) GetByLongUrl(ctx context.Context, longUrl string) (entity.URL, error) {
	key := c.createUrlKey(longUrl)

	c.logger.Debug("Getting URL from cache ", slog.String("key", key))
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		c.logger.Debug("Failed to get URL from cache: %s", err)
		return entity.URL{}, err
	}

	c.logger.Debug("Retrieved URL from cache ", slog.String("key", key), slog.String("value", value))
	return entity.URL{Short: key, Origin: value}, nil
}

// SetByLongUrl saves a URL in the cache using its long URL.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - url: the URL to save in the cache.
//
// Returns:
// - error: an error if the operation failed.
func (c *redisUserTokenCache) SetByLongUrl(ctx context.Context, url entity.URL) error {
	key := c.createUrlKey(url.Origin)

	c.logger.Debug("Saving URL to cache ", slog.String("key", key), slog.String("value", url.Origin))
	err := c.client.Set(ctx, key, url.Origin, c.urlConfig.LiveCaheExpiration()).Err()
	if err != nil {
		c.logger.Debug("Failed to save URL to cache: %s", err)
		return err
	}

	c.logger.Debug("Saved URL to cache ", slog.String("key", key), slog.String("value", url.Origin))
	return nil
}
