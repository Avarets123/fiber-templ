package vacancy

import (
	"fiber-templ/pkg/templ"
	"fiber-templ/views/components"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	router fiber.Router
}

func ApplyHanlder(router fiber.Router) {
	h := &handler{
		router: router,
	}

	vacancyGroup := h.router.Group("/vacancy")

	vacancyGroup.Post("/", h.create)

}

func (h *handler) create(c *fiber.Ctx) error {

	email := c.FormValue("email")

	form := VacancyCreateForm{
		Email: email,
	}

	if errMsg := validateVacancyCreate(&form); errMsg != "" {
		component := components.Notification(errMsg, components.NotificationFail)
		return templ.Render(c, component)
	}

	component := components.Notification("Вакансия создана!", components.NotificationSuccess)

	return templ.Render(c, component)
}
