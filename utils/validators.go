package utils

import "github.com/go-playground/validator"

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateModel(model interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
