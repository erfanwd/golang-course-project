package validations

import (
	"github.com/erfanwd/golang-course-project/common"
	"github.com/go-playground/validator/v10"
)

func MobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.MobileNumberValidate(value)
}
