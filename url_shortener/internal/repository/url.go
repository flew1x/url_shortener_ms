package repository

import (
	"context"
	"log/slog"

	"github.com/flew1x/url_shortener_auth_ms/internal/config"
	"github.com/flew1x/url_shortener_auth_ms/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IURLRepository interface {
	// Create creates a new URL in the repository.
	Create(ctx context.Context, url entity.URL) error

	// GetByOrigin returns a URL from the repository by its origin.
	GetByOrigin(ctx context.Context, origin string) (entity.URL, error)

	// GetByShort returns a URL from the repository by its short.
	GetByShort(ctx context.Context, short string) (entity.URL, error)

	// DeleteByID deletes a URL from the repository by its ID.
	Delete(ctx context.Context, short string) error

	// Update updates a URL in the repository by its ID.
	Update(ctx context.Context, url entity.URL) error
}

type urlRepository struct {
	logger     *slog.Logger
	config     *config.IURLConfig
	collection *mongo.Collection
}

func NewURLRepository(logger *slog.Logger, config *config.IURLConfig, database *mongo.Database) IURLRepository {
	return &urlRepository{logger: logger, config: config, collection: database.Collection(URLS_COLLECTION)}
}

// Create creates a new URL in the repository.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - url: the URL to create in the repository.
//
// Returns:
// - error: an error if the operation failed.
func (l *urlRepository) Create(ctx context.Context, url entity.URL) error {
	l.logger.Debug("Creating URL in repository", "origin", url.Origin, "short", url.Short)
	_, err := l.collection.InsertOne(ctx, url)
	if err != nil {
		l.logger.Error("Error creating URL in repository: " + err.Error())
		return err
	}
	l.logger.Debug("URL created successfully")
	return nil
}

// Delete deletes a URL from the repository by its short.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - short: the shortened URL to delete.
//
// Returns:
// - error: an error if the operation failed.
func (l *urlRepository) Delete(ctx context.Context, short string) error {
	_, err := l.collection.DeleteOne(ctx, entity.URL{Short: short})
	if err != nil {
		l.logger.Error("error deleting url: " + err.Error())
		return err
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
func (l *urlRepository) GetByOrigin(ctx context.Context, origin string) (entity.URL, error) {
	var url entity.URL
	err := l.collection.FindOne(ctx, entity.URL{Origin: origin}).Decode(&url)
	if err != nil {
		l.logger.Error("error getting url: " + err.Error())
		return entity.URL{}, err
	}

	return url, nil
}

// GetByShort retrieves a URL from the repository by its short.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - short: the shortened URL to retrieve from the repository.
//
// Returns:
// - entity.URL: the URL retrieved from the repository.
// - error: an error if the operation failed.
func (l *urlRepository) GetByShort(ctx context.Context, short string) (entity.URL, error) {
	var url entity.URL
	filter := bson.M{"short": short}
	l.logger.Debug("Getting URL by short: " + short)
	err := l.collection.FindOne(ctx, filter).Decode(&url)
	if err != nil {
		l.logger.Error("error getting url: " + err.Error())
		return entity.URL{}, err
	}
	l.logger.Debug("Retrieved URL: " + url.Origin)
	return url, nil
}

// Update updates a URL in the repository by its ID.
//
// Parameters:
// - ctx: the context.Context for the operation.
// - url: the URL to update in the repository.
//
// Returns:
// - error: an error if the operation failed.
func (l *urlRepository) Update(ctx context.Context, url entity.URL) error {
	_, err := l.collection.UpdateOne(ctx, entity.URL{Short: url.Short}, url)
	if err != nil {
		l.logger.Error("error updating url: " + err.Error())
		return err
	}

	return nil
}
