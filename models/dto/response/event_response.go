package response

import "time"

// ! package dto itu apa?

// ● Tempat struktur data untuk request & response API
// ● Penghubung antara client dan backend
// ● Biasanya berisi:
// ● struct request (input dari user)
// ● struct response (output ke client)
// ● DTO (Data Transfer Object)

type EventResponse struct {
	ID          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
}
