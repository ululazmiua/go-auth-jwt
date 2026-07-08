package app

import (
	"GO-AUTH-JWT/controller"
	"GO-AUTH-JWT/exception"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)


type Server struct {
	App *fiber.App
}


// untuk dependecy injection google wire, maka perlu di modifikasi seperti ini
func NewServer() *fiber.App {
	appFiber := fiber.New(fiber.Config{ // ? fiber.New(fiber.Config) *fiber.App, digunakan untuk membuat object Fiber baru
		ReadTimeout:  5 * time.Second,        // ? ReadTimeout adalah waktu maksimal untuk membaca request dari client
		WriteTimeout: 5 * time.Second,        // ? WriteTimeout adalah waktu maksimal untuk menulis response ke client
		IdleTimeout:  5 * time.Second,        // ? IdleTimeout adalah waktu maksimal untuk menjaga koneksi tetap terbuka tanpa aktivitas
		ErrorHandler: exception.ErrorHandler, // ? ErrorHandler untuk menangani error
	})
	appFiber.Use(recover.New())

	return appFiber
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
