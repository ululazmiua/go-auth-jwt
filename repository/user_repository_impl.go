package repository

import (
	"GO-AUTH-JWT/helper"
	"GO-AUTH-JWT/models/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

// ! implentasi repository
// user_repository_impl.go itu untuk apa?
// ● File implementasi dari interface repository
// ● Tempat isi nyata akses database user
// ● Menjalankan query SQL / ORM
// ● Biasanya berisi:
// ● real query SELECT
// ● INSERT data user
// ● UPDATE data user
// ● DELETE data user

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository { // ? digunakan untuk membuat object UserRepositoryImpl yang mengimplementasikan interface UserRepository
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, db *gorm.DB, user *domain.User) domain.User {
	User := domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	err := db.WithContext(ctx).Create(&User).Error
	helper.PanicIfError(err)

	return User
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, db *gorm.DB, email string) (_ domain.User, _ error) {
	user := domain.User{}
	err := db.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, id int64) (_ domain.User, _ error) {
	user := domain.User{}
	err := db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, db *gorm.DB, user *domain.User) (_ domain.User) {

	err := db.WithContext(ctx).Save(&user).Error
	helper.PanicIfError(err)

	return *user // ? kenapa di return *user, bukan user? karena user adalah pointer, jadi harus di dereference dulu
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id int64) {
	err := db.WithContext(ctx).Delete(&domain.User{}, "id = ?", id).Error
	helper.PanicIfError(err)
}
