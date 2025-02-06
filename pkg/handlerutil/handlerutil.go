package handlerutil

import (
	"errors"
	"samsamoohooh-api/pkg/token"

	"github.com/labstack/echo/v4"
)

const (
	Token = "handlerutil_token"
)

func Locals[T any](ctx echo.Context, key string, obj ...T) (T, error) {
	var zero T
	if len(obj) != 0 {
		ctx.Set(key, obj[0])
		return zero, nil
	}

	v, ok := ctx.Get(key).(T)
	if !ok {
		return zero, errors.New("key not found")
	}

	return v, nil
}

func GetRequestUserID(ctx echo.Context) (int, error) {
	payload, err := Locals[token.Payload](ctx, Token)
	if err != nil {
		return 0, err
	}

	return payload.UserID, nil
}
