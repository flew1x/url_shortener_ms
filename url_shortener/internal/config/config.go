package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var cfg *koanf.Koanf

// Config represents the configuration for the URL shortener service.
type Config struct {
	// - UrlConfig: the configuration for the URL shortener service.
	UrlConfig IUrlConfig `koanf:"url_shortener"`

	// - RedisConfig: the configuration for the Redis cache.
	RedisConfig IRedisConfig `koanf:"redis"`

	// - ServerConfig: the configuration for the HTTP server.
	ServerConfig IServerConfig `koanf:"server"`

	// - LoggerConfig: the configuration for the logger.
	LoggerConfig ILoggerConfig `koanf:"logger"`

	// - MongoConfig: the configuration for the MongoDB database.
	MongoConfig IMongoConfig `koanf:"mongo"`
}

// NewConfig returns a new instance of Config with the UrlConfig field initialized
// to a new instance of IUrlConfig.
//
// Returns:
// - *Config: a new instance of Config.
func NewConfig() *Config {
	return &Config{
		UrlConfig:    NewUrlConfig(),
		RedisConfig:  NewRedisConfig(),
		ServerConfig: NewServerConfig(),
		LoggerConfig: NewLoggerConfig(),
		MongoConfig:  NewMongoConfig(),
	}
}

// InitConfig initializes the global configuration by loading the
// config from the given YAML file and parsing it.
// It panics if there is an error loading or parsing the config.
func (c *Config) InitConfig(configPath, configFile string) {
	cfg = koanf.New(".")
	filePath := filepath.Join(configPath, configFile)
	config := file.Provider(filePath)
	if err := cfg.Load(config, yaml.Parser()); err != nil {
		panic(fmt.Errorf(ConfigLoadError, err))
	}
}

// MustString returns a string value for the given field from the global config.
// It panics if the field is not found or if the value is not a string.
//
// Parameters:
// - field: the field to retrieve the value for.
//
// Returns:
// - string: the value of the field.
func mustString(field string) string {
	return cfg.MustString(field)
}

// MustInt returns an int value for the given field from the global config.
// It panics if the field is not found or if the value is not an int.
//
// Parameters:
// - field: the field to retrieve the value for.
//
// Returns:
// - int: the value of the field.
func mustInt(field string) int {
	return cfg.MustInt(field)
}

// MustFromEnv returns a string value for the given environment variable.
// It panics if the environment variable is not set.
//
// Parameters:
// - field: the name of the environment variable to retrieve the value for.
//
// Returns:
// - string: the value of the environment variable.
func mustStringFromEnv(field string) string {
	value := os.Getenv(field)
	if value == "" {
		panic(fmt.Sprintf("%s env var is required", field))
	}

	return value
}

// MustDuration returns a time.Duration value for the given field from the global config.
// It panics if the field is not found or if the value is not a duration.
//
// Parameters:
// - field: the field to retrieve the value for.
//
// Returns:
// - time.Duration: the value of the field.
func mustDuration(field string) time.Duration {
	return cfg.MustDuration(field)
}
