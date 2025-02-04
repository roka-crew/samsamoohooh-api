package router

import (
	"context"
	"log"
	"samsamoohooh-api/pkg/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

type Router struct {
	cfg *config.Config
	app *echo.Echo
}

func NewRouter(
	lc fx.Lifecycle,
	cfg *config.Config,
) *Router {
	r := &Router{
		cfg: cfg,
	}

	app := echo.New()
	app.Use(middleware.Recover())

	// app.HTTPErrorHandler = echo.New().DefaultHTTPErrorHandler

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := r.listen()
				if err != nil {
					log.Panicf("listen server error: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := r.app.Shutdown(ctx); err != nil {
				return err
			}

			return nil
		},
	})

	r.app = app
	return r
}

func (r Router) listen() error {
	return r.app.Start(r.cfg.Listen)
}
