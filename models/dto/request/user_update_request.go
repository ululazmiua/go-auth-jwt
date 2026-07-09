package request

// ! package dto itu apa?

// ● Tempat struktur data untuk request & response API
// ● Penghubung antara client dan backend
// ● Biasanya berisi:
// ● struct request (input dari user)
// ● struct response (output ke client)
// ● DTO (Data Transfer Object)

type UserUpdateRequest struct {
	ID       int64    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required,min=1,max=255"` // ? validate="required" artinya field ini wajib diisi, min=3 artinya minimal 3 karakter, max=255 artinya maksimal 255 karakter
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=0123456789,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
}
