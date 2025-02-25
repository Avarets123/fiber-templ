package home

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	router fiber.Router
}
type User struct {
	Id   int
	Name string
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

	users := []User{
		{Id: 1, Name: "Vasya"},
		{Id: 2, Name: "Andrey"},
		{Id: 3, Name: "Sasha"},
	}

	names := []string{"Alex", "Tanya"}

	data := struct {
		Names []string
		Users []User
	}{Names: names, Users: users}

	return c.Render("page2", data)
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
