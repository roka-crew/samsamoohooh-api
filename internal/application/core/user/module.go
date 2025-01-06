package user

import (
	"samsamoohooh-api/internal/application/core/user/handler"
	"samsamoohooh-api/internal/application/core/user/service"
	"samsamoohooh-api/internal/application/core/user/store"
	"samsamoohooh-api/internal/application/port"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"user_module",
	fx.Provide(
		handler.NewUserHandler,

		fx.Annotate(
			service.NewUserService,
			fx.As(new(port.UserService)),
		),

		fx.Annotate(
			store.NewUserStore,
			fx.As(new(port.UserStore)),
		),
	),

	fx.Invoke(func(h *handler.UserHandler) {}),
)
