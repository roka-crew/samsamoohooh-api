package validator

import (
	"fmt"
	"samsamoohooh-api/pkg/httperr"
	"strings"

	stdv "github.com/go-playground/validator/v10"
)

type Validator struct {
	engine *stdv.Validate
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(out any) error {
	err := v.engine.Struct(out)
	if err != nil {
		validationErrors, ok := err.(stdv.ValidationErrors)
		if !ok {
			return httperr.New(err).
				SetType(httperr.ValidationFailed)
		}

		var details []string
		for _, err := range validationErrors {
			details = append(details, fmt.Sprintf(
				"Field: %s, Tag: %s, Value: %v",
				err.Field(),
				err.Tag(),
				err.Value(),
			))
		}

		return httperr.New(err).
			SetType(httperr.ValidationFailed).
			SetDetail("validation field: %s", strings.Join(details, "; "))
	}

	return nil
}

func (v *Validator) ValidateParams(out any) error {
	err := v.engine.Struct(out)
	if err != nil {
		validationErrors, ok := err.(stdv.ValidationErrors)
		if !ok {
			return httperr.New(err).
				SetStatus(httperr.StatusInternalServerError).
				SetType(httperr.ValidationFailed)
		}

		var details []string
		for _, err := range validationErrors {
			details = append(details, fmt.Sprintf(
				"Field: %s, Tag: %s, Value: %v",
				err.Field(),
				err.Tag(),
				err.Value(),
			))
		}

		return httperr.New(err).
			SetType(httperr.ValidationFailed).
			SetStatus(httperr.StatusInternalServerError).
			SetDetail("failed params valdiation, validation field: %s", strings.Join(details, "; "))
	}

	return nil
}
