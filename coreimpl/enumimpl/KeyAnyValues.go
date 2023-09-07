package enumimpl

import "reflect"

func KeyAnyValues(names []string, values interface{}) []KeyAnyVal {
	if len(names) == 0 {
		return []KeyAnyVal{}
	}

	slice := make([]KeyAnyVal, len(names))
	reflectValues := reflect.ValueOf(values)
	for i, name := range names {
		rfVal := reflectValues.Index(i)
		anyVal := rfVal.Interface()
		slice[i] = KeyAnyVal{
			Key:      name,
			AnyValue: anyVal,
		}
	}

	return slice
}
