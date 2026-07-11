package service

import (
	"GO-AUTH-JWT/models/dto/request"
	"GO-AUTH-JWT/models/dto/response"
	"context"
	"mime/multipart"
)

// ! event_service.go itu untuk apa?

// ● File khusus logika bisnis untuk event
// ● Mengatur proses data event
// ● Menghubungkan handler dengan repository
// ● Biasanya berisi:
// ● validasi data event
// ● pemanggilan repository
// ● aturan bisnis (misalnya cek duplikat, dll)

type EventService interface {
	Create(ctx context.Context, request request.EventCreateRequest, fileImage *multipart.FileHeader) response.EventResponse
	Update(ctx context.Context, request request.EventUpdateRequest, fileImage *multipart.FileHeader) response.EventResponse
	Delete(ctx context.Context, eventId int64, userId int64)
	FindById(ctx context.Context, eventId int64, userId int64) response.EventResponse
	FindAll(ctx context.Context, userId int64) []response.EventResponse
}
