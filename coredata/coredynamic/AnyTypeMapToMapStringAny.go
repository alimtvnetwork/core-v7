package coredynamic

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

// AnyTypeMapToMapStringAny
//
//  expectation : map[key:interface{}]interface{} to map[string]interface{}
func AnyTypeMapToMapStringAny(
	reflectVal reflect.Value,
) (map[string]interface{}, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return AnyTypeMapToMapStringAny(
			reflect.Indirect(reflect.ValueOf(reflectVal)))
	}

	if reflectVal.Kind() != reflect.Map {
		return map[string]interface{}{},
			errcore.TypeMismatchType.
				Error("Reflection is not Map", reflectVal)
	}

	mapKeys := reflectVal.MapKeys()
	newMap := make(
		map[string]interface{},
		reflectVal.Len()+1)

	for _, key := range mapKeys {
		value := reflectVal.MapIndex(key)
		keyAny := key.Interface()
		var keyAsString string

		keyAsString, isString := keyAny.(string)
		if isString {
			newMap[keyAsString] = value.Interface()
			continue
		}

		keyAsString = fmt.Sprintf(
			constants.SprintValueFormat,
			keyAny)
		newMap[keyAsString] = value.Interface()
	}

	return newMap, nil
}
