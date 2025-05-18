package handler

import "github.com/go-playground/validator/v10"

type RequestValidator struct {
	validator *validator.Validate
}

func (val *RequestValidator) Validate(i any) error {
	return val.validator.Struct(i)
}
