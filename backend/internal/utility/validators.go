package utility

import (
	"reflect"
	"regexp"

	validator "gopkg.in/go-playground/validator.v8"
)

func FormidaName(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if ch, ok := field.Interface().(string); ok {
		re := regexp.MustCompile("^[\\p{L}_\\-\\.0-9 ]{1,}$")
		if re.MatchString(ch) {
			return true
		}

		return false
	}
	return false
}
