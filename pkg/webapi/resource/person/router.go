package person

import (
	"github.com/gofiber/fiber"
	"snake/pkg/repositories/person"
)

func DefineRoutes(app *fiber.App, repository *person.Repository) {
	handler := NewHandler(repository)

	personHandlers := app.Group("/person")
	personHandlers.Get("/list", handler.GetAll)
	personHandlers.Get("/find/:id", handler.Find)
	personHandlers.Post("/store", handler.Store)
	personHandlers.Put("/update/:id", handler.Update)
	personHandlers.Delete("/remove/:id", handler.Remove)
}
