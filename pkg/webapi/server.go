package webapi

import (
	"github.com/gofiber/fiber"
	"snake/pkg/webapi/resource/person"
	"snake/pkg/webapi/setup"
)

func ServeAndListen(port string, dependency *setup.Dependency) {
	app := fiber.New()

	person.DefineRoutes(app, dependency.Person)

	_ = app.Listen(port)
}
