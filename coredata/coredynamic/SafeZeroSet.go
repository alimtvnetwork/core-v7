package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func SafeZeroSet(rv reflect.Value) {
	if reflectinternal.Is.Null(rv) {
		return
	}

	rv.Elem().Set(reflect.Zero(rv.Elem().Type()))
}
