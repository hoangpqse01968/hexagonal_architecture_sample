//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/application/repository"
	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/application/service"
	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/handler"
	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/infrastructure"

	"github.com/google/wire"
)

func BuildUserHandler(ctx context.Context) *handler.GetUserHandler {
	panic(wire.Build(
		wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
		service.NewUserService,
		wire.Bind(new(repository.UserRepository), new(*infrastructure.UserInfrastructure)),
		infrastructure.NewUserInfrastructure,
		handler.NewGetUserHandler,
	))
}
