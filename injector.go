//go:build wireinject
// +build wireinject

package main

import (
	"GO-AUTH-JWT/app"
	"GO-AUTH-JWT/controller"
	"GO-AUTH-JWT/repository"
	"GO-AUTH-JWT/service"
	"GO-AUTH-JWT/storage"

	"github.com/go-playground/validator/v10"

	"github.com/google/wire"
)

var userRepository = repository.NewUserRepository

var authSet = wire.NewSet(
	service.NewAuthService,
	controller.NewAuthController,
)

var userSet = wire.NewSet(
	service.NewUserService,
	controller.NewUserController,
)

var eventSet = wire.NewSet(
	repository.NewEventRepository,
	service.NewEventService,
	storage.NewImagekitUploader,
	storage.InitImageKit,
	controller.NewEventController,
)

func InitializedServer() *app.Server {
	wire.Build(
		app.NewServer,
		app.OpenConnection,
		validator.New,                    // ? validator.New digunakan untuk inisialisasi validator
		wire.Value([]validator.Option{}), // ? wire.Value(value interface{}) dipakai untuk memasukkan nilai langsung sebagai dependency, []validator.Option{} artinya memasukkan []validator.Option{} sebagai dependency ke wire
		userRepository,
		authSet,
		userSet,
		eventSet,

		app.BuildServer,
	)
	return nil
}
