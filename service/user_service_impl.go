package service

import (
	"GO-AUTH-JWT/helper"
	"GO-AUTH-JWT/models/domain"
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"GO-AUTH-JWT/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// ! category_service_impl.go itu untuk apa?
// ● File implementasi dari interface service
// ● Tempat isi nyata logika bisnis category
// ● Mengolah data sebelum/sesudah ke repository
// ● Biasanya berisi:
// ● validasi input category
// ● pemanggilan category_repository
// ● pengaturan transaksi (tx)
// ● aturan bisnis category

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

func (service *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) (_ response.UserResponse) {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user = service.UserRepository.Save(ctx, service.DB, &user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) (_ response.UserResponse) {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindById(ctx, service.DB, int64(request.ID))
	helper.PanicIfError(err)

	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password

	user = service.UserRepository.Update(ctx, service.DB, &user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int64) {
	service.UserRepository.Delete(ctx, service.DB, userId)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int64) (_ response.UserResponse) {
	user, err := service.UserRepository.FindById(ctx, service.DB, userId)
	helper.PanicIfError(err)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindByEmail(ctx context.Context, email string) (_ response.UserResponse) {
	user, err := service.UserRepository.FindByEmail(ctx, service.DB, email)
	helper.PanicIfError(err)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) (_ []response.UserResponse) {
	users := service.UserRepository.FindAll(ctx, service.DB)

	return helper.ToUsersResponse(users)
}
