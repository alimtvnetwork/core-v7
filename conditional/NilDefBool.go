package conditional

import "gitlab.com/evatix-go/core/constants"

func NilDefBool(
	valuePointer *bool,
) bool {
	if valuePointer == nil {
		return false
	}

	return *valuePointer
}

func NilDefBoolPtr(
	valuePointer *bool,
) *bool {
	if valuePointer == nil {
		return constants.FalseBoolPtr
	}

	return valuePointer
}

func NilBoolVal(
	valuePointer *bool,
	defVal bool,
) bool {
	if valuePointer == nil {
		return defVal
	}

	return *valuePointer
}

func NilBoolValPtr(
	valuePointer *bool,
	defVal bool,
) *bool {
	if valuePointer == nil {
		return &defVal
	}

	return valuePointer
}
