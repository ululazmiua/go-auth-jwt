package main

import (
	"GO-AUTH-JWT/app"
	"GO-AUTH-JWT/controller"
	"GO-AUTH-JWT/exception"
	"GO-AUTH-JWT/repository"
	"GO-AUTH-JWT/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

// ! main.go itu untuk apa?

// ● Titik awal aplikasi Go dijalankan
// ● Mengatur semua dependency
// ● Menghubungkan controller, service, repository
// ● Setup database connection
// ● Setup router/server HTTP
// ● Biasanya berisi:
// ● inisialisasi DB
// ● wiring struct (dependency injection)
// ● start server

func main() {

	/*// ! layer clean architecture
	-> controller
	-> service
	-> repository
	-> database

	● Prinsip utama:
	● Layer atas boleh tahu layer bawah,
	● layer bawah tidak tahu layer atas
	*/

	db := app.OpenConnection()  // ? app.OpenConnection() digunakan untuk inisialisasi database
	validate := validator.New() // ? validator.New() digunakan untuk inisialisasi validator

	// ! Dependency Injection
	userRepository := repository.NewUserRepository() // ? repository.NewUserRepository(db) digunakan untuk inisialisasi repository user
	// eventRepository := repository.NewEventRepository() // ? repository.NewEventRepository(db) digunakan untuk inisialisasi repository event

	// userService := service.NewUserService(userRepository, db, validate)
	// eventService := service.NewEventService(eventRepository, db, validate)
	authService := service.NewAuthService(userRepository, db, validate)

	authController := controller.NewAuthController(authService)
	// eventController := controller.NewEventController(eventService)
	// userController := controller.NewUserController(userService)

	/*
		! Configuration
		● Saat kita membuat fiber.App menggunakan fiber.New() terdapat parameter fiber.Config yang bisa kita gunakan
		● Ada banyak sekali konfigurasi yang bisa kita ubah, dan kita akan bahas secara bertahap
		● Contoh yang bisa kita gunakan adalah mengubah konfigurasi timeout
	*/
	appFiber := fiber.New(fiber.Config{ // ? fiber.New(fiber.Config) *fiber.App, digunakan untuk membuat object Fiber baru
		ReadTimeout:  5 * time.Second,        // ? ReadTimeout adalah waktu maksimal untuk membaca request dari client
		WriteTimeout: 5 * time.Second,        // ? WriteTimeout adalah waktu maksimal untuk menulis response ke client
		IdleTimeout:  5 * time.Second,        // ? IdleTimeout adalah waktu maksimal untuk menjaga koneksi tetap terbuka tanpa aktivitas
		ErrorHandler: exception.ErrorHandler, // ? ErrorHandler untuk menangani error
	})
	appFiber.Use(recover.New())

	// ! Routing
	router := app.NewRouter(appFiber, authController)

	// ! Start Server
	err := router.Listen("localhost:3000", fiber.ListenConfig{
		EnablePrefork: true, // ? EnablePrefork adalah opsi untuk mengaktifkan preforking, jika diaktifkan, maka Fiber akan menjalankan server dalam mode preforking, sehingga lebih cepat dalam menjalankan server, karena Fiber akan membuat beberapa proses worker untuk menangani request secara bersamaan, sehingga meningkatkan performa server(menggunakan semua core CPU yang tersedia)
	}) // ? app.Listen(addr string, config ...fiber.listenConfig) error, digunakan untuk menjalankan aplikasi Fiber pada alamat dan port tertentu
	if err != nil {
		panic(err)
	}
}
