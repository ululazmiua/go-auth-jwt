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
	"golang.org/x/crypto/bcrypt"
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

type AuthServiceImpl struct {
	UserRepository repository.UserRepository // ? tidak perlu dibuat pointer karena merupakan interface
	DB             *gorm.DB                  // ? dibuat pointer supaya tidak membuat salinan dari objek DB, melainkan menggunakan referensi yang sama
	validate       *validator.Validate       // perlu dibuat pointer, karena ia merupakan struct, pacakge go-playground/validator digunakan untuk validasi,
}

func NewAuthService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		validate:       validate,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request request.UserCreateRequest) (_ response.UserResponse) {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	_, err = service.UserRepository.FindByEmail(ctx, service.DB, request.Email)
	if err == nil {
		panic(exception.NewEmailAlreadyExistsError("email already exists"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	helper.PanicIfError(err)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	user = service.UserRepository.Save(ctx, service.DB, &user)

	return helper.ToUserResponse(user)
}

func (service *AuthServiceImpl) Login(ctx context.Context, request request.UserLoginRequest) (_ response.LoginResponse) {
	user, err := service.UserRepository.FindByEmail(ctx, service.DB, request.Email)
	if err != nil {
		panic(exception.NewUnauthorizedError("invalid email or password"))
	}
	// ! Add password validation logic here
	err = bcrypt.CompareHashAndPassword( // ? bcrypt.CompareHashAndPassword(hashedPassword []byte, password []byte) error => digunakan untuk membandingkan password yang diinputkan user dengan password yang ada di database, jika password sama maka akan mengembalikan nil, jika tidak maka akan mengembalikan error
		[]byte(user.Password),
		[]byte(request.Password),
	)
	if err != nil {
		panic(exception.NewUnauthorizedError("invalid email or password"))
	}

	// generate token logic
	// ! Add JWT token logic here
	// ! Add JWT token logic here
	// ! Add JWT token logic here
	// ! Add JWT token logic here
	// ! Add JWT token logic here
	// ! Add JWT token logic here
	return helper.ToLoginResponse(user, "token")
}
