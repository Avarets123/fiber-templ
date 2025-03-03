package vacancy

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Repository struct {
	db     *pgxpool.Pool
	logger *zerolog.Logger
}

func NewRepo(db *pgxpool.Pool, logger *zerolog.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) GetAll() ([]Vacancy, error) {
	query := `SELECT * FROM vacancies ORDER BY createdat`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[Vacancy])

}

func (r *Repository) CreateVacancy(vacancy *VacancyCreateForm) error {

	query := `INSERT INTO vacancies (email, salary, location, company, role, direction, createdat)
			  VALUES (@email, @salary, @location, @company, @role, @direction, @createdat)
	`

	args := pgx.NamedArgs{
		"email":     vacancy.Email,
		"salary":    vacancy.Salary,
		"location":  vacancy.Location,
		"company":   vacancy.Company,
		"role":      vacancy.Role,
		"direction": vacancy.Direction,
		"createdat": time.Now(),
	}

	_, err := r.db.Exec(context.Background(), query, args)

	if err != nil {
		return err
	}

	return nil
}
