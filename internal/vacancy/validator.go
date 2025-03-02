package vacancy

import (
	"fmt"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

func validateVacancyCreate(vacancy *VacancyCreateForm) string {

	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name: "Email", Field: vacancy.Email,
		},
		&validators.StringIsPresent{
			Name: "Role", Field: vacancy.Role,
		},

		&validators.StringIsPresent{
			Name: "Direction", Field: vacancy.Direction,
		},
		&validators.StringIsPresent{
			Name: "Location", Field: vacancy.Location,
		},

		&validators.StringIsPresent{
			Name: "Salary", Field: vacancy.Salary,
		},

		&validators.StringIsPresent{
			Name: "Company", Field: vacancy.Company,
		},
	)

	if len(errors.Errors) == 0 {
		return ""
	}

	errMsg := ""

	for k, v := range errors.Errors {

		errMsg += fmt.Sprintf("[%s]: %s;", k, v)

	}

	return errMsg

}
