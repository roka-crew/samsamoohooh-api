package handler

import (
	"net/http"
	"samsamoohooh-api/domain"
	"samsamoohooh-api/internal/user/presenter"
	"samsamoohooh-api/pkg/handlerutil"
	"samsamoohooh-api/router"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	router      *router.Router
	userService domain.UserService
}

func NewUserHandler(
	router *router.Router,
	userService domain.UserService,
) *userHandler {
	handler := &userHandler{
		router:      router,
		userService: userService,
	}

	users := router.Group("/users")
	{
		users.GET("/me", handler.FindUserByMe)
	}

	return handler
}

func (h userHandler) FindUserByMe(c echo.Context) error {
	var (
		err error
		req domain.FindUserByMeRequest
		res presenter.FindUserByMeResponse
	)

	req.RequestUserID, err = handlerutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	foundUser, err := h.userService.FindUserByMe(c.Request().Context(), req)
	if err != nil {
		return err
	}

	switch err {
	case nil:
		return c.JSON(http.StatusOK, res.FromModel(*foundUser))
	default:
		return err
	}
}
