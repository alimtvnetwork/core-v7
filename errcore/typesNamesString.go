package errcore

import (
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func typesNamesString(
	anyItems ...interface{},
) string {
	slice := make([]string, len(anyItems))

	for i, item := range anyItems {
		slice[i] = reflect.TypeOf(item).Name()
	}

	return strings.Join(slice, constants.CommaSpace)
}
