package home

import (
	"encoding/json"
	"fiber-templ/internal/vacancy"
	"fiber-templ/pkg/templ"
	"fiber-templ/views"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	router     fiber.Router
	repository *vacancy.Repository
}
type User struct {
	Id   int
	Name string
}

func ApplyHanlder(r fiber.Router, repository *vacancy.Repository) {

	h := &handler{
		router:     r,
		repository: repository,
	}

	api := h.router.Group("/api")

	api.Get("/", h.home)
	api.Get("/error", h.error)

}

func (h *handler) home(c *fiber.Ctx) error {

	// users := []User{
	// 	{Id: 1, Name: "Vasya"},
	// 	{Id: 2, Name: "Andrey"},
	// 	{Id: 3, Name: "Sasha"},
	// }

	// names := []string{"Alex", "Tanya"}

	// data := struct {
	// 	Names []string
	// 	Users []User
	// }{Names: names, Users: users}

	// return c.Render("page2", data)

	vacancies, err := h.repository.GetAll()
	if err != nil {
		return c.SendStatus(500)
	}

	component := views.Main(vacancies)

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
