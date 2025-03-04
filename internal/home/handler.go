package home

import (
	"encoding/json"
	"fiber-templ/internal/vacancy"
	"fiber-templ/pkg/templ"
	"fiber-templ/views"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type handler struct {
	router     fiber.Router
	repository *vacancy.Repository
	logger     *zerolog.Logger
}
type User struct {
	Id   int
	Name string
}

func ApplyHanlder(r fiber.Router, repository *vacancy.Repository, logger *zerolog.Logger) {

	h := &handler{
		router:     r,
		repository: repository,
		logger:     logger,
	}

	api := h.router.Group("/")

	api.Get("/", h.home)
	api.Get("/error", h.error)

}

func (h *handler) home(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := 2

	offset := (page - 1) * limit

	vacancies, err := h.repository.GetAll(limit, offset)
	if err != nil {
		h.logger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}

	itemsCount := h.repository.GetCount()
	totalPage := math.Ceil(float64(itemsCount) / float64(limit))

	h.logger.Info().Msgf("%f", totalPage)
	h.logger.Info().Msgf("%d", itemsCount)

	component := views.Main(vacancies, int(totalPage), page)

	return templ.Render(c, component)
}

func (h *handler) error(c *fiber.Ctx) error {
	// panic("Error")
	// return c.SendString("ERROR")

	errBody, _ := json.Marshal(map[string]interface{}{
		"success": false,
		"message": "something went wrong!",
	})

	return fiber.NewError(400, string(errBody))

}
