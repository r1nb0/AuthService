package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/r1nb0/UserService/configs"
	"github.com/r1nb0/UserService/internal/utils"
	"log"
)

func checkPassword(pass string) bool {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("error of getting config: %s", err.Error())
	}
	if len(pass) < cfg.Password.MinLength {
		return false
	}
	if cfg.Password.IncludeChars && !utils.HasLetter(pass) {
		return false
	}
	if cfg.Password.IncludeDigits && !utils.HasDigit(pass) {
		return false
	}
	if cfg.Password.IncludeLowercase && !utils.HasLower(pass) {
		return false
	}
	if cfg.Password.IncludeUppercase && !utils.HasUpper(pass) {
		return false
	}
	if cfg.Password.IncludeSpecial && !utils.HasSpecial(pass) {
		return false
	}
	return true
}

func PasswordValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}
	return checkPassword(value)
}
