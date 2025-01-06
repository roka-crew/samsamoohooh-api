package router

import (
	"context"
	"fmt"
	"log"
	"samsamoohooh-api/internal/infra/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Router struct {
	*fiber.App

	config *config.Config
}

func NewRouter(
	lc fx.Lifecycle,
	config *config.Config,
) *Router {
	app := fiber.New(fiber.Config{})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	return Router{
		config: config,
		App:    app,
	}.serve(lc)
}

func (r Router) serve(lc fx.Lifecycle) *Router {
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
			if err := r.Shutdown(); err != nil {
				return fmt.Errorf("error shutting down server: %v", err)
			}

			return nil
		},
	})

	return &r
}

func (r *Router) listen() error {
	return r.App.Listen(r.config.Server.Addr)
}
