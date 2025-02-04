package errors

import "github.com/pkg/errors"

type Error struct {
	err error
}

func New(err any) *Error {
	switch err.(type) {
	case string:
		return &Error{err: errors.New(err.(string))}
	case error:
		return &Error{err: err.(error)}
	default:
		return &Error{err: errors.New("unknown error")}
	}
}

func (e Error) Error() string {
	return e.err.Error()
}
func (e Error) Unwrap() error {
	return errors.Unwrap(e.err)
}
