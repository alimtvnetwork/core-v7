package coredynamic

import "reflect"

// ZeroSet
//
// Sets empty bytes to the struct or the value but don't make it nil.
//
// It only makes all fields to nil or zero values.
//
// Warning :
//   - Must be set as a pointer reflect value.
func ZeroSet(rvPointer reflect.Value) {
	elem := rvPointer.Elem()

	elem.Set(reflect.Zero(elem.Type()))
}
