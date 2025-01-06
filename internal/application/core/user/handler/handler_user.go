package handler

import (
	"samsamoohooh-api/internal/application/port"
	"samsamoohooh-api/internal/application/presenter"
	"samsamoohooh-api/internal/infra/utils/handlerutil"
	"samsamoohooh-api/internal/infra/validator"
	"samsamoohooh-api/internal/router"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.UserService
	validator   *validator.Validator
	router      *router.Router
}

func NewUserHandler(
	userService port.UserService,
	validator *validator.Validator,
	router *router.Router,
) *UserHandler {
	userHandler := &UserHandler{
		userService: userService,
		validator:   validator,
		router:      router,
	}

	userHandler.Route(router.V0)

	return userHandler
}

func (h *UserHandler) Route(r fiber.Router) {
	users := r.Group("/users")
	{
		users.Get("/:id", h.FindUser)
		users.Patch("/:id", h.PatchUser)
	}
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

func (h *UserHandler) PatchUser(c *fiber.Ctx) error {
	req, err := handlerutil.Bind[presenter.PatchUserRequest](c)
	if err != nil {
		return err
	}

	err = h.validator.Validate(req)
	if err != nil {
		return err
	}

	patchedUser, err := h.userService.PatchUser(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewFindUserResponse(patchedUser))
}
