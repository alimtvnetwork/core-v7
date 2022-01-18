package reflectinternal

import "reflect"

func SafeTypeName(any interface{}) string {
	rt := reflect.TypeOf(any)

	if IsNull(rt) {
		return ""
	}

	return rt.String()
}
