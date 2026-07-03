package repository

import (
	"GO-AUTH-JWT/models/domain"
	"context"

	"gorm.io/gorm"
)

// ! Repository hanya berisi query database.
// Repository tidak boleh berisi logika bisnis.

type UserRepository interface {
	Save(ctx context.Context, db *gorm.DB, user *domain.User) domain.User            // ? untuk register user
	FindByEmail(ctx context.Context, db *gorm.DB, email string) (domain.User, error) // ? untuk login
	FindById(ctx context.Context, db *gorm.DB, id int64) (domain.User, error)        // ? (untuk update/delete dan validasi)
	Update(ctx context.Context, db *gorm.DB, user *domain.User) domain.User
	Delete(ctx context.Context, db *gorm.DB, id int64)
}
