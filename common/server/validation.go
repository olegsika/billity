package server

import "github.com/go-playground/validator/v10"

// CustomValidator holds custom validator
type CustomValidator struct {
	V *validator.Validate
}

// Validate validates the request
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.V.Struct(i)
}
