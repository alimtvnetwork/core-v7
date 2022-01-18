package coredynamic

import (
	"reflect"

	"gitlab.com/evatix-go/core/constants"
)

func TypesIndexOf(
	typeLookup reflect.Type,
	acceptedTypes ...reflect.Type,
) (index int) {
	for i, acceptedType := range acceptedTypes {
		if acceptedType == typeLookup {
			return i
		}
	}

	return constants.InvalidIndex
}
