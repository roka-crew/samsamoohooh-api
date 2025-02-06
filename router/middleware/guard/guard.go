package guard

import (
	"net/http"
	"samsamoohooh-api/pkg/handlerutil"
	"samsamoohooh-api/pkg/token"
	"strings"

	"github.com/labstack/echo/v4"
)

type Guard struct {
	tokenService token.Token
}

func New(
	tokenService token.Token,
) *Guard {
	return &Guard{
		tokenService: tokenService,
	}
}

func (g Guard) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Get token from header
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "token is required",
			})
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": `'Bearer' token is required`,
			})
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Validate token
		payload, err := g.tokenService.ParseToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "valid token is required",
			})
		}

		// Set payload to context
		_, err = handlerutil.Locals(c, handlerutil.Token, payload)
		if err != nil {
			return err
		}

		return next(c)
	}
}
