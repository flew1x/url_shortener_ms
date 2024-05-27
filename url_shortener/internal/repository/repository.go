package repository

import (
	"log/slog"

	"github.com/flew1x/url_shortener_ms/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	GetUrlRepository() IURLRepository
}

type repository struct {
	UrlRepository IURLRepository
}

func NewRepository(logger *slog.Logger, config config.IConfig, database *mongo.Database) IRepository {
	return &repository{
		UrlRepository: NewURLRepository(logger, config.GetUrlConfig(), database),
	}
}

// GetUrlRepository returns the URL repository.
//
// Returns:
// - IURLRepository: the URL repository.
func (r *repository) GetUrlRepository() IURLRepository {
	return r.UrlRepository
}
