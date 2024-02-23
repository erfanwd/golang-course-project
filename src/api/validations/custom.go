package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func MobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	res, err := regexp.MatchString(`(0|\+98)?([ ]|-|[()]){0,2}9[1|2|3|4]([ ]|-|[()]){0,2}(?:[0-9]([ ]|-|[()]){0,2}){8}`, value)
	if err != nil {
		return false		
	}
	
	return res
}