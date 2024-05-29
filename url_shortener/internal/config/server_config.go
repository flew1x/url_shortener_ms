package config

const (
	SERVER_PORT    = "server_bind_port"
	SERVER_BIND_IP = "server_bind_ip"
	SERVER_SCHEME  = "server_scheme"

	LIMIT_PER_SECOND = "rate_limit_per_second"
)

type IServerConfig interface {
	// GetPort returns the port of the HTTP server.
	GetPort() string

	// GetBindIP returns the bind IP address of the HTTP server.
	GetBindIP() string

	// GetLimitPerSecond returns the limit per second of the HTTP server.
	GetLimitPerSecond() int

	// GetScheme returns the scheme of the HTTP server.
	GetScheme() string
}

type ServerConfig struct{}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{}
}

// GetPort returns the port of the HTTP server.
//
// Returns:
// - int: the port of the HTTP server.
func (s *ServerConfig) GetPort() string {
	return mustString(SERVER_PORT)
}

// GetBindIP returns the bind IP address of the HTTP server.
//
// Returns:
// - string: the bind IP address of the HTTP server.
func (s *ServerConfig) GetBindIP() string {
	return mustString(SERVER_BIND_IP)
}

// GetLimitPerSecond returns the limit per second of the HTTP server.
//
// Returns:
// - int: the limit per second of the HTTP server.
func (s *ServerConfig) GetLimitPerSecond() int {
	return mustInt(LIMIT_PER_SECOND)
}

// GetScheme returns the scheme of the HTTP server.
//
// Returns:
// - string: the scheme of the HTTP server.
//
// The scheme is one of "http" or "https".
func (s *ServerConfig) GetScheme() string {
	return mustString(SERVER_SCHEME)
}
