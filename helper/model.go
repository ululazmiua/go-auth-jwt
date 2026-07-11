package helper

import (
	"GO-AUTH-JWT/models/domain"
	"GO-AUTH-JWT/models/dto/response"
)

// ! package Helper
// package helper itu apa?
// ● Tempat fungsi bantuan (utility) aplikasi
// ● Dipakai ulang di banyak bagian program
// ● Bukan logika bisnis utama
// ● Biasanya berisi:
// ● helper database (commit/rollback tx)
// ● helper response JSON
// ● helper error handling
// ● helper konversi data

func ToUserResponse(user domain.User) response.UserResponse {
	return response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUsersResponse(users []domain.User) []response.UserResponse {
	var usersResponse []response.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, ToUserResponse(user))
	}
	return usersResponse
}

func ToEventResponse(event domain.Event) response.EventResponse {
	return response.EventResponse{
		ID:          event.ID,
		UserId:      event.UserID,
		Name:        event.Name,
		Image:       event.Image,
		ImageId:     event.ImageId,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    event.DateTime,
	}
}

func ToEventsResponse(events []domain.Event) []response.EventResponse {
	var eventsResponse []response.EventResponse
	for _, event := range events {
		eventsResponse = append(eventsResponse, ToEventResponse(event))
	}
	return eventsResponse
}

func ToLoginResponse(user domain.User, token string) response.LoginResponse {
	return response.LoginResponse{
		Token: token,
		User:  ToUserResponse(user),
	}
}
