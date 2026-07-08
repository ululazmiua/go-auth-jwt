package exception

import (
	"GO-AUTH-JWT/models/dto/response"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// ! package exception itu apa?

// ● Tempat khusus untuk mengelola error aplikasi
// ● package exception = pusat penanganan kesalahan (error handling layer)
// ● Supaya aplikasi tidak panic sembarangan dan response API tetap rapi

func ErrorHandler(c fiber.Ctx, err error) error {
	// debug
	// fmt.Printf("Tipe error asli: %T | Nilai: %v\n", err, err)

	if invalidEmailPassword(c, err) {
		return nil
	}
	if emailAlreadyExistsError(c, err) {
		return nil
	}
	if notFoundError(c, err) {
		return nil
	}
	if validationError(c, err) {
		return nil
	}
	InternalServerError(c, err)
	return nil
}

func invalidEmailPassword(ctx fiber.Ctx, err any) bool {
	exception, ok := err.(InvalidEmailPassword)
	if ok {
		ctx.Set("Content-Type", "application/json")
		ctx.Status(http.StatusUnauthorized) // 401

		webResponse := response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Data:   exception.error, // atau exception.Message tergantung structmu
		}
		ctx.JSON(webResponse)
		return true
	}
	return false
}
func emailAlreadyExistsError(ctx fiber.Ctx, err any) bool {
	exception, ok := err.(EmailAlreadyExistsError)
	if ok {
		ctx.Set("Content-Type", "application/json")
		ctx.Status(http.StatusBadRequest) // 400

		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   exception.error,
		}
		ctx.JSON(webResponse)
		return true
	}
	return false
}
func notFoundError(ctx fiber.Ctx, err any) bool {
	exception, ok := err.(NotFoundError) // ? err.(NotFoundError) / var.(struct) => digunakan untuk mengecek apakah error tersebut adalah NotFoundError
	if ok {
		ctx.Set("Content-Type", "application/json")
		ctx.Status(http.StatusNotFound) // ? menambahkan status code nya adalah 404

		webResponse := response.WebResponse{
			Code:   http.StatusNotFound, // 404
			Status: http.StatusText(http.StatusNotFound),
			Data:   exception.error, // ? exception.error => digunakan untuk mengambil pesan error
		}

		ctx.JSON(webResponse)
		return true
	} else {
		return false
	}
}

func validationError(ctx fiber.Ctx, err any) bool {
	exception, ok := err.(validator.ValidationErrors) // ? err.(validator.ValidationErrors) / var.(struct) => digunakan untuk mengecek apakah error tersebut adalah validator.ValidationErrors
	if ok {
		ctx.Set("Content-Type", "application/json")
		ctx.Status(http.StatusBadRequest)

		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest, // 400
			Status: http.StatusText(http.StatusBadRequest),
			Data:   exception.Error(), // ? exception.error => digunakan untuk mengambil pesan error
		}

		ctx.JSON(webResponse)
		return true
	} else {
		return false
	}
}

func InternalServerError(ctx fiber.Ctx, err any) {
	ctx.Set("Content-Type", "application/json")
	ctx.Status(http.StatusInternalServerError)

	webResponse := response.WebResponse{
		Code:   http.StatusInternalServerError, // 500
		Status: http.StatusText(http.StatusInternalServerError),
		Data:   err,
	}

	ctx.JSON(webResponse)
}
