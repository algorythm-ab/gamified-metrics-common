package methods

import "github.com/go-playground/validator/v10"

// Validations implements validator.Func
func Validations(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}
