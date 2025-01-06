package handler

import (
	"github.com/gofiber/fiber/v2"
	"samsamoohooh-api/internal/application/core/user/service"
	"samsamoohooh-api/internal/application/presenter"
	"samsamoohooh-api/internal/infra/utils/handlerutil"
	"samsamoohooh-api/internal/infra/validator"
)

type UserHandler struct {
	userService service.UserService
	validator   validator.Validator
}

func NewUserHandler(
	userService service.UserService,
	validator validator.Validator,
) *UserHandler {
	userHandler := &UserHandler{
		userService: userService,
		validator:   validator,
	}

	return userHandler
}

func (h *UserHandler) FindUser(c *fiber.Ctx) error {
	req, err := handlerutil.Bind[presenter.FindUserRequest](c)
	if err != nil {
		return err
	}

	err = h.validator.Validate(req)
	if err != nil {
		return err
	}

	foundUser, err := h.userService.FindUser(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewFindUserResponse(foundUser))
}
