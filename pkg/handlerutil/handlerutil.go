package handlerutil

import (
	"github.com/gofiber/fiber/v2"
	"samsamoohooh-api/pkg/httperr"
)

func Bind[T any](c *fiber.Ctx) (*T, error) {
	var req T
	if err := c.BodyParser(&req); err != nil {
		return nil, httperr.New(err).
			SetType(httperr.RequestParsingFailed).
			SetDetail("unable to parse request body")
	}

	if err := c.ParamsParser(&req); err != nil {
		return nil, httperr.New(err).
			SetType(httperr.RequestParsingFailed).
			SetDetail("unable to parse request params")
	}

	if err := c.QueryParser(&req); err != nil {
		return nil, httperr.New(err).
			SetType(httperr.RequestParsingFailed).
			SetDetail("unable to parse request query")
	}

	return &req, nil
}
