package reflectinternal

import (
	"reflect"
)

func TypeName(any interface{}) string {
	return reflect.TypeOf(any).String()
}
