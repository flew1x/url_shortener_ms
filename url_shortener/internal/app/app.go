package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/flew1x/url_shortener_auth_ms/internal/cache"
	"github.com/flew1x/url_shortener_auth_ms/internal/config"
	http_v1 "github.com/flew1x/url_shortener_auth_ms/internal/controllers/http/v1"
	"github.com/flew1x/url_shortener_auth_ms/internal/repository"
	"github.com/flew1x/url_shortener_auth_ms/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	config *config.Config
	router *gin.Engine
	logger *slog.Logger
}

// createAddress constructs the address string for a server.
//
// Parameters:
// - address: the IP address of the server.
// - port: the port number of the server.
//
// Returns:
// - string: the address string in the format "address:port".
func createAddress(address string, port string) string {
	return address + ":" + port
}

func createMongoURL(config *config.Config) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.MongoConfig.GetMongoUsername(),
		url.QueryEscape(config.MongoConfig.GetMongoPassword()),
		config.MongoConfig.GetMongoHost(),
		config.MongoConfig.GetMongoPort(),
	)
}

// InitialServer initializes a new Server instance.
//
// Parameters:
// - ctx: the context.Context for the function.
// - config: the configuration object.
// - logger: the logger object.
//
// Returns:
// - *Server: a new instance of Server.
// - error: an error if there was an issue initializing the server.
func InitialServer(ctx context.Context, config *config.Config, logger *slog.Logger) (*Server, error) {
	// Connect to Redis
	redisOptions := &redis.Options{
		Addr:     createAddress(config.RedisConfig.GetRedisHost(), config.RedisConfig.GetRedisPort()),
		Password: config.RedisConfig.GetRedisPassword(),
		DB:       config.RedisConfig.GetRedisUrlDB(),
	}
	redisClient := redis.NewClient(redisOptions)

	// Initialize cache
	cache := cache.NewCache(logger, config.RedisConfig, config.URLConfig, redisClient)

	mongoDatabase, err := mongoDatabase(ctx, logger, config)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// Initialize repositories
	repositories := repository.NewRepository(logger, config, mongoDatabase)

	// Initialize services
	services := service.NewService(logger, repositories, cache, config)

	// Initialize handlers
	handlers := http_v1.NewHandler(logger, services, config, cache)

	// Initialize router
	router := handlers.InitRoutes()

	logger.Info("Starting the application...")

	// Initialize and return App
	return &Server{config: config, router: router, logger: logger}, nil
}

// mongoDatabase initializes a new MongoDB database connection.
//
// Parameters:
// - ctx: the context.Context for the function.
// - logger: the logger object.
// - config: the configuration object.
//
// Returns:
// - *mongo.Database: the initialized MongoDB database connection.
// - error: an error if there was an issue initializing the database connection.
func mongoDatabase(ctx context.Context, logger *slog.Logger, config *config.Config) (*mongo.Database, error) {
	// Connect to MongoDB
	mongoURL := createMongoURL(config)

	logger.Info("Connecting to MongoDB... " + mongoURL)

	clientOptions := options.Client().ApplyURI(mongoURL)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	mongoDatabase := mongoClient.Database(config.MongoConfig.GetMongoDatabase())

	return mongoDatabase, nil
}

// Run runs the Server.
//
// It starts the HTTP server with the provided context.
func (a *Server) Run() {
	a.StartHTTP(context.Background())
}

// StartHTTP starts the HTTP server.
//
// ctx context.Context
func (a *Server) StartHTTP(ctx context.Context) {
	address := createAddress(a.config.ServerConfig.GetBindIP(), a.config.ServerConfig.GetPort())
	err := a.router.Run(address)
	if err != nil {
		panic(err)
	}

	a.logger.Info("Server started", slog.String("address", address))
}
