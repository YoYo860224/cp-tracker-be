//go:build wireinject
// +build wireinject

package routes

import (
	"cp_tracker/db"
	"cp_tracker/user/controllers"
	"cp_tracker/user/services"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/wire"
)

type App struct {
	MongoDB        *mongo.Database
	UserService    *services.UserService
	UserController *controllers.UserController
}

func InitializeApp() (App, error) {
	wire.Build(
		db.ProvideMongoDB,
		services.NewUserService,
		controllers.NewUserController,
		wire.Struct(new(App), "*"),
	)
	return App{}, nil
}
