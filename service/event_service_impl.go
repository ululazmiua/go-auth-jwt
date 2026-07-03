package service

import (
	"GO-AUTH-JWT/exception"
	"GO-AUTH-JWT/helper"
	"GO-AUTH-JWT/models/domain"
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"GO-AUTH-JWT/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// ! event_service_impl.go itu untuk apa?
// ● File implementasi dari interface service
// ● Tempat isi nyata logika bisnis event
// ● Mengolah data sebelum/sesudah ke repository
// ● Biasanya berisi:
// ● validasi input event
// ● pemanggilan event_repository
// ● pengaturan transaksi (tx)
// ● aturan bisnis event

type EventServiceImpl struct {
	EventRepository repository.EventRepository
	DB              *gorm.DB
	validate        *validator.Validate
}

func NewEventService(eventRepository repository.EventRepository, DB *gorm.DB, validate *validator.Validate) EventService {
	return &EventServiceImpl{
		EventRepository: eventRepository,
		DB:              DB,
		validate:        validate,
	}
}

func (service *EventServiceImpl) Create(ctx context.Context, request request.EventCreateRequest) (_ response.EventResponse) {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	event := domain.Event{
		UserID:      request.UserId,
		Name:        request.Name,
		Description: request.Description,
		Location:    request.Location,
		DateTime:    request.DateTime,
	}
	event = service.EventRepository.Save(ctx, service.DB, &event, request.UserId)

	return helper.ToEventResponse(event)
}

func (service *EventServiceImpl) Update(ctx context.Context, request request.EventUpdateRequest) (_ response.EventResponse) {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	event, err := service.EventRepository.FindById(ctx, service.DB, request.ID, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) // ? exception.NewNotFoundError(err.Error()) => digunakan untuk mengecek apakah category dengan id tersebut ada di db, err.Error() => digunakan untuk mengambil pesan error
	}

	event.Name = request.Name
	event.Description = request.Description
	event.Location = request.Location
	event.DateTime = request.DateTime

	event = service.EventRepository.Update(ctx, service.DB, &event, request.UserId)

	return helper.ToEventResponse(event)
}

func (service *EventServiceImpl) Delete(ctx context.Context, eventId int64, userId int64) {
	event, err := service.EventRepository.FindById(ctx, service.DB, eventId, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) // ? exception.NewNotFoundError(err.Error()) => digunakan untuk mengecek apakah category dengan id tersebut ada di db, err.Error() => digunakan untuk mengambil pesan error
	}
	service.EventRepository.Delete(ctx, service.DB, event.ID, userId)
}

func (service *EventServiceImpl) FindById(ctx context.Context, eventId int64, userId int64) (_ response.EventResponse) {
	event, err := service.EventRepository.FindById(ctx, service.DB, eventId, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) // ? exception.NewNotFoundError(err.Error()) => digunakan untuk mengecek apakah category dengan id tersebut ada di db, err.Error() => digunakan untuk mengambil pesan error
	}

	return helper.ToEventResponse(event)
}

func (service *EventServiceImpl) FindAll(ctx context.Context, userId int64) (_ []response.EventResponse) {
	events := service.EventRepository.FindAll(ctx, service.DB, userId)

	return helper.ToEventsResponse(events)
}
