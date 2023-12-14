package corefuncs

import (
	"runtime"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func GetFunc(i interface{}) *runtime.Func {
	return reflectinternal.GetFunc(i)
}
