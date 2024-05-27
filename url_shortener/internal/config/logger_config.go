package config

const (
	LOG_LEVEL = "logging_mode"
)

type ILoggerConfig interface {
	// GetLogLevel returns the logger configuration's log level.
	GetLogLevel() string
}

type loggerConfig struct{}

func NewLoggerConfig() ILoggerConfig {
	return &loggerConfig{}
}

// GetLogLevel returns the logger configuration's log level.
//
// Returns:
// - string: the log level.
func (l *loggerConfig) GetLogLevel() string {
	return mustString(LOG_LEVEL)
}
