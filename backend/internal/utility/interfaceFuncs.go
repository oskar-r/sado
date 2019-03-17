package utility

import (
	"reflect"
)

func ClaimToInt(claim interface{}, claimID string) (int, bool) {
	v := reflect.Indirect(reflect.ValueOf(claim))
	var cl int
	ok := true
	if v.FieldByName(claimID).IsValid() {
		switch v.FieldByName(claimID).Interface().(type) {
		case int:
			cl = int(v.FieldByName(claimID).Int())
		default:
			ok = false
		}
	}
	return cl, ok
}
func InterfaceToIntSlice(v interface{}) ([]int, bool) {
	i := []int{}
	ok := true
	switch v := v.(type) {
	case []int:
		for _, value := range v {
			i = append(i, int(value))
		}
	default:
		ok = false
	}
	return i, ok
}
func ClaimToString(claim interface{}, claimID string) (string, bool) {
	v := reflect.Indirect(reflect.ValueOf(claim))
	var cl string
	ok := true
	if v.FieldByName(claimID).IsValid() {
		switch v.FieldByName(claimID).Interface().(type) {
		case string:
			cl = v.FieldByName(claimID).String()
		default:
			ok = false
		}
	}
	return cl, ok
}
