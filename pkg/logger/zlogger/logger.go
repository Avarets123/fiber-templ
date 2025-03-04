package zlogger

import (
	"fiber-templ/config"
	"os"

	"github.com/rs/zerolog"
)

func New(cfg *config.LoggerConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(cfg.LogLevel))
	var logger zerolog.Logger

	if cfg.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		// consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, FormatCaller: func(i interface{}) string {
		// 	pathSlices := strings.Split(i.(string), "/")
		// 	finalPath := strings.Join(pathSlices[len(pathSlices)-3:], "/")
		// 	return finalPath
		// }}
		logger = zerolog.New(consoleWriter).With().Timestamp().Caller().Logger()
	}

	logger.Info().Msg("ZeroLogger inited!")

	return &logger

}
