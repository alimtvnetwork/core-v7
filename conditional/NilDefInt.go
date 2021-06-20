package conditional

import "gitlab.com/evatix-go/core/constants"

func NilDefInt(
	valuePointer *int,
) int {
	if valuePointer == nil {
		return constants.Zero
	}

	return *valuePointer
}

func NilDefIntPtr(
	valuePointer *int,
) *int {
	if valuePointer == nil {
		return constants.ZeroPtr
	}

	return valuePointer
}

func NilDefValInt(
	valuePointer *int,
	defVal int,
) int {
	if valuePointer == nil {
		return defVal
	}

	return *valuePointer
}
