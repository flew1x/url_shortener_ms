package config

const (
	REDIS_HOST     = "REDIS_HOST"
	REDIS_PASSWORD = "REDIS_PASSWORD"
	REDIS_PORT     = "REDIS_PORT"
	REDIS_URL_DB   = "redis_url_db"
)

type IRedisConfig interface {
	// GetRedisAddress returns the address of the Redis server.
	GetRedisHost() string

	// GetRedisPassword returns the password of the Redis server.
	GetRedisPassword() string

	// GetRedisPort returns the port of the Redis server.
	GetRedisPort() string

	// GetRedisDB returns the database number of the Redis server to use for URL shortening.
	GetRedisUrlDB() int
}

type RedisConfig struct{}

// NewRedisConfig returns an instance of IRedisConfig interface
// implementing the Redis configuration.
//
// Returns:
// - IRedisConfig: an instance of IRedisConfig interface.
func NewRedisConfig() *RedisConfig {
	return &RedisConfig{}
}

// NewRedisConfig returns an instance of IRedisConfig interface
// implementing the Redis configuration.
//
// Returns:
// - IRedisConfig: an instance of IRedisConfig interface.

func (r *RedisConfig) GetRedisHost() string {
	return mustStringFromEnv(REDIS_HOST)
}

// GetRedisPassword returns the password of the Redis server.
//
// Returns:
// - string: the password of the Redis server.
func (r *RedisConfig) GetRedisPassword() string {
	return mustStringFromEnv(REDIS_PASSWORD)
}

// GetRedisPort returns the port of the Redis server.
//
// Returns:
// - string: the port of the Redis server.
func (r *RedisConfig) GetRedisPort() string {
	return mustStringFromEnv(REDIS_PORT)
}

// GetRedisUrlDB returns the database number of the Redis server to use for URL shortening.
//
// Returns:
// - int: the database number of the Redis server.
func (r *RedisConfig) GetRedisUrlDB() int {
	return mustInt(REDIS_URL_DB)
}
