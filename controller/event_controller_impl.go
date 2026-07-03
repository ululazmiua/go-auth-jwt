package controller

import (
	"GO-AUTH-JWT/service"

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

func (controller *EventControllerImpl) Create(ctx *fiber.Ctx) (_ error) {
	panic("not implemented") // TODO: Implement
}

func (controller *EventControllerImpl) Update(ctx *fiber.Ctx) (_ error) {
	panic("not implemented") // TODO: Implement
}

func (controller *EventControllerImpl) Delete(ctx *fiber.Ctx) (_ error) {
	panic("not implemented") // TODO: Implement
}

func (controller *EventControllerImpl) FindById(ctx *fiber.Ctx) (_ error) {
	panic("not implemented") // TODO: Implement
}

func (controller *EventControllerImpl) FindAll(ctx *fiber.Ctx) (_ error) {
	panic("not implemented") // TODO: Implement
}
