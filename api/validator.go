package api

import (
	"bank/utils"

	// "github.com/go-playground/locales/currency"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func=func(FieldLevel validator.FieldLevel) bool {
	if currency,ok:=FieldLevel.Field().Interface().(string);ok{
		return utils.IsSupportedCurrency(currency)
	}
	return false
}