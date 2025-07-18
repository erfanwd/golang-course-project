package validations

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)


func UniqueValidator(db *gorm.DB) validator.Func {
	return func(fld validator.FieldLevel) bool {
		value := fld.Field().String()
		tableName := fld.Param()
		if tableName == "" || value == "" || db == nil {
			return false
		}

		fieldName := strings.ToLower(fld.StructFieldName())

		var exists bool
		query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = ?)", tableName, fieldName)
		err := db.Raw(query, value).Scan(&exists).Error
		if err != nil {
			fmt.Println("DB Error:", err)
			return false
		}

		fmt.Println("Value:", value, "Exists:", exists)

		return !exists
	}
}
