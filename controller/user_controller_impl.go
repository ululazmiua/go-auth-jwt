package controller

import (
	"GO-AUTH-JWT/helper"
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"GO-AUTH-JWT/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// ! user_controller_impl.go itu untuk apa?
// ● File implementasi dari interface controller
// ● Tempat isi nyata handler API user
// ● Menjalankan proses request–response
// ● Memanggil user_service
// ● Biasanya berisi:
// ● kode real untuk GET user
// ● POST user
// ● UPDATE user
// ● DELETE user

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) Update(ctx fiber.Ctx) (_ error) {
	userUpdateRequest := request.UserUpdateRequest{}
	err := ctx.Bind().Body(&userUpdateRequest)

	UserId, err := strconv.Atoi(ctx.Params("id"))
	helper.PanicIfError(err)

	userUpdateRequest.ID = UserId
	

	userResponse := controller.userService.Update(ctx.Context(), userUpdateRequest)
	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   userResponse,
	})
}

func (controller *UserControllerImpl) Delete(ctx fiber.Ctx) (_ error) {
	userUpdateRequest := ctx.Params("id")
	userId, err := strconv.Atoi(userUpdateRequest)
	helper.PanicIfError(err)

	controller.userService.Delete(ctx.Context(), int64(userId))
	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	})
}
