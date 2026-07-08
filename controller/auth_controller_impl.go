package controller

import (
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"GO-AUTH-JWT/service"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

// ! auth_controller_impl.go itu untuk apa?
// ● File implementasi dari interface controller
// ● Tempat isi nyata handler API auth
// ● Menjalankan proses request–response
// ● Memanggil auth_service
// ● Biasanya berisi:
// ● kode real untuk GET auth
// ● POST auth
// ● UPDATE auth
// ● DELETE auth

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Register(ctx fiber.Ctx) error {
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	request := request.UserCreateRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}
	context := context.Background()

	userResponse := controller.AuthService.Register(context, request)

	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   userResponse,
	})

}

func (controller *AuthControllerImpl) Login(ctx fiber.Ctx) error {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	request := request.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	context := context.Background()
	loginResponse := controller.AuthService.Login(context, request)

	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   loginResponse,
	})
}
