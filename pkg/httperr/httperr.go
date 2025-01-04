package httperr

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type httperr struct {
	type_    string
	status   int
	detail   string
	instance string
	cause    error
}

func (h *httperr) Error() string {
	if h.type_ != "" && h.status == 0 {
		h.status = statusOf(h.type_)
	}

	jsonData, err := json.Marshal(h)
	if err != nil {
		return fmt.Sprintf("failed to marshal httperr: %v", err)
	}

	return string(jsonData)
}

func New(err ...error) *httperr {
	httperr := &httperr{}
	if len(err) > 0 {
		httperr.cause = err[0]
	}

	return httperr
}

func (h *httperr) SetType(type_ string) *httperr {
	h.type_ = type_
	return h
}

func (h *httperr) SetStatus(status int) *httperr {
	h.status = status
	return h
}

func (h *httperr) SetDetail(format string, a ...any) *httperr {
	h.detail = fmt.Sprintf(format, a...)
	return h
}

func (h *httperr) SetInstance(instance string) *httperr {
	h.instance = instance
	return h
}

func (h *httperr) SetCause(cause error) *httperr {
	h.cause = errors.Wrap(cause, "httperr cause")
	return h
}

func (h *httperr) MarshalJSON() ([]byte, error) {
	type alias struct {
		Type_    string `json:"type"`
		Status   int    `json:"status"`
		Detail   string `json:"detail"`
		Instance string `json:"instance"`
	}

	return json.Marshal(alias{
		Type_:    h.type_,
		Status:   h.status,
		Detail:   h.detail,
		Instance: h.instance,
	})
}

func (h *httperr) Cause() error {
	return h.cause
}

func As(err error) (*httperr, bool) {
	if err == nil {
		return nil, false
	}

	if httpErr, ok := err.(*httperr); ok {
		return httpErr, true
	}

	return nil, false
}
