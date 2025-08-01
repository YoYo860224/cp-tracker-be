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
	MongoDB            *mongo.Database
	UserService        *services.UserService
	UserController     *controllers.UserController
	UserDataService    *services.UserDataService
	UserDataController *controllers.UserDataController
	InviteCodeService  *services.InviteCodeService
}

func InitializeApp() (App, error) {
	wire.Build(
		db.ProvideMongoDB,
		services.NewUserService,
		controllers.NewUserController,
		services.NewUserDataService,
		controllers.NewUserDataController,
		services.NewInviteCodeService,
		wire.Struct(new(App), "*"),
	)
	return App{}, nil
}
