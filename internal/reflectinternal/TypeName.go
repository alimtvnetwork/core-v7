package reflectinternal

import (
	"reflect"
)

func TypeName(any interface{}) string {
	rfType := reflect.TypeOf(any)

	if rfType == nil {
		return ""
	}

	return rfType.String()
}
