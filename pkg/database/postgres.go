package database

import (
	"context"
	"fiber-templ/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func NewPostgresDb(cfg *config.DatabaseCfg, logger *zerolog.Logger) *pgxpool.Pool {

	dbPool, err := pgxpool.New(context.Background(), cfg.Url)

	if dbPool == nil {
		logger.Error().Msg("Error in connection to database!")
		logger.Error().Msg(err.Error())
		panic(err)
	}

	logger.Info().Msg("Connection to database successfully!")
	return dbPool

}
