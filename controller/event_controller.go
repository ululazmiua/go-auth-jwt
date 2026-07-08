package controller

import "github.com/gofiber/fiber/v3"

// ! event_controller.go itu untuk apa?
// ● File khusus endpoint API untuk event
// ● Menerima request dari client
// ● Mengambil data input (JSON, params)
// ● Memanggil event_service untuk diproses
// ● Mengirim response ke client
// ● Biasanya berisi:
// ● handler GET event
// ● handler POST event
// ● handler PUT event
// ● handler DELETE event

type EventController interface {
	Create(ctx fiber.Ctx) error
	Update(ctx fiber.Ctx) error
	Delete(ctx fiber.Ctx) error
	FindById(ctx fiber.Ctx) error
	FindAll(ctx fiber.Ctx) error
}
