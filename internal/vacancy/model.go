package vacancy

import "time"

type VacancyCreateForm struct {
	Email,
	Salary,
	Company,
	Role,
	Location,
	Direction string
}

type Vacancy struct {
	Id        string    `db:"id"`
	Email     string    `db:"email"`
	Salary    string    `db:"salary"`
	Company   string    `db:"company"`
	Role      string    `db:"role"`
	Location  string    `db:"location"`
	Direction string    `db:"direction"`
	CreatedAt time.Time `db:"createdat"`
}
