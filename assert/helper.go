package assert

import "reflect"

func IsNil(obj interface{}) bool {
	if obj == nil {
		return true
	}

	value := reflect.TypeOf(obj)

	if value == nil {
		return true
	}
	return false
}
