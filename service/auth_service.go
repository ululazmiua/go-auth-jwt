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

type AuthService interface {
	Register(ctx context.Context, request request.UserCreateRequest) response.UserResponse
	Login(ctx context.Context, request request.UserLoginRequest) response.LoginResponse
}
