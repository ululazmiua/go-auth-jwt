package app

import (
	"GO-AUTH-JWT/controller"

	"github.com/gofiber/fiber/v3"
)

// ! package app itu apa?
// ● Tempat setup komponen utama aplikasi
// ● Biasanya berisi inisialisasi database
// ● Setup router & middleware
// ● Menyiapkan dependency (repository, service, controller)

func NewRouter(app *fiber.App, authController controller.AuthController, userController controller.UserController, eventController controller.EventController) *fiber.App {
	// auth
	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)
	// user
	app.Put("/user/:id", userController.Update)
	app.Delete("/user/:id", userController.Delete)
	// event
	app.Get("/user/:userId/event/:eventId", eventController.FindById)
	app.Get("/user/:userId/event", eventController.FindAll)
	app.Post("/user/:userId/event", eventController.Create)
	app.Put("/user/:userId/event/:eventId", eventController.Update)
	app.Delete("/user/:userId/event/:eventId", eventController.Delete)
	return app
}
