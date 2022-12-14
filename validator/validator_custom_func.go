package validator

import (
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"	
)

func IsFormatPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var (
		hasNumber bool
		hasLetter bool
		hasSuitableLen bool		
	)
	
	if utf8.RuneCountInString(password) >= 6 && utf8.RuneCountInString(password) <= 30 {
		hasSuitableLen = true
	}

	for _, c := range password {
		if unicode.IsNumber(c) {
			hasNumber = true 
		}
		if unicode.IsLetter(c) {
			hasLetter = true
		}
	}

	return hasNumber && hasLetter && hasSuitableLen
}

func IsTodoContent(fl validator.FieldLevel) bool {
	content := fl.Field().String()
	var hasLetter = false
	for _, c := range content {
		if unicode.IsLetter(c) {
			hasLetter = true
		}
	}

	return hasLetter
}