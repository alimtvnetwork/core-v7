package corefuncs

import "gitlab.com/auk-go/core/internal/reflectinternal"

func GetFuncName(i interface{}) string {
	return reflectinternal.GetFunc.Name(i)
}
