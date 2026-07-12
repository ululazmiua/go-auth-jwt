package repository

import (
	"GO-AUTH-JWT/helper"
	"GO-AUTH-JWT/models/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

// ! implentasi repository
// event_repository_impl.go itu untuk apa?
// ● File implementasi dari interface repository
// ● Tempat isi nyata akses database event
// ● Menjalankan query SQL / ORM
// ● Biasanya berisi:
// ● real query SELECT
// ● INSERT data event
// ● UPDATE data event
// ● DELETE data event

type EventRepositoryImpl struct {
}

func NewEventRepository() EventRepository {
	return &EventRepositoryImpl{}
}

func (repository *EventRepositoryImpl) Save(ctx context.Context, db *gorm.DB, event *domain.Event, userID int64) domain.Event {
	var user domain.User
	err := db.WithContext(ctx).First(&user, userID).Error
	helper.PanicIfError(err)

	newEvent := domain.Event{
		UserID:      user.ID,
		Name:        event.Name,
		Description: event.Description,
		Image:       event.Image,
		ImageId:     event.ImageId,
		Location:    event.Location,
		DateTime:    event.DateTime,
	}

	err = db.WithContext(ctx).Create(&newEvent).Error
	helper.PanicIfError(err)

	return newEvent
}

func (repository *EventRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, id int64, UserId int64) (_ domain.Event, _ error) {
	var Event domain.Event
	err := db.WithContext(ctx).First(&Event, "id = ? AND user_id = ?", id, UserId).Error
	if err != nil {
		return Event, errors.New("event not found")
	}

	return Event, nil
}

func (repository *EventRepositoryImpl) Update(ctx context.Context, db *gorm.DB, event *domain.Event, UserId int64) (_ domain.Event) {
	var Event domain.Event
	err := db.WithContext(ctx).Where("id = ? AND user_id = ?", event.ID, UserId).First(&Event).Error
	helper.PanicIfError(err)

	Event.Name = event.Name
	Event.Description = event.Description
	Event.Location = event.Location
	Event.Image = event.Image
	Event.ImageId = event.ImageId
	Event.DateTime = event.DateTime
	Event.UserID = event.UserID

	err = db.WithContext(ctx).Save(&Event).Error
	helper.PanicIfError(err)

	return Event
}

func (repository *EventRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id int64, UserId int64) {
	var Event domain.Event
	err := db.WithContext(ctx).Where("id = ? AND user_id = ?", id, UserId).First(&Event).Error
	helper.PanicIfError(err)

	err = db.WithContext(ctx).Delete(&Event).Error
	helper.PanicIfError(err)
}

func (repository *EventRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, UserId int64) (_ []domain.Event) {
	var Events []domain.Event
	err := db.WithContext(ctx).Where("user_id = ?", UserId).Find(&Events).Error
	helper.PanicIfError(err)

	return Events
}
