package handlerutil

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"samsamoohooh-api/internal/infra/validator"
	"samsamoohooh-api/pkg/httperr"
)

type HandlerUtil struct {
	validator *validator.Validator
}

func NewHandlerUtil(
	validator *validator.Validator,
) *HandlerUtil {
	return &HandlerUtil{
		validator: validator,
	}
}

func (h *HandlerUtil) Bind(c *fiber.Ctx, out any) error {
	if err := c.BodyParser(out); err != nil && !errors.Is(err, fiber.ErrUnprocessableEntity) {
		return httperr.New(err).
			SetType(httperr.RequestParsingFailed).
			SetDetail("unable to parse request body")
	}

	if err := c.ParamsParser(out); err != nil {
		return httperr.New(err).
			SetType(httperr.RequestParsingFailed).
			SetDetail("unable to parse request params")
	}

	if err := c.QueryParser(out); err != nil && !errors.Is(err, fiber.ErrUnprocessableEntity) {
		return httperr.New(err).
			SetType(httperr.RequestParsingFailed).
			SetDetail("unable to parse request query")
	}

	if err := h.validator.Validate(out); err != nil {
		return err
	}

	return nil
}
