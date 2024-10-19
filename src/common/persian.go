package common

import (
	"log"
	"regexp"
)

func MobileNumberValidate(value string) bool {
	res, err := regexp.MatchString(`(0|\+98)?([ ]|-|[()]){0,2}9[1|2|3|4]([ ]|-|[()]){0,2}(?:[0-9]([ ]|-|[()]){0,2}){8}`, value)
	if err != nil {
		log.Print(err.Error())
	}

	return res
}
