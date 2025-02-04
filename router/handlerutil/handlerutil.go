package handlerutil

import "github.com/labstack/echo/v4"

const (
	Token = "handlerutil_token"
)

func Local(ctx echo.Context, obj any) {
	ctx.Set(Token, obj)
}

func Get[T any](ctx echo.Context) (T, bool) {
	t, ok := ctx.Get(Token).(T)
	return t, ok
}
