package service

import (
	"context"
	"log/slog"
	"net/url"

	"github.com/flew1x/url_shortener_auth_ms/internal/cache"
	"github.com/flew1x/url_shortener_auth_ms/internal/config"
	"github.com/flew1x/url_shortener_auth_ms/internal/entity"
	"github.com/flew1x/url_shortener_auth_ms/internal/repository"
	"github.com/flew1x/url_shortener_auth_ms/pkg/utils"
)

type IURLService interface {
	// Create creates a new URL in the repository.
	Create(ctx context.Context, originUrl string) (shortUrl string, err error)

	// GetByOrigin returns a URL from the repository by its origin.
	GetByOrigin(ctx context.Context, origin string) (entity.URL, error)

	// GetByShort returns a URL from the repository by its short.
	GetByShort(ctx context.Context, short string) (entity.URL, error)

	// GetByID returns a URL from the repository by its ID.
	GetByID(ctx context.Context, shortID string) (entity.URL, error)

	// DeleteByID deletes a URL from the repository by its ID.
	Delete(ctx context.Context, short string) error

	// Update updates a URL in the repository by its ID.
	Update(ctx context.Context, url entity.URL) error

	// BuildShortUrl builds the short URL from the given short ID.
	BuildShortUrl(short string) url.URL
}

type urlService struct {
	logger        *slog.Logger
	urlRepository repository.IURLRepository
	cache         cache.IUrlCache
	config        *config.Config
}

func NewURLService(logger *slog.Logger, urlRepository repository.IURLRepository, cache cache.IUrlCache, config *config.Config) IURLService {
	return &urlService{logger: logger, urlRepository: urlRepository, cache: cache, config: config}
}

func (l *urlService) generateShortUrl(length int) url.URL {
	rand := utils.SeededRand()

	b := make([]byte, length)
	for i := range b {
		b[i] = SYMBOLS[rand.Intn(len(SYMBOLS))]
	}

	url := url.URL{
		Scheme: l.config.ServerConfig.GetScheme(),
		Host:   l.config.ServerConfig.GetBindIP(),
		Path:   string(b),
	}

	return url
}

// BuildShortUrl builds the short URL from the given short ID.
//
// Parameters:
// - short: the short ID of the URL.
//
// Returns:
// - url.URL: the built short URL.
func (l *urlService) BuildShortUrl(short string) url.URL {
	return url.URL{
		Scheme: l.config.ServerConfig.GetScheme(),
		Host:   l.config.ServerConfig.GetBindIP(),
		Path:   short,
	}
}

// Create creates a new URL entry in the repository and returns its short URL.
//
// It validates the given origin URL and returns an error if the URL is not valid.
// If the URL is present in the cache, it retrieves it from the cache.
// Otherwise, it generates a new short URL and creates a new URL entry in the repository.
//
// Parameters:
// - ctx: the context.Context for the function.
// - originURL: the original URL to be shortened.
//
// Returns:
// - shortURL: the shortened URL.
// - err: an error if the URL is not valid or if there was an issue creating the short URL.
func (s *urlService) Create(ctx context.Context, originURL string) (shortURL string, err error) {
	// Validate the origin URL
	if err = s.validateOrigin(originURL); err != nil {
		s.logger.Error("Error validating origin URL: %v", err)
		return "", err
	}

	// Check if the URL is present in the cache
	if cachedURL, err := s.cache.Get(ctx, originURL); err == nil {
		// Parse the cached URL and return it
		s.logger.Debug("URL found in cache: %v", slog.Any("url", cachedURL))
		parsedURL, err := url.Parse(cachedURL.Origin)
		if err != nil {
			s.logger.Error("Error parsing cached URL: %v", err)
			return "", err
		}
		return parsedURL.String(), nil
	}

	// Generate a new short URL
	shortGeneratedURL := s.generateShortUrl(s.config.UrlConfig.LengthShortURL())

	urlObject := entity.NewURL(shortGeneratedURL.Path, shortGeneratedURL.String(), originURL)

	if err := s.urlRepository.Create(ctx, urlObject); err != nil {
		s.logger.Error("Error creating URL entry: %v", err)
		return "", err
	}

	cachedUrl := entity.NewCachedURL(urlObject.Short, urlObject.Origin)
	err = s.cache.Set(ctx, cachedUrl)
	if err != nil {
		s.logger.Error("Error setting URL in cache: %v", err)
		return "", err
	}

	s.logger.Debug("URL created: %v", slog.Any("url", urlObject))
	return urlObject.Short, nil
}

// ValidateOrigin validates the given origin URL.
//
// It parses the URL and returns an error if the URL is not valid.
//
// Parameters:
// - originURL: the original URL to be validated.
//
// Returns:
// - error: an error if the URL is not valid, nil otherwise.
func (l *urlService) validateOrigin(originURL string) error {
	l.logger.Debug("Validating URL: %v", slog.Any("url", originURL))

	// Parse URL
	parsedURL, err := url.Parse(originURL)
	if err != nil {
		l.logger.Error("Error parsing URL: " + err.Error())
		return ErrNotValidURL
	}

	// Check scheme
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return ErrNotValidURL
	}

	// Check host
	if parsedURL.Host == "" {
		return ErrNotValidURL
	}

	return nil
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
func (l *urlService) GetByOrigin(ctx context.Context, origin string) (entity.URL, error) {
	if err := l.validateOrigin(origin); err != nil {
		return entity.URL{}, err
	}

	url, err := l.urlRepository.GetByOrigin(ctx, origin)
	if err != nil {
		l.logger.Error("error getting url: " + err.Error())
		return entity.URL{}, err
	}

	return url, nil
}

func (l *urlService) GetByShort(ctx context.Context, short string) (entity.URL, error) {
	if err := l.validateOrigin(short); err != nil {
		return entity.URL{}, err
	}

	url, err := l.urlRepository.GetByShort(ctx, short)
	if err != nil {
		l.logger.Error("error getting url: " + err.Error())
		return entity.URL{}, err
	}

	return url, nil
}

// GetByShortID retrieves a URL from the repository or cache by its short.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - short: the shortened URL to retrieve.
//
// Returns:
// - url.URL: the parsed URL retrieved from the repository or cache.
// - error: an error if the operation failed.
func (l *urlService) GetByID(ctx context.Context, shortID string) (entity.URL, error) {
	l.logger.Debug("Getting URL from cache")
	// Try to get the URL from the cache
	cacheUrl, err := l.cache.Get(ctx, shortID)
	if err == nil {
		l.logger.Debug("URL found in cache")
		// If the URL is found in the cache, return the parsed URL
		parsedUrl, err := url.Parse(cacheUrl.Origin)
		if err != nil {
			l.logger.Error("Error parsing cached URL: " + err.Error())
			return entity.URL{
				Origin: cacheUrl.Origin,
				Short:  cacheUrl.Short,
			}, err
		}

		return entity.URL{Origin: cacheUrl.Origin, Short: cacheUrl.Short}, nil
	}

	l.logger.Debug("URL not found in cache, getting it from repository")
	// If the URL is not found in the cache, retrieve it from the repository
	urlObject, err := l.urlRepository.GetByShort(ctx, shortID)
	if err != nil {
		l.logger.Error("Error getting URL from repository: " + err.Error())
		return entity.URL{}, err
	}

	l.logger.Debug("URL retrieved from repository, caching it")
	// Cache the URL
	cachedUrl := entity.NewCachedURL(urlObject.Origin, urlObject.Short, urlObject.ID)
	err = l.cache.Set(ctx, cachedUrl)
	if err != nil {
		l.logger.Error("Error caching URL: " + err.Error())
		return entity.URL{}, err
	}

	return entity.URL{ID: urlObject.ID, Origin: urlObject.Origin, Short: urlObject.Short, CreatedAt: urlObject.CreatedAt}, nil
}

// Delete deletes a URL from the repository by its short.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - short: the shortened URL to delete.
//
// Returns:
// - error: an error if the operation failed.
func (l *urlService) Delete(ctx context.Context, short string) error {
	if err := l.urlRepository.Delete(ctx, short); err != nil {
		l.logger.Error("error deleting url: " + err.Error())
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
func (l *urlService) Update(ctx context.Context, url entity.URL) error {
	if err := l.validateOrigin(url.Origin); err != nil {
		return err
	}

	if err := l.urlRepository.Update(ctx, url); err != nil {
		l.logger.Error("error updating url: " + err.Error())
		return err
	}

	return nil
}
