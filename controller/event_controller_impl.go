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
	// userId := ctx.Locals("userId").(float64)	// ctx.Locals() => untuk mengambil data dari context
	// eventRequest.UserId = int64(userId)
	userId, err := strconv.Atoi(ctx.Params("userId"))
	helper.PanicIfError(err)

	eventRequest.UserId = int64(userId)

	eventResponse := controller.EventService.Create(ctx.Context(), eventRequest)

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
	// userId := ctx.Locals("userId").(float64)	// ctx.Locals() => untuk mengambil data dari context
	// eventRequest.UserId = int64(userId)
	userId, err := strconv.Atoi(ctx.Params("userId"))
	helper.PanicIfError(err)

	eventRequest.ID = int64(EventId)
	eventRequest.UserId = int64(userId)

	eventResponse := controller.EventService.Update(ctx.Context(), eventRequest)

	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   eventResponse,
	})

}

func (controller *EventControllerImpl) Delete(ctx fiber.Ctx) (_ error) {
	// ! ambil id user dari token
	UserId, err := strconv.Atoi(ctx.Params("userId"))
	EventId, err := strconv.Atoi(ctx.Params("eventId"))

	helper.PanicIfError(err)

	controller.EventService.Delete(ctx.Context(), int64(EventId), int64(UserId))
	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	})
}

func (controller *EventControllerImpl) FindById(ctx fiber.Ctx) (_ error) {
	// ! ambil id user dari token
	UserId, err := strconv.Atoi(ctx.Params("userId"))
	EventId, err := strconv.Atoi(ctx.Params("eventId"))

	helper.PanicIfError(err)

	eventResponse := controller.EventService.FindById(ctx.Context(), int64(EventId), int64(UserId))
	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   eventResponse,
	})
}

func (controller *EventControllerImpl) FindAll(ctx fiber.Ctx) (_ error) {
	UserId, err := strconv.Atoi(ctx.Params("userId"))
	helper.PanicIfError(err)

	EventsResponse := controller.EventService.FindAll(ctx.Context(), int64(UserId))

	return ctx.JSON(response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   EventsResponse,
	})
}
