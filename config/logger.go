package config

type LoggerConfig struct {
	LogLevel int
	Format   string
}

func NewLoggerCfg() *LoggerConfig {
	return &LoggerConfig{
		LogLevel: getEnvInt("LOG_LEVEL", 0),
		Format:   getEnvStr("LOG_FORMAT", "color"),
	}
}
