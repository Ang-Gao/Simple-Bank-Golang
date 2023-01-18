package api

import (
	"myproject/util"

	"github.com/go-playground/validator/v10"
)

//use reflection to analyze the `tags`
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
