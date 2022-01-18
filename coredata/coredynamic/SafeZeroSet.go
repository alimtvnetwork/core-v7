package coredynamic

import (
	"reflect"

	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

func SafeZeroSet(rv reflect.Value) {
	if reflectinternal.IsNull(rv) {
		return
	}

	rv.Elem().Set(reflect.Zero(rv.Elem().Type()))
}
