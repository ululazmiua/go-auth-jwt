package service

import (
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"context"
)

// ! user_service.go itu untuk apa?

// ● File khusus logika bisnis untuk user
// ● Mengatur proses data user
// ● Menghubungkan handler dengan repository
// ● Biasanya berisi:
// ● validasi data user
// ● pemanggilan repository
// ● aturan bisnis (misalnya cek duplikat, dll)

type UserService interface {
	Create(ctx context.Context, request request.UserCreateRequest) response.UserResponse
	Update(ctx context.Context, request request.UserUpdateRequest) response.UserResponse
	Delete(ctx context.Context, userId int64)
	FindById(ctx context.Context, userId int64) response.UserResponse
	FindByEmail(ctx context.Context, email string) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
}
