package repository

import (
	"log/slog"

	"github.com/flew1x/url_shortener_auth_ms/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	UrlRepository IURLRepository
}

func NewRepository(logger *slog.Logger, config *config.Config, database *mongo.Database) *Repository {
	return &Repository{
		UrlRepository: NewURLRepository(logger, &config.UrlConfig, database),
	}
}
