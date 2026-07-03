package controller

import "github.com/gofiber/fiber/v3"

// ! user_controller.go itu untuk apa?
// ● File khusus endpoint API untuk user
// ● Menerima request dari client
// ● Mengambil data input (JSON, params)
// ● Memanggil user_service untuk diproses
// ● Mengirim response ke client
// ● Biasanya berisi:
// ● handler GET user
// ● handler POST user
// ● handler PUT user
// ● handler DELETE user

type AuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
