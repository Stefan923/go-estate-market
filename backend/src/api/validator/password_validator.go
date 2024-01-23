package validator

import (
	"github.com/Stefan923/go-estate-market/config"
	"github.com/Stefan923/go-estate-market/util"
	"github.com/go-playground/validator/v10"
)

func PasswordValidator(field validator.FieldLevel) bool {
	value, success := field.Field().Interface().(string)
	if !success {
		field.Param()
		return false
	}

	return checkPassword(value)
}

func checkPassword(password string) bool {
	passwordConfig := config.GetConfig().Auth.Password
	if len(password) < passwordConfig.MinLength {
		return false
	}

	if passwordConfig.IncludeChars && !util.HasLetter(password) {
		return false
	}

	if passwordConfig.IncludeDigits && !util.HasDigits(password) {
		return false
	}

	if passwordConfig.IncludeLowercase && !util.HasLower(password) {
		return false
	}

	if passwordConfig.IncludeUppercase && !util.HasUpper(password) {
		return false
	}

	return true
}
