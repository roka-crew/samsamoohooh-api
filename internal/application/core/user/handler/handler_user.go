package handler

import (
	"samsamoohooh-api/internal/application/port"
	"samsamoohooh-api/internal/application/presenter"
	"samsamoohooh-api/internal/infra/handlerutil"
	"samsamoohooh-api/internal/router"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.UserService
	handlerUtil *handlerutil.HandlerUtil
	router      *router.Router
}

func NewUserHandler(
	userService port.UserService,
	handlerUtil *handlerutil.HandlerUtil,
	router *router.Router,
) *UserHandler {
	userHandler := &UserHandler{
		userService: userService,
		handlerUtil: handlerUtil,
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
	var req = &presenter.FindUserRequest{}
	if err := h.handlerUtil.Bind(c, req); err != nil {
		return nil
	}

	foundUser, err := h.userService.FindUser(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewFindUserResponse(foundUser))
}

func (h *UserHandler) PatchUser(c *fiber.Ctx) error {
	var req = &presenter.PatchUserRequest{}
	if err := h.handlerUtil.Bind(c, req); err != nil {
		return err
	}

	patchedUser, err := h.userService.PatchUser(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewFindUserResponse(patchedUser))
}
