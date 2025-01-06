package main

import (
	"samsamoohooh-api/internal/infra/config"
	"samsamoohooh-api/internal/infra/router"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Supply("./configs/env.yaml"),
		fx.Provide(config.NewConfig),
		fx.Provide(router.NewRouter),
		fx.Invoke(func(*router.Router) {}),
	).Run()
}
