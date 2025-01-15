package main

import (
	"samsamoohooh-api/internal/application/core/user"
	"samsamoohooh-api/internal/infra"
	"samsamoohooh-api/internal/router"

	"go.uber.org/fx"

	_ "samsamoohooh-api/docs/swagger"
)

func main() {
	fx.New(
		fx.Supply("./configs/env.yaml"),
		infra.Module,
		user.Module,
		fx.Provide(router.NewRouter),
		fx.Invoke(func(*router.Router) {}),
	).Run()
}
