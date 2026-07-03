package repository

import (
	"GO-AUTH-JWT/models/domain"
	"context"

	"gorm.io/gorm"
)

// ! Repository hanya berisi query database.
// Repository tidak boleh berisi logika bisnis.

type EventRepository interface {
	Save(ctx context.Context, db *gorm.DB, event *domain.Event, UserId int64) domain.Event
	FindById(ctx context.Context, db *gorm.DB, id int64, UserId int64) (domain.Event, error)
	Update(ctx context.Context, db *gorm.DB, event *domain.Event, UserId int64) domain.Event
	Delete(ctx context.Context, db *gorm.DB, id int64, UserId int64)
	FindAll(ctx context.Context, db *gorm.DB, UserId int64) []domain.Event
}
