package infra

import (
	"samsamoohooh-api/internal/application/port"
	"samsamoohooh-api/internal/infra/authenticate/token"
	"samsamoohooh-api/internal/infra/config"
	"samsamoohooh-api/internal/infra/persistence/mysql"
	"samsamoohooh-api/internal/infra/validator"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"infra_module",
	fx.Provide(
		config.NewConfig,
		validator.NewValidator,
		mysql.NewMysql,
		fx.Annotate(
			token.NewTokenService,
			fx.As(new(port.TokenService)),
		),
	),
)
