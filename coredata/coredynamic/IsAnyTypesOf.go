package coredynamic

import "reflect"

func IsAnyTypesOf(
	typeLookup reflect.Type,
	acceptedTypes ...reflect.Type,
) (isFound bool) {
	for _, acceptedType := range acceptedTypes {
		if acceptedType == typeLookup {
			return true
		}
	}

	return false
}
