package request

// ! package dto itu apa?

// ● Tempat struktur data untuk request & response API
// ● Penghubung antara client dan backend
// ● Biasanya berisi:
// ● struct request (input dari user)
// ● struct response (output ke client)
// ● DTO (Data Transfer Object)

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=0123456789,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
}
