package service

import (
	"context"
	"log/slog"
	"net/url"

	"github.com/flew1x/url_shortener_ms/internal/cache"
	"github.com/flew1x/url_shortener_ms/internal/config"
	"github.com/flew1x/url_shortener_ms/internal/entity"
	"github.com/flew1x/url_shortener_ms/internal/repository"
	"github.com/flew1x/url_shortener_ms/pkg/utils"
)

type IURLService interface {
	// Create creates a new URL in the repository.
	Create(ctx context.Context, originUrl string) (shortUrl string, err error)

	// GetByOrigin returns a URL from the repository by its origin.
	GetByOrigin(ctx context.Context, origin string) (entity.IURL, error)

	// GetByShort returns a URL from the repository by its short.
	GetByShort(ctx context.Context, short string) (entity.IURL, error)

	// DeleteByID deletes a URL from the repository by its ID.
	Delete(ctx context.Context, short string) error

	// Update updates a URL in the repository by its ID.
	Update(ctx context.Context, url entity.IURL) error

	// BuildShortURL builds the short URL from the given short ID.
	BuildShortURL(short string) url.URL
}

type URLService struct {
	logger        *slog.Logger
	urlRepository repository.IURLRepository
	cache         cache.IUrlCache
	config        *config.Config
}

func NewURLService(logger *slog.Logger, urlRepository repository.IURLRepository, cache cache.IUrlCache, config *config.Config) *URLService {
	return &URLService{logger: logger, urlRepository: urlRepository, cache: cache, config: config}
}

// generateShortUrl generates a random short URL of the given length.
//
// Parameters:
// - length: the length of the short URL.
//
// Returns:
// - url.URL: the generated short URL.
func (l *URLService) generateShortUrl(length int) url.URL {
	rand := utils.SeededRand()

	b := make([]byte, length)
	for i := range b {
		b[i] = SYMBOLS[rand.Intn(len(SYMBOLS))]
	}

	path := "s/" + string(b)

	url := url.URL{
		Scheme: l.config.ServerConfig.GetScheme(),
		Host:   l.config.ServerConfig.GetBindIP(),
		Path:   path,
	}

	return url
}

// BuildShortURL builds the short URL from the given short ID.
//
// Parameters:
// - short: the short ID of the URL.
//
// Returns:
// - url.URL: the built short URL.
func (l *URLService) BuildShortURL(short string) url.URL {
	path := "s/" + short

	return url.URL{
		Scheme: l.config.ServerConfig.GetScheme(),
		Host:   l.config.ServerConfig.GetBindIP(),
		Path:   path,
	}
}

// Create creates a new URL entry in the repository and returns its short URL.
//
// Parameters:
// - ctx: the context.Context for the function.
// - originURL: the original URL to be shortened.
//
// Returns:
// - shortURL: the shortened URL.
// - err: an error if the URL is not valid or if there was an issue creating the short URL.
func (s *URLService) Create(ctx context.Context, originURL string) (shortURL string, err error) {
	// Validate the origin URL
	if err = utils.ValidateOrigin(originURL); err != nil {
		s.logger.Error("Error validating origin URL " + err.Error())
		return "", err
	}

	// Log the origin URL
	s.logger.Debug("Creating URL ", slog.Any("origin", originURL))

	// Check if the URL is present in the cache
	if cachedURL, err := s.cache.GetByLongUrl(ctx, originURL); err == nil {
		s.logger.Debug("URL found in cache ", slog.Any("url", cachedURL.GetShort()))
		return cachedURL.GetShort(), nil
	}

	// Generate a new short URL
	shortGeneratedURL := s.generateShortUrl(s.config.URLConfig.LengthShortURL())

	// Log the generated short URL
	s.logger.Debug("Generated short URL ", slog.Any("short", shortGeneratedURL.String()))

	urlObject := entity.NewURL(shortGeneratedURL.String(), originURL)

	// Save the URL in the repository
	if err := s.urlRepository.Create(ctx, urlObject); err != nil {
		s.logger.Error("Error creating URL " + err.Error())
		return "", err
	}

	// Set the URL in the cache by long URL
	if err := s.cache.SetByLongUrl(ctx, urlObject); err != nil {
		s.logger.Error("Error setting URL in cache by long URL " + err.Error())
		return "", err
	}

	// Set the URL in the cache by short URL
	if err := s.cache.SetByShortUrl(ctx, urlObject); err != nil {
		s.logger.Error("Error setting URL in cache by short URL " + err.Error())
		return "", err
	}

	// Return the short URL
	return urlObject.GetShort(), nil
}

// GetByOrigin retrieves a URL from the repository by its origin.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - origin: the original URL to retrieve from the repository.
//
// Returns:
// - entity.URL: the URL retrieved from the repository.
// - error: an error if the operation failed.
func (l *URLService) GetByOrigin(ctx context.Context, origin string) (entity.IURL, error) {
	if err := utils.ValidateOrigin(origin); err != nil {
		return nil, err
	}

	url, err := l.urlRepository.GetByOrigin(ctx, origin)
	if err != nil {
		l.logger.Error("error getting url " + err.Error())
		return nil, err
	}

	return url, nil
}

// GetByShortID retrieves a URL from the repository or cache by its short.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - shortID: the shortened URL to retrieve.
//
// Returns:
// - entity.URL: the URL retrieved from the repository or cache.
// - error: an error if the operation failed.
func (l *URLService) GetByShort(ctx context.Context, shortID string) (entity.IURL, error) {
	// Log the beginning of the function
	l.logger.Debug("GetByShort function started ")

	// Validate the shortened URL
	err := utils.ValidateOrigin(shortID)
	if err != nil {
		l.logger.Error("error validating origin " + err.Error())
		return nil, err
	}

	l.logger.Debug("Getting URL by short ID" + shortID)

	// Get the URL from the cache
	cacheURL, err := l.cache.GetByShortUrl(ctx, shortID)
	if err == nil {
		l.logger.Debug("URL found in cache" + cacheURL.GetOrigin())
		return cacheURL, nil
	}

	l.logger.Debug("URL not found in cache. Retrieving from repository ")

	// Get the URL from the repository
	repoURL, err := l.urlRepository.GetByShort(ctx, shortID)
	if err != nil {
		l.logger.Error("error getting URL from repository " + err.Error())
		return nil, err
	}

	l.logger.Debug("URL found in repository " + repoURL.GetOrigin())

	// Set the URL in the cache
	err = l.cache.SetByShortUrl(ctx, repoURL)
	if err != nil {
		l.logger.Error("error setting URL in cache " + err.Error())
		return nil, err
	}

	err = l.cache.SetByLongUrl(ctx, repoURL)
	if err != nil {
		l.logger.Error("error setting URL in cache " + err.Error())
		return nil, err
	}

	// Log the end of the function
	l.logger.Debug("GetByShort function ended ")

	return repoURL, nil
}

// Delete deletes a URL from the repository by its short.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - short: the shortened URL to delete.
//
// Returns:
// - error: an error if the operation failed.
func (l *URLService) Delete(ctx context.Context, short string) error {
	if err := l.urlRepository.Delete(ctx, short); err != nil {
		l.logger.Error("error deleting url " + err.Error())
		return err
	}

	return nil
}

// Update updates a URL in the repository by its ID.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - url: the URL to update in the repository.
//
// Returns:
// - error: an error if the operation failed.
func (l *URLService) Update(ctx context.Context, url entity.IURL) error {
	if err := utils.ValidateOrigin(url.GetOrigin()); err != nil {
		return err
	}

	if err := l.urlRepository.Update(ctx, url); err != nil {
		l.logger.Error("error updating url " + err.Error())
		return err
	}

	return nil
}
