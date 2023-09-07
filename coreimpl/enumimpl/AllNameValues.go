package enumimpl

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/constants"
)

func AllNameValues(nameStrings []string, anyEnumVal interface{}) []string {
	reflectValues := reflect.ValueOf(anyEnumVal)
	length := reflectValues.Len()
	slice := make([]string, length)

	for i := 0; i < length; i++ {
		rfVal := reflectValues.Index(i)
		anyVal := rfVal.Interface()

		slice[i] = fmt.Sprintf(
			constants.EnumNameValueFormat,
			nameStrings[i],
			anyVal)
	}

	return slice
}
