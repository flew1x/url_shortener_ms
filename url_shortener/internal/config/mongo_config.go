package config

const (
	MONGO_HOST     = "MONGO_HOST"
	MONGO_PASSWORD = "MONGO_PASSWORD"
	MONGO_PORT     = "MONGO_PORT"
	MONGO_USERNAME = "MONGO_USERNAME"
	MONGO_DATABASE = "MONGO_DATABASE"
)

type IMongoConfig interface {
	GetMongoHost() string
	GetMongoPassword() string
	GetMongoPort() string
	GetMongoUsername() string
	GetMongoDatabase() string
}

type MongoConfig struct{}

func NewMongoConfig() *MongoConfig {
	return &MongoConfig{}
}

// GetMongoHost returns the host of the MongoDB server.
//
// Returns:
// - string: the host of the MongoDB server.
func (m *MongoConfig) GetMongoHost() string {
	return mustStringFromEnv(MONGO_HOST)
}

// GetMongoPassword returns the password of the MongoDB server.
//
// Returns:
// - string: the password of the MongoDB server.
func (m *MongoConfig) GetMongoPassword() string {
	return mustStringFromEnv(MONGO_PASSWORD)
}

// GetMongoPort returns the port of the MongoDB server.
//
// Returns:
// - string: the port of the MongoDB server.
func (m *MongoConfig) GetMongoPort() string {
	return mustStringFromEnv(MONGO_PORT)
}

// GetMongoUsername returns the username of the MongoDB server.
//
// Returns:
// - string: the username of the MongoDB server.
func (m *MongoConfig) GetMongoUsername() string {
	return mustStringFromEnv(MONGO_USERNAME)
}

// GetMongoDatabase returns the name of the MongoDB database.
//
// Returns:
// - string: the name of the MongoDB database.
func (m *MongoConfig) GetMongoDatabase() string {
	return mustStringFromEnv(MONGO_DATABASE)
}
