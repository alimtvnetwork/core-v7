package coredynamic

import "reflect"

func TypeName(any interface{}) string {
	rf := reflect.TypeOf(any)

	if rf == nil {
		return ""
	}

	return rf.String()
}
