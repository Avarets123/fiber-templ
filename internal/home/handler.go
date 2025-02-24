package home

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	router fiber.Router
}

func ApplyHanlder(r fiber.Router) {

	h := &handler{
		router: r,
	}

	api := h.router.Group("/api")

	api.Get("/", h.home)
	api.Get("/error", h.error)

}

func (h *handler) home(c *fiber.Ctx) error {
	data := struct {
		Count           int
		IsAdmin, CanUse bool
	}{Count: 10, IsAdmin: true, CanUse: true}

	return c.Render("page", data)
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
