package router

import (
	"context"
	"log"
	"samsamoohooh-api/pkg/config"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"

	_ "samsamoohooh-api/docs/swagger"
)

type Router struct {
	cfg *config.Config
	*echo.Echo
}

func New(
	lc fx.Lifecycle,
	cfg *config.Config,
) *Router {
	r := &Router{
		cfg: cfg,
	}

	app := echo.New()
	// app.Use(middleware.Recover())

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	app.HTTPErrorHandler = func(err error, c echo.Context) {
		c.Logger().Error(err)
		c.JSON(500, map[string]interface{}{
			"message": "internal server error",
		})
	}

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
			if err := r.Echo.Shutdown(ctx); err != nil {
				return err
			}

			return nil
		},
	})

	r.Echo = app
	return r
}

func (r Router) listen() error {
	return r.Echo.Start(r.cfg.Listen)
}
