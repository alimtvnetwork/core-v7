package coredynamic

import (
	"reflect"

	"gitlab.com/evatix-go/core/errcore"
)

func MapKeysStringSlice(reflectVal reflect.Value) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return MapKeysStringSlice(
			reflect.Indirect(reflect.ValueOf(reflectVal)))
	}

	if reflectVal.Kind() != reflect.Map {
		return []string{},
			errcore.TypeMismatch.Error("Reflection is not Map", reflectVal)
	}

	mapKeys := reflectVal.MapKeys()
	length := len(mapKeys)
	keys := make([]string, length)

	for i, key := range reflectVal.MapKeys() {
		keyAny := key.Interface()
		keyAsString, isString := keyAny.(string)

		if !isString {
			return keys, errcore.TypeMismatch.Error("Not string type", keyAny)
		}

		keys[i] = keyAsString
	}

	return keys, nil
}
