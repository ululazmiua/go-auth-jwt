package app

import (
	"GO-AUTH-JWT/controller"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	App *fiber.App
}

func BuildServer(app *fiber.App, authController controller.AuthController, userController controller.UserController, eventController controller.EventController) *Server {

	RegisterRoutes(
		app,
		authController,
		userController,
		eventController,
	)

	return &Server{
		App: app,
	}
}
