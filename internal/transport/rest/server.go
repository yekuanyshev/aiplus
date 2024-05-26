package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yekuanyshev/aiplus/internal/service"
	"github.com/yekuanyshev/aiplus/internal/transport/rest/handler"
)

type Server struct {
	app *fiber.App
}

func New(service service.Manager) *Server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	handlerManager := handler.NewManager(service)
	setRoutes(app, handlerManager)

	return &Server{
		app: app,
	}
}

func (s *Server) Start(listen string) error {
	return s.app.Listen(listen)
}

func (s *Server) Stop() error {
	return s.app.Shutdown()
}
