package main

import (
	userHandler "samsamoohooh-api/internal/user/handler"
	userService "samsamoohooh-api/internal/user/service"
	userStore "samsamoohooh-api/internal/user/store"
	"samsamoohooh-api/pkg/config"
	"samsamoohooh-api/pkg/mysql"
	"samsamoohooh-api/pkg/token"
	"samsamoohooh-api/router"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Supply("configs/env.yaml"),
		fx.Provide(
			config.New,
			mysql.New,
			token.New,

			userHandler.NewUserHandler,
			userService.NewUserService,
			userStore.NewUserStore,

			router.New,
		),
		fx.Invoke(
			func(m *mysql.Mysql) {},

			userHandler.NewUserHandler,

			func(r *router.Router) {},
		),
	).Run()
}
