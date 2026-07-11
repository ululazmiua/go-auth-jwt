package controller

import (
	"GO-AUTH-JWT/exception"
	"GO-AUTH-JWT/helper"
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"GO-AUTH-JWT/service"
	"net/http"

	"strconv"

	"github.com/gofiber/fiber/v3"
)

// ! event_controller_impl.go itu untuk apa?
// ● File implementasi dari interface controller
// ● Tempat isi nyata handler API event
// ● Menjalankan proses request–response
// ● Memanggil event_service
// ● Biasanya berisi:
// ● kode real untuk GET event
// ● POST event
// ● UPDATE event
// ● DELETE event

type EventControllerImpl struct {
	EventService service.EventService
}

func NewEventController(eventService service.EventService) EventController {
	return &EventControllerImpl{
		EventService: eventService,
	}
}

func (controller *EventControllerImpl) Create(ctx fiber.Ctx) (_ error) {
	eventRequest := request.EventCreateRequest{}
	err := ctx.Bind().Body(&eventRequest)
	helper.PanicIfError(err)

	// ! ambil id user dari token
	userId := ctx.Locals("userId").(float64) // ctx.Locals() => untuk mengambil data dari context
	eventRequest.UserId = int64(userId)

	// ! ambil file image dari form
	fileImage, err := ctx.FormFile("image")
	if err != nil {
		panic(exception.NewCustomBadRequestError("Gambar wajib ada!"))
	}

	eventResponse := controller.EventService.Create(ctx.Context(), eventRequest, fileImage)

	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   eventResponse,
	})
}

func (controller *EventControllerImpl) Update(ctx fiber.Ctx) (_ error) {
	eventRequest := request.EventUpdateRequest{}
	err := ctx.Bind().Body(&eventRequest) // ? ctx.Bind() => untuk memparsing request body, .Body(out any) => untuk menampung hasil parsing
	helper.PanicIfError(err)

	EventId, err := strconv.Atoi(ctx.Params("eventId"))
	helper.PanicIfError(err)

	// ! ambil id user dari token
	userId := ctx.Locals("userId").(float64) // ctx.Locals() => untuk mengambil data dari context
	eventRequest.UserId = int64(userId)

	eventRequest.ID = int64(EventId)
	eventRequest.UserId = int64(userId)

	// ! ambil file image dari form
	fileImage, err := ctx.FormFile("image")
	if err != nil {
		panic(exception.NewCustomBadRequestError("Gambar wajib ada!"))
	}

	eventResponse := controller.EventService.Update(ctx.Context(), eventRequest, fileImage)

	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   eventResponse,
	})

}

func (controller *EventControllerImpl) Delete(ctx fiber.Ctx) (_ error) {

	// ! ambil id user dari token
	userId := ctx.Locals("userId").(float64) // ctx.Locals() => untuk mengambil data dari context
	EventId, err := strconv.Atoi(ctx.Params("eventId"))

	helper.PanicIfError(err)

	controller.EventService.Delete(ctx.Context(), int64(EventId), int64(userId))
	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	})
}

func (controller *EventControllerImpl) FindById(ctx fiber.Ctx) (_ error) {
	// ! ambil id user dari token
	userId := ctx.Locals("userId").(float64) // ctx.Locals() => untuk mengambil data dari context

	EventId, err := strconv.Atoi(ctx.Params("eventId"))

	helper.PanicIfError(err)

	eventResponse := controller.EventService.FindById(ctx.Context(), int64(EventId), int64(userId))
	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   eventResponse,
	})
}

func (controller *EventControllerImpl) FindAll(ctx fiber.Ctx) (_ error) {
	// ! ambil id user dari token
	userId := ctx.Locals("userId").(float64) // ctx.Locals() => untuk mengambil data dari context

	EventsResponse := controller.EventService.FindAll(ctx.Context(), int64(userId))

	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   EventsResponse,
	})
}
