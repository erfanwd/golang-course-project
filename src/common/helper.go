package common


import "github.com/jinzhu/copier"


func StructToStructMapper[T any](data any) (T, error) {
	var result T
	err := copier.Copy(&result, data)
	if err != nil {
		return result, err
	}
	return result, nil
}