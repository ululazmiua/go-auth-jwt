package service

import (
	"GO-AUTH-JWT/exception"
	"GO-AUTH-JWT/helper"
	"GO-AUTH-JWT/models/domain"
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"GO-AUTH-JWT/repository"
	"GO-AUTH-JWT/storage"
	"context"
	"fmt"
	"mime/multipart"
	"strconv"

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
	Storage         storage.Uploader
	validate        *validator.Validate
}

func NewEventService(eventRepository repository.EventRepository, DB *gorm.DB, storage storage.Uploader, validate *validator.Validate) EventService {
	return &EventServiceImpl{
		EventRepository: eventRepository,
		DB:              DB,
		Storage:         storage,
		validate:        validate,
	}
}

func (service *EventServiceImpl) Create(ctx context.Context, request request.EventCreateRequest, fileImage *multipart.FileHeader) (_ response.EventResponse) {
	file, err := fileImage.Open()
	if err != nil {
		panic(exception.NewCustomBadRequestError("Gambar wajib ada!"))
	}

	responseUpload, err := service.Storage.Upload(ctx, file, fileImage.Filename)
	request.Image = responseUpload.URL
	request.ImageId = responseUpload.FileIDImageKit

	err = service.validate.Struct(request)
	helper.PanicIfError(err)

	event := domain.Event{
		UserID:      request.UserId,
		Name:        request.Name,
		Description: request.Description,
		Image:       request.Image,
		ImageId:     request.ImageId,
		Location:    request.Location,
		DateTime:    request.DateTime,
	}
	event = service.EventRepository.Save(ctx, service.DB, &event, request.UserId)

	return helper.ToEventResponse(event)
}

func (service *EventServiceImpl) Update(ctx context.Context, request request.EventUpdateRequest, fileImage *multipart.FileHeader) response.EventResponse {
	// Validasi request
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	event, err := service.EventRepository.FindById(ctx, service.DB, request.ID, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Simpan id image lama
	oldImageID := event.ImageId
	fmt.Println("gambar lama: ", oldImageID)

	file, err := fileImage.Open()
	if err != nil {
		panic(exception.NewCustomBadRequestError("Gambar wajib ada"))
	}
	defer file.Close()

	// Upload gambar baru
	uploadResult, err := service.Storage.Upload(ctx, file, fileImage.Filename)
	if err != nil {
		panic(exception.NewCustomInternalServerError(err.Error()))
	}

	// Update object event
	event.Name = request.Name
	event.Description = request.Description
	event.Location = request.Location
	event.DateTime = request.DateTime
	event.Image = uploadResult.URL
	event.ImageId = uploadResult.FileIDImageKit

	// Simpan ke database
	event = service.EventRepository.Update(ctx, service.DB, &event, request.UserId)

	// Hapus gambar lama
	err = service.Storage.Delete(ctx, oldImageID)
	if err != nil {
		fmt.Println(err.Error())
	}

	return helper.ToEventResponse(event)
}

func (service *EventServiceImpl) Delete(ctx context.Context, eventId int64, userId int64) {
	event, err := service.EventRepository.FindById(ctx, service.DB, eventId, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) // ? exception.NewNotFoundError(err.Error()) => digunakan untuk mengecek apakah category dengan id tersebut ada di db, err.Error() => digunakan untuk mengambil pesan error
	}

	err = service.Storage.Delete(ctx, event.ImageId)
	if err != nil {
		panic(exception.NewCustomInternalServerError(err.Error()))
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

func (service *EventServiceImpl) FindAll(ctx context.Context, userId int64, querySearch string, page string, limit string) ([]response.EventResponse, int64) {
	// hitung total data
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		limitInt = 6
	}

	// hitung offset
	offset := (pageInt - 1) * limitInt // ? dikurang satu karena ingin menampilkan page ke berapa

	events, totalData := service.EventRepository.FindAll(ctx, service.DB, userId, querySearch, offset, limitInt)

	return helper.ToEventsResponse(events), totalData
}
