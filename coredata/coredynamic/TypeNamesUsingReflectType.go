package coredynamic

import "reflect"

func TypeNamesUsingReflectType(
	isFullName bool,
	reflectTypes ...reflect.Type,
) []string {
	slice := make([]string, len(reflectTypes))

	if isFullName {
		for i, item := range reflectTypes {
			slice[i] = item.String()
		}

		return slice
	}

	for i, item := range reflectTypes {
		slice[i] = item.Name()
	}

	return slice
}
