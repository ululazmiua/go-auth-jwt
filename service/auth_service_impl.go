package service

import (
	"GO-AUTH-JWT/exception"
	"GO-AUTH-JWT/helper"
	"GO-AUTH-JWT/models/domain"
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"GO-AUTH-JWT/repository"
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

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

	hashedPassword, err := bcrypt.GenerateFromPassword( // ? bcrypt.GenerateFromPassword(password []byte, cost int) ([]byte, error) => digunakan untuk mengenkripsi password menggunakan algoritma bcrypt, cost digunakan untuk menentukan kompleksitas algoritma misalnya 10 atau 12 yang artinya password akan dienkripsi sebanyak 10 atau 12 kali
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
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindByEmail(ctx, service.DB, request.Email)
	if err != nil {
		panic(exception.NewInvalidEmailPassword("invalid email or password"))
	}
	// ! Add password validation logic here
	err = bcrypt.CompareHashAndPassword( // ? bcrypt.CompareHashAndPassword(hashedPassword []byte, password []byte) error => digunakan untuk membandingkan password yang diinputkan user dengan password yang ada di database, jika password sama maka akan mengembalikan nil, jika tidak maka akan mengembalikan error
		[]byte(user.Password),
		[]byte(request.Password),
	)
	if err != nil {
		panic(exception.NewInvalidEmailPassword("invalid email or password"))
	}

	// ! JWT logic
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // ? jwt.NewWithClaims(signingMethod jwt.SigningMethod, claims jwt.Claims) *Token => digunakan untuk membuat token baru dengan signing method dan claims yang diberikan, jika signing method adalah HS256 maka claims harus bertipe jwt.MapClaims, jika signing method adalah RS256 maka claims harus bertipe jwt.RegisteredClaims. Jwt.MapClaims() => Mengisi informasi yang ingin disimpan di dalam JWT(misal id dan sebagainya)
		"sub": user.ID,                                   // ? sub digunakan untuk menyimpan id user
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // ? exp digunakan untuk menyimpan waktu kadaluarsa token
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))) // ? SignedString(key []byte) (string, error) => digunakan untuk menambahkan signature ke token yaitu JWT_SECRET difile .env
	helper.PanicIfError(err)

	return helper.ToLoginResponse(user, tokenString)
}
