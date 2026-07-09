package app

import (
	"GO-AUTH-JWT/controller"
	"GO-AUTH-JWT/middleware"

	"github.com/gofiber/fiber/v3"
)

// ! package app itu apa?
// ● Tempat setup komponen utama aplikasi
// ● Biasanya berisi inisialisasi database
// ● Setup router & middleware
// ● Menyiapkan dependency (repository, service, controller)

func RegisterRoutes(app *fiber.App, authController controller.AuthController, userController controller.UserController, eventController controller.EventController) {
	middleware.AuthMiddleware(app)
	// auth
	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)
	// user
	app.Put("/user", userController.Update)
	app.Delete("/user", userController.Delete)
	// event
	app.Get("/user/event/:eventId", eventController.FindById)
	app.Get("/user/event", eventController.FindAll)
	app.Post("/user/event", eventController.Create)
	app.Put("/user/event/:eventId", eventController.Update)
	app.Delete("/user/event/:eventId", eventController.Delete)

}
