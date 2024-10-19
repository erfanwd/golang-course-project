package validations

import (
	"github.com/erfanwd/golang-course-project/common"
	"github.com/go-playground/validator/v10"
)

func PasswordValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}

	return common.PasswordValidate(value)
}
