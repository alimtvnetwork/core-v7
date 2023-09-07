package coredynamic

import (
	"reflect"
)

// MapKeysStringSliceAny
//
//  expectation : map[key:string]don't care values
func MapKeysStringSliceAny(any interface{}) ([]string, error) {
	reflectVal := reflect.ValueOf(any)

	return MapKeysStringSlice(reflectVal)
}
