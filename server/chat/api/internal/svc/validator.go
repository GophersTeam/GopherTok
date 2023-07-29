package svc

import (
	"github.com/bytedance/go-tagexpr/v2/validator"
	"net/http"
)

type Validator struct {
	validator *validator.Validator
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New("vd"),
	}
}

func (v Validator) Validate(r *http.Request, data any) error {
	return v.validator.Validate(data)
}
