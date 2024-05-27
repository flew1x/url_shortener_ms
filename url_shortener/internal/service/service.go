package service

import (
	"log/slog"

	"github.com/flew1x/url_shortener_ms/internal/cache"
	"github.com/flew1x/url_shortener_ms/internal/config"
	"github.com/flew1x/url_shortener_ms/internal/repository"
)

type IService interface {
	GetUrlShortener() IURLService
}

type service struct {
	UrlShortener IURLService
}

func NewService(logger *slog.Logger, repository repository.IRepository, cache cache.ICache, config config.IConfig) IService {
	return &service{
		UrlShortener: NewURLService(logger, repository.GetUrlRepository(), cache.GetUrlCache(), config),
	}
}

// GetUrlShortener returns an instance of IURLService
// which is used to manage URLs.
//
// Returns:
// - IURLService: an instance of IURLService
func (s *service) GetUrlShortener() IURLService {
	return s.UrlShortener
}
