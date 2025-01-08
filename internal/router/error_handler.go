package router

import (
	"samsamoohooh-api/pkg/httperr"

	"github.com/gofiber/fiber/v2"
)

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	if castedHttperr, ok := httperr.Cast(err); ok {
		castedHttperr.SetInstance(c.Path())
		return c.Status(castedHttperr.Status()).SendString(castedHttperr.Error())
	}

	return fiber.DefaultErrorHandler(c, err)
}
