package utils

import (
	"errors"
	"reflect"
)

// Contain 存在包含
func Contain(needles interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == needles {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(needles)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
