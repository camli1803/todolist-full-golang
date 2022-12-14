package validator

import (
	validator "github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (*CustomValidator) New() *validator.Validate {
	var customValidator CustomValidator
	customValidator.Validator = validator.New()

	customValidator.Validator.RegisterValidation("password", IsFormatPassword)
	customValidator.Validator.RegisterValidation("todo_content", IsTodoContent)
	
	return customValidator.Validator
}
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
