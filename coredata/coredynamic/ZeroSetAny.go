package coredynamic

import (
	"reflect"

	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

// ZeroSetAny
//
// Sets empty bytes to the struct or the value but don't make it nil.
//
// It only makes all fields to nil or zero values.
//
// Warning :
//  - Must be set as a pointer any.
func ZeroSetAny(anyItem interface{}) {
	if reflectinternal.IsNull(anyItem) {
		return
	}

	SafeZeroSet(reflect.ValueOf(anyItem))
}
