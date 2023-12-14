package isany

import (
	"reflect"
	"runtime"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// Function
//
// Returns false for nil, struct, anything else
func Function(item interface{}) (isFunc bool, name string) {
	if reflectinternal.Is.Null(item) {
		return false, ""
	}

	rv := reflect.ValueOf(item)

	if rv.Kind() != reflect.Func {
		return false, ""
	}

	name = runtime.FuncForPC(rv.Pointer()).Name()

	return name != "", name
}
