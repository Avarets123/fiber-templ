package vacancy

import (
	"fiber-templ/pkg/templ"
	"fiber-templ/views/components"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type handler struct {
	router      fiber.Router
	vacancyRepo *Repository
	logger      *zerolog.Logger
}

func ApplyHanlder(router fiber.Router, vacancyRepo *Repository, logger *zerolog.Logger) {
	h := &handler{
		router:      router,
		vacancyRepo: vacancyRepo,
		logger:      logger,
	}

	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.create)
	vacancyGroup.Get("/", h.getAll)

}

func (h *handler) getAll(c *fiber.Ctx) error {
	vacancies, err := h.vacancyRepo.GetAll(10, 5)
	if err != nil {
		h.logger.Error().Msg(err.Error())
		c.SendStatus(500)
	}

	return c.JSON(vacancies)

}

func (h *handler) create(c *fiber.Ctx) error {

	form := VacancyCreateForm{
		Email:     c.FormValue("email"),
		Salary:    c.FormValue("salary"),
		Company:   c.FormValue("company"),
		Role:      c.FormValue("role"),
		Location:  c.FormValue("location"),
		Direction: c.FormValue("direction"),
	}

	if errMsg := validateVacancyCreate(&form); errMsg != "" {
		component := components.Notification(errMsg, components.NotificationFail)
		return templ.Render(c, component)
	}

	err := h.vacancyRepo.CreateVacancy(&form)
	if err != nil {
		component := components.Notification(err.Error(), components.NotificationFail)
		return templ.Render(c, component)
	}

	component := components.Notification("Вакансия создана!", components.NotificationSuccess)

	return templ.Render(c, component)
}
