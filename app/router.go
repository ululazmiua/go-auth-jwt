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

func NewRouter(app *fiber.App, authController controller.AuthController, userController controller.UserController) *fiber.App {

	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)
	app.Put("/user/:id", userController.Update)
	app.Delete("/user/:id", userController.Delete)

	return app
}
