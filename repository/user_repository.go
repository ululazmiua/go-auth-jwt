package repository

import (
	"GO-AUTH-JWT/models/domain"
	"context"

	"gorm.io/gorm"
)

// ! Repository hanya berisi query database.
// Repository tidak boleh berisi logika bisnis.

type UserRepository interface {
	Save(ctx context.Context, db *gorm.DB, user *domain.User) domain.User
	FindByEmail(ctx context.Context, db *gorm.DB, email string) (domain.User, error)
	FindById(ctx context.Context, db *gorm.DB, id int64) (domain.User, error)
	Update(ctx context.Context, db *gorm.DB, user *domain.User) domain.User
	Delete(ctx context.Context, db *gorm.DB, id int64)
	FindAll(ctx context.Context, db *gorm.DB) []domain.User
}
