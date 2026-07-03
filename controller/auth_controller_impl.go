package controller

import (
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"GO-AUTH-JWT/service"
	"fmt"
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

	fmt.Println(name, email, password)

	request := request.UserCreateRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}
	userResponse := controller.AuthService.Register(ctx, request)

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
	loginResponse := controller.AuthService.Login(ctx, request)

	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   loginResponse,
	})
}
