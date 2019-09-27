package common

import (
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate will perform struct level validation on the struct.
func Validate(s interface{}) error {
	return validate.Struct(s)
}
