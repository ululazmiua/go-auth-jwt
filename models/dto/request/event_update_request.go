package request

import "time"

// ! package dto itu apa?

// ● Tempat struktur data untuk request & response API
// ● Penghubung antara client dan backend
// ● Biasanya berisi:
// ● struct request (input dari user)
// ● struct response (output ke client)
// ● DTO (Data Transfer Object)

type EventUpdateRequest struct {
	ID          int64     `json:"id" validate:"required"`
	UserId      int64     `json:"user_id" validate:"required"`
	Name        string    `json:"name" validate:"required,min=1,max=255"` // ? validate="required" artinya field ini wajib diisi, min=3 artinya minimal 3 karakter, max=255 artinya maksimal 255 karakter
	Description string    `json:"description" validate:"required,min=1"`
	Location    string    `json:"location" validate:"required,min=1"`
	DateTime    time.Time `json:"date_time" validate:"required"`
}
