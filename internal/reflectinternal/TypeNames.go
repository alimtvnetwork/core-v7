package reflectinternal

import "reflect"

func TypeNames(
	isFullName bool,
	anyItems ...interface{},
) []string {
	slice := make([]string, len(anyItems))

	if isFullName {
		for i, item := range anyItems {
			slice[i] = reflect.TypeOf(item).String()
		}

		return slice
	}

	for i, item := range anyItems {
		slice[i] = reflect.TypeOf(item).Name()
	}

	return slice
}
