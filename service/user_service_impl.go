package service

import (
	"GO-AUTH-JWT/exception"
	"GO-AUTH-JWT/helper"
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"GO-AUTH-JWT/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// ! user_service_impl.go itu untuk apa?
// ● File implementasi dari interface service
// ● Tempat isi nyata logika bisnis user
// ● Mengolah data sebelum/sesudah ke repository
// ● Biasanya berisi:
// ● validasi input user
// ● pemanggilan user_repository
// ● pengaturan transaksi (tx)
// ● aturan bisnis user

type UserServiceImpl struct {
	UserRepository repository.UserRepository // ? tidak perlu dibuat pointer karena merupakan interface
	DB             *gorm.DB                  // ? dibuat pointer supaya tidak membuat salinan dari objek DB, melainkan menggunakan referensi yang sama
	validate       *validator.Validate       // perlu dibuat pointer, karena ia merupakan struct, pacakge go-playground/validator digunakan untuk validasi,
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		validate:       validate,
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) (_ response.UserResponse) {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindById(ctx, service.DB, int64(request.ID))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) // ? exception.NewNotFoundError(err.Error()) => digunakan untuk mengecek apakah category dengan id tersebut ada di db, err.Error() => digunakan untuk mengambil pesan error
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password

	user = service.UserRepository.Update(ctx, service.DB, &user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int64) {
	user, err := service.UserRepository.FindById(ctx, service.DB, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error())) // ? exception.NewNotFoundError(err.Error()) => digunakan untuk mengecek apakah category dengan id tersebut ada di db, err.Error() => digunakan untuk mengambil pesan error
	}
	service.UserRepository.Delete(ctx, service.DB, user.ID)
}
