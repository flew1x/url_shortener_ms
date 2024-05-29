package cache

import (
	"context"
	"log/slog"

	"github.com/flew1x/url_shortener_ms/internal/config"
	"github.com/flew1x/url_shortener_ms/internal/entity"
	"github.com/redis/go-redis/v9"
)

type IUrlCache interface {
	// Get retrieves a URL from the cache using its long URL.
	GetByShortUrl(ctx context.Context, shortUrl string) (entity.IURL, error)

	// Set saves a URL in the cache using its short URL.
	SetByShortUrl(ctx context.Context, url entity.IURL) error

	// Get retrieves a URL from the cache using its short URL.
	GetByLongUrl(ctx context.Context, longUrl string) (entity.IURL, error)

	// Set saves a URL in the cache using its long URL.
	SetByLongUrl(ctx context.Context, url entity.IURL) error
}

// redisUserTokenCache is an implementation of IUrlCache interface
// that stores data in Redis cache.
type redisUserTokenCache struct {
	// logger is used for logging.
	logger *slog.Logger

	// config is a configuration for Redis client.
	config config.IRedisConfig

	// urlConfig is a configuration for URLs.
	urlConfig config.IURLConfig

	// client is a Redis client.
	client *redis.Client
}

func NewUrlCache(logger *slog.Logger, config config.IRedisConfig, urlConfig config.IURLConfig, client *redis.Client) IUrlCache {
	return &redisUserTokenCache{logger: logger, config: config, urlConfig: urlConfig, client: client}
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
func (c *redisUserTokenCache) GetByShortUrl(ctx context.Context, shortURL string) (entity.IURL, error) {
	value, err := c.client.Get(ctx, shortURL).Result()
	if err != nil {
		c.logger.Debug("Failed to get URL from cache", slog.String("err", err.Error()))
		return nil, err
	}

	c.logger.Debug("Retrieved URL from cache", slog.String("key", shortURL), slog.String("value", value))

	url := entity.NewURL(
		shortURL,
		value,
	)

	return url, nil
}

// SetByShortUrl saves a URL in the cache using its short URL.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - url: the URL to save in the cache.
//
// Returns:
// - error: an error if the operation failed.
func (c *redisUserTokenCache) SetByShortUrl(ctx context.Context, shortUrl entity.IURL) error {
	if err := c.client.Set(ctx, shortUrl.GetShort(), shortUrl.GetOrigin(), c.urlConfig.LiveCaheExpiration()).Err(); err != nil {
		c.logger.Debug("Failed to save URL to cache", slog.String("err", err.Error()))
		return err
	}

	c.logger.Debug("Saved URL to cache", slog.String("key", shortUrl.GetShort()), slog.String("value", shortUrl.GetOrigin()))
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
func (c *redisUserTokenCache) GetByLongUrl(ctx context.Context, longUrl string) (entity.IURL, error) {
	value, err := c.client.Get(ctx, longUrl).Result()
	if err != nil {
		c.logger.Debug("Failed to get URL from cache", slog.String("err", err.Error()))
		return nil, err
	}

	c.logger.Debug("Retrieved URL from cache", slog.String("key", longUrl), slog.String("value", value))

	url := entity.NewURL(
		value,
		longUrl,
	)

	return url, nil
}

// SetByLongUrl saves a URL in the cache using its long URL.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - url: the URL to save in the cache.
//
// Returns:
// - error: an error if the operation failed.
func (c *redisUserTokenCache) SetByLongUrl(ctx context.Context, longUrl entity.IURL) error {
	if err := c.client.Set(ctx, longUrl.GetOrigin(), longUrl.GetShort(), c.urlConfig.LiveCaheExpiration()).Err(); err != nil {
		c.logger.Debug("Failed to save URL to cache", slog.String("err", err.Error()))
		return err
	}

	c.logger.Debug("Saved URL to cache", slog.String("key", longUrl.GetOrigin()), slog.String("value", longUrl.GetShort()))
	return nil
}
