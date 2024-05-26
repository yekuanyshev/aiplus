package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yekuanyshev/aiplus/internal/transport/rest/handler"
)

func setRoutes(app *fiber.App, handler *handler.Manager) {
	api := app.Group("/api")

	cityRouter := api.Group("/city")
	{
		cityRouter.Get("/", handler.City.List)
		cityRouter.Post("/", handler.City.Create)
		cityRouter.Put("/:id", handler.City.Update)
	}

	employeeRouter := api.Group("/employee")
	{
		employeeRouter.Get("/:id", handler.Employee.GetByID)
		employeeRouter.Post("/", handler.Employee.Create)
	}
}
