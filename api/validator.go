package api

import (
	"github.com/Federico191/Simple_Bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		//check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
