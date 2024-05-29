package config

const (
	LOG_LEVEL = "logging_mode"
)

type ILoggerConfig interface {
	// GetLogLevel returns the logger configuration's log level.
	GetLogLevel() string
}

type LoggerConfig struct{}

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{}
}

// GetLogLevel returns the logger configuration's log level.
//
// Returns:
// - string: the log level.
func (l *LoggerConfig) GetLogLevel() string {
	return mustString(LOG_LEVEL)
}
