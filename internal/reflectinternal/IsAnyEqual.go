package reflectinternal

import (
	"reflect"
)

func IsAnyEqual(left, right interface{}) bool {
	isEqual, isConclusive := IsConclusive(left, right)

	if isConclusive {
		return isEqual
	}

	return reflect.DeepEqual(left, right)
}
