package service

import (
	"log/slog"

	"github.com/flew1x/url_shortener_auth_ms/internal/cache"
	"github.com/flew1x/url_shortener_auth_ms/internal/config"
	"github.com/flew1x/url_shortener_auth_ms/internal/repository"
)

type Service struct {
	UrlShortener IURLService
}

func NewService(logger *slog.Logger, repository *repository.Repository, cache *cache.Cache, config *config.Config) *Service {
	return &Service{
		UrlShortener: NewURLService(logger, repository.UrlRepository, cache.UrlCache, config),
	}
}
