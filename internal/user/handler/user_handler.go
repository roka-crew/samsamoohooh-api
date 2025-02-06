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
		users.PATCH("/me", handler.PatchByMeUser)
	}

	return handler
}

// FindUserByMe godoc
//
//	@Summary		Find user by me
//	@Tags			User
//	@Description	Find user by me - ✅
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	presenter.FindUserByMeResponse	"사용자 조회 성공"
//	@Router			/users/me [get]
//	@Security		BearerAuth
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

	switch {
	case err == nil:
		return c.JSON(http.StatusOK, res.FromModel(foundUser))
	default:
		return err
	}
}

// PatchByMeUser godoc
//
//	@Summary		Patch user by me
//	@Tags			User
//	@Description	Patch user by me - ✅
//	@Accept			json
//	@Produce		json
//	@Param			PatchUserByMeRequest	body	domain.PatchUserByMeRequest	true	"사용자 수정 요청"
//	@Success		204						"사용자 수정 성공"
//	@Router			/users/me [patch]
//	@Security		BearerAuth
func (h userHandler) PatchByMeUser(c echo.Context) error {
	var (
		err error
		req domain.PatchUserByMeRequest
	)

	if err := c.Bind(&req); err != nil {
		return err
	}

	req.RequestUserID, err = handlerutil.GetRequestUserID(c)
	if err != nil {
		return err
	}

	err = h.userService.PatchByMeUser(c.Request().Context(), req)

	switch {
	case err == nil:
		return c.NoContent(http.StatusNoContent)
	default:
		return err
	}
}
